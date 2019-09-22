// Copyright 2019 The FoodUnit Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package session provides interfaces and implementations for session managers.
package session

import (
	"context"
	"github.com/alexedwards/scs/v2"
	"net/http"
)

type Manager interface {
	RenewToken(ctx context.Context) error
	Get(ctx context.Context, key string) interface{}
	GetString(ctx context.Context, key string) string
	GetBool(ctx context.Context, key string) bool
	Put(ctx context.Context, key string, val interface{})
	LoadAndSave(next http.Handler) http.Handler
}

func NewManager() Manager {
	manager := scs.New()
	manager.Cookie.HttpOnly = true

	return manager
}
