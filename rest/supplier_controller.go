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

// Package rest provides methods for handling REST API requests.
package rest

import (
	"github.com/go-chi/render"
	"net/http"
)

var (
	// SupplierController serves as a ready-to-go controller instance.
	SupplierController = supplierController{}
)

// supplierController groups HTTP handlers for Supplier-related requests.
type supplierController struct{}

// GetInfo handles core.SupplierService.GetInfo.
func (supplierController) GetInfo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, "GetInfo")
	})
}

// GetMenu handles core.SupplierService.GetMenu.
func (supplierController) GetMenu() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, "GetMenu")
	})
}
