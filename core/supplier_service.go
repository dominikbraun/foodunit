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

// Package core provides implementations of the DL services. They provide
// the actual business logic and can be called by any interface adapter.
package core

import "github.com/dominikbraun/foodunit/dl"

// SupplierService implements dl.SupplierService.
type SupplierService struct {
	dao dl.DataAccess
}

// NewSupplierService creates a new instance of SupplierService.
func NewSupplierService(dao dl.DataAccess) *SupplierService {
	return &SupplierService{dao: dao}
}

// Get implements dl.SupplierService.Get.
func (s SupplierService) Get(id uint64) (dl.Supplier, error) {
	return s.dao.Suppliers().FindByID(id)
}
