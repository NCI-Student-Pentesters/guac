//
// Copyright 2022 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package file

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/guacsec/guac/pkg/events"
	"github.com/guacsec/guac/pkg/handler/processor"
)

const (
	FileCollector = "FileCollector"
)

type fileCollector struct {
	path        string
	lastChecked time.Time
	poll        bool
	interval    time.Duration
	useBlobURL  bool
}

func NewFileCollector(ctx context.Context, path string, poll bool, interval time.Duration, useBlobURL bool) *fileCollector {
	return &fileCollector{
		path:       path,
		poll:       poll,
		interval:   interval,
		useBlobURL: useBlobURL,
	}
}

// RetrieveArtifacts collects the documents from the collector. It emits each collected
// document through the channel to be collected and processed by the upstream processor.
// The function should block until all the artifacts are collected and return a nil error
// or return an error from the collector crashing. This function can keep running and check
// for new artifacts as they are being uploaded by polling on an interval or run once and
// grab all the artifacts and end.
func (f *fileCollector) RetrieveArtifacts(ctx context.Context, docChannel chan<- *processor.Document) error {
	if _, err := os.Stat(f.path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("path: %s does not exist", f.path)
		}
		return fmt.Errorf("unknown error on os.Stat for FileCollector path: %w", err)
	}

	readFunc := func(path string, dirEntry fs.DirEntry, err error) error {
		// If the context has been canceled it contains an err which we can throw.
		// When it gets thrown a second time will cancel the walk.
		// See filepath.WalkDir for more info.
		if ctx.Err() != nil {
			return ctx.Err() // nolint:wrapcheck
		}
		// NOTE: Explicitly rethrowing new errors if a particular directory has an error.
		// If we rethrow the error it kills the whole walk. Still useful to make it explicit that we ran into an error.
		if err != nil {
			return fmt.Errorf("path: %s is invalid", path)
		}
		if dirEntry.IsDir() {
			return nil
		}
		info, err := dirEntry.Info()
		if err != nil {
			return fmt.Errorf("unknown error on dirEntry.Info while walking path: %w", err)
		}
		if !info.ModTime().After(f.lastChecked) {
			return nil
		}

		blob, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file: %s, err: %w", path, err)
		}

		var docRef string
		if f.useBlobURL {
			docRef = events.GetKey(blob) // this is the blob store path
		} else {
			docRef = ""
		}

		doc := &processor.Document{
			Blob:   blob,
			Type:   processor.DocumentUnknown,
			Format: processor.FormatUnknown,
			SourceInformation: processor.SourceInformation{
				Collector:   string(FileCollector),
				Source:      fmt.Sprintf("file:///%s", path),
				DocumentRef: docRef,
			},
		}

		docChannel <- doc

		return nil
	}

	for {
		if err := filepath.WalkDir(f.path, readFunc); err != nil {
			return fmt.Errorf("error walking path: %s, err: %w", f.path, err)
		}
		f.lastChecked = time.Now()
		if !f.poll {
			break
		}
		select {
		// If the context has been canceled it contains an err which we can throw.
		case <-ctx.Done():
			return ctx.Err() // nolint:wrapcheck
		case <-time.After(f.interval):
		}
	}

	return nil
}

// Type returns the collector type
func (f *fileCollector) Type() string {
	return FileCollector
}
