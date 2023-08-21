// Copyright 2019 WhizUs GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package common

import (
	"context"
)

const (
	UsersEndpoint        = "/_plugins/_security/api/internalusers/"
	RolesEndpoint        = "/_plugins/_security/api/roles/"
	RolesMappingEndpoint = "/_plugins/_security/api/rolesmapping/"
	ActiongroupEndpoint  = "/_plugins/_security/api/actiongroups/"
	TenantEndpoint       = "/_plugins/_security/api/tenants/"
	HealthEndpoint       = "/_plugins/_security/health"
)

type Service struct {
	Client ClientInterface
}

type Patch struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type ClientInterface interface {
	Do(ctx context.Context, reqBytes interface{}, endpoint string, method string) ([]byte, error)
	Get(ctx context.Context, path string, T interface{}) error
	GetBaseURL() string
	Modify(ctx context.Context, path string, method string, reqBytes interface{}) error
}

type Modifyable interface {
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, name string, patches *[]Patch) error
	UpdateBatch(ctx context.Context, patches *[]Patch) error
}
