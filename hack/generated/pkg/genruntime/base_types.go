/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type MetaObject interface {
	runtime.Object
	metav1.Object
	KubernetesResource
}

type KubernetesResource interface {
	// Owner returns the ResourceReference of the owner, or nil if there is no owner
	Owner() *ResourceReference

	// AzureName returns the Azure name of the resource
	AzureName() string
}

// ArmResourceSpec is an ARM resource specification
type ArmResourceSpec interface {
	GetApiVersion() string

	GetType() string

	GetName() string
}

type ArmResource interface {
	ArmResourceSpec
	// ArmResourceStatus TODO: ???

	GetId() string
}

// TODO: I think that this is throwaway?
type ArmResourceImpl struct {
	ArmResourceSpec
	Id string
}

var _ ArmResource = &ArmResourceImpl{}

func (resource *ArmResourceImpl) GetId() string {
	return resource.Id
}
