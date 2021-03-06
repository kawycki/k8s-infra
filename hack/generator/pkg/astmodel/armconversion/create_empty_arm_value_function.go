/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/dave/dst"
)

// CreateEmptyArmValueFunc represents a function that creates
// an empty value suitable for calling PopulateFromARM against.
// It should be equivalent to ConvertToARM("") on a default struct value.
type CreateEmptyArmValueFunc struct {
	idFactory   astmodel.IdentifierFactory
	armTypeName astmodel.TypeName
}

var _ astmodel.Function = &CreateEmptyArmValueFunc{}

func (f CreateEmptyArmValueFunc) AsFunc(c *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {
	fn := &astbuilder.FuncDetails{
		Name:          "CreateEmptyArmValue",
		ReceiverIdent: f.idFactory.CreateIdentifier(receiver.Name(), astmodel.NotExported),
		ReceiverType: &dst.StarExpr{
			X: dst.NewIdent(receiver.Name()),
		},
		Body: []dst.Stmt{
			astbuilder.Returns(&dst.CompositeLit{
				Type: dst.NewIdent(f.armTypeName.Name()),
			}),
		},
	}

	fn.AddReturns("interface{}")
	fn.AddComments("returns an empty ARM value suitable for deserializing into")

	return fn.DefineFunc()
}

func (f CreateEmptyArmValueFunc) Equals(other astmodel.Function) bool {
	otherFunc, ok := other.(CreateEmptyArmValueFunc)
	return ok && f.armTypeName == otherFunc.armTypeName
}

func (f CreateEmptyArmValueFunc) Name() string {
	return "CreateEmptyArmValue"
}

func (f CreateEmptyArmValueFunc) References() astmodel.TypeNameSet {
	return nil
}

func (f CreateEmptyArmValueFunc) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet()
}
