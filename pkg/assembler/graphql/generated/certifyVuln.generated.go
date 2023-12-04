// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _CertifyVuln_id(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVuln) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVuln_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVuln_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVuln",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVuln_package(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVuln) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVuln_package(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Package, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Package)
	fc.Result = res
	return ec.marshalNPackage2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackage(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVuln_package(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVuln",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Package_id(ctx, field)
			case "type":
				return ec.fieldContext_Package_type(ctx, field)
			case "namespaces":
				return ec.fieldContext_Package_namespaces(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Package", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVuln_vulnerability(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVuln) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVuln_vulnerability(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Vulnerability, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Vulnerability)
	fc.Result = res
	return ec.marshalNVulnerability2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVulnerability(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVuln_vulnerability(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVuln",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Vulnerability_id(ctx, field)
			case "type":
				return ec.fieldContext_Vulnerability_type(ctx, field)
			case "vulnerabilityIDs":
				return ec.fieldContext_Vulnerability_vulnerabilityIDs(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Vulnerability", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyVuln_metadata(ctx context.Context, field graphql.CollectedField, obj *model.CertifyVuln) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyVuln_metadata(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Metadata, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.ScanMetadata)
	fc.Result = res
	return ec.marshalNScanMetadata2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐScanMetadata(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyVuln_metadata(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyVuln",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "timeScanned":
				return ec.fieldContext_ScanMetadata_timeScanned(ctx, field)
			case "dbUri":
				return ec.fieldContext_ScanMetadata_dbUri(ctx, field)
			case "dbVersion":
				return ec.fieldContext_ScanMetadata_dbVersion(ctx, field)
			case "scannerUri":
				return ec.fieldContext_ScanMetadata_scannerUri(ctx, field)
			case "scannerVersion":
				return ec.fieldContext_ScanMetadata_scannerVersion(ctx, field)
			case "origin":
				return ec.fieldContext_ScanMetadata_origin(ctx, field)
			case "collector":
				return ec.fieldContext_ScanMetadata_collector(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type ScanMetadata", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _ScanMetadata_timeScanned(ctx context.Context, field graphql.CollectedField, obj *model.ScanMetadata) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ScanMetadata_timeScanned(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.TimeScanned, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ScanMetadata_timeScanned(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ScanMetadata",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ScanMetadata_dbUri(ctx context.Context, field graphql.CollectedField, obj *model.ScanMetadata) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ScanMetadata_dbUri(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DbURI, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ScanMetadata_dbUri(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ScanMetadata",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ScanMetadata_dbVersion(ctx context.Context, field graphql.CollectedField, obj *model.ScanMetadata) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ScanMetadata_dbVersion(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DbVersion, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ScanMetadata_dbVersion(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ScanMetadata",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ScanMetadata_scannerUri(ctx context.Context, field graphql.CollectedField, obj *model.ScanMetadata) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ScanMetadata_scannerUri(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ScannerURI, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ScanMetadata_scannerUri(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ScanMetadata",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ScanMetadata_scannerVersion(ctx context.Context, field graphql.CollectedField, obj *model.ScanMetadata) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ScanMetadata_scannerVersion(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ScannerVersion, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ScanMetadata_scannerVersion(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ScanMetadata",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ScanMetadata_origin(ctx context.Context, field graphql.CollectedField, obj *model.ScanMetadata) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ScanMetadata_origin(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Origin, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ScanMetadata_origin(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ScanMetadata",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ScanMetadata_collector(ctx context.Context, field graphql.CollectedField, obj *model.ScanMetadata) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ScanMetadata_collector(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Collector, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ScanMetadata_collector(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ScanMetadata",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputCertifyVulnSpec(ctx context.Context, obj interface{}) (model.CertifyVulnSpec, error) {
	var it model.CertifyVulnSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "package", "vulnerability", "timeScanned", "dbUri", "dbVersion", "scannerUri", "scannerVersion", "origin", "collector"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalOID2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		case "package":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("package"))
			data, err := ec.unmarshalOPkgSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgSpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.Package = data
		case "vulnerability":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("vulnerability"))
			data, err := ec.unmarshalOVulnerabilitySpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐVulnerabilitySpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.Vulnerability = data
		case "timeScanned":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("timeScanned"))
			data, err := ec.unmarshalOTime2ᚖtimeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
			it.TimeScanned = data
		case "dbUri":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("dbUri"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.DbURI = data
		case "dbVersion":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("dbVersion"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.DbVersion = data
		case "scannerUri":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("scannerUri"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.ScannerURI = data
		case "scannerVersion":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("scannerVersion"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.ScannerVersion = data
		case "origin":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Origin = data
		case "collector":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("collector"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Collector = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputScanMetadataInput(ctx context.Context, obj interface{}) (model.ScanMetadataInput, error) {
	var it model.ScanMetadataInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"timeScanned", "dbUri", "dbVersion", "scannerUri", "scannerVersion", "origin", "collector"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "timeScanned":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("timeScanned"))
			data, err := ec.unmarshalNTime2timeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
			it.TimeScanned = data
		case "dbUri":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("dbUri"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.DbURI = data
		case "dbVersion":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("dbVersion"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.DbVersion = data
		case "scannerUri":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("scannerUri"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ScannerURI = data
		case "scannerVersion":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("scannerVersion"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ScannerVersion = data
		case "origin":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Origin = data
		case "collector":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("collector"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Collector = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var certifyVulnImplementors = []string{"CertifyVuln", "Node"}

func (ec *executionContext) _CertifyVuln(ctx context.Context, sel ast.SelectionSet, obj *model.CertifyVuln) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, certifyVulnImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("CertifyVuln")
		case "id":
			out.Values[i] = ec._CertifyVuln_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "package":
			out.Values[i] = ec._CertifyVuln_package(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "vulnerability":
			out.Values[i] = ec._CertifyVuln_vulnerability(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "metadata":
			out.Values[i] = ec._CertifyVuln_metadata(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var scanMetadataImplementors = []string{"ScanMetadata"}

func (ec *executionContext) _ScanMetadata(ctx context.Context, sel ast.SelectionSet, obj *model.ScanMetadata) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, scanMetadataImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ScanMetadata")
		case "timeScanned":
			out.Values[i] = ec._ScanMetadata_timeScanned(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "dbUri":
			out.Values[i] = ec._ScanMetadata_dbUri(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "dbVersion":
			out.Values[i] = ec._ScanMetadata_dbVersion(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "scannerUri":
			out.Values[i] = ec._ScanMetadata_scannerUri(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "scannerVersion":
			out.Values[i] = ec._ScanMetadata_scannerVersion(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "origin":
			out.Values[i] = ec._ScanMetadata_origin(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "collector":
			out.Values[i] = ec._ScanMetadata_collector(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNCertifyVuln2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVulnᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.CertifyVuln) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNCertifyVuln2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVuln(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNCertifyVuln2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVuln(ctx context.Context, sel ast.SelectionSet, v *model.CertifyVuln) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._CertifyVuln(ctx, sel, v)
}

func (ec *executionContext) unmarshalNCertifyVulnSpec2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVulnSpec(ctx context.Context, v interface{}) (model.CertifyVulnSpec, error) {
	res, err := ec.unmarshalInputCertifyVulnSpec(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNScanMetadata2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐScanMetadata(ctx context.Context, sel ast.SelectionSet, v *model.ScanMetadata) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._ScanMetadata(ctx, sel, v)
}

func (ec *executionContext) unmarshalNScanMetadataInput2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐScanMetadataInput(ctx context.Context, v interface{}) (model.ScanMetadataInput, error) {
	res, err := ec.unmarshalInputScanMetadataInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNScanMetadataInput2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐScanMetadataInputᚄ(ctx context.Context, v interface{}) ([]*model.ScanMetadataInput, error) {
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]*model.ScanMetadataInput, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNScanMetadataInput2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐScanMetadataInput(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) unmarshalNScanMetadataInput2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐScanMetadataInput(ctx context.Context, v interface{}) (*model.ScanMetadataInput, error) {
	res, err := ec.unmarshalInputScanMetadataInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
