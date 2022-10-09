//go:build go1.18
// +build go1.18

package typeparams

import (
	"go/ast"
	"go/types"
)

func HasTypeParam(t types.Type) bool {
	switch t := t.(type) {
	case *types.TypeParam:
		return true
	case *types.Named:
		return t.TypeParams() != nil
	case *types.Signature:
		return t.TypeParams() != nil
	}
	return false
}

func NamedHasTypeParam(t *types.Named) bool {
	return t.TypeParams() != nil
}

func FuncTypeHasTypeParam(t *ast.FuncType) bool {
	return t.TypeParams != nil
}

func TypeSpecRemoveTypeParam(spec *ast.TypeSpec) {
	spec.TypeParams = nil
}
