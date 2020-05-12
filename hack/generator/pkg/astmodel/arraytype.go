/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
)

// ArrayType is used for fields that contain an array of values
type ArrayType struct {
	element Type
}

// NewArrayType creates a new array with elements of the specified type
func NewArrayType(element Type) *ArrayType {
	return &ArrayType{element}
}

// assert we implemented Type correctly
var _ Type = (*ArrayType)(nil)

// AsType renders the Go abstract syntax tree for an array type
func (array *ArrayType) AsType() ast.Expr {
	return &ast.ArrayType{
		Elt: array.element.AsType(),
	}
}

func (array *ArrayType) RequiredImports() []PackageReference {
	return array.element.RequiredImports()
}

func (array *ArrayType) References(t Type) bool {
	return array == t || array.element.References(t)
}