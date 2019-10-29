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

// Package user provides services and types for User-related data.
package user

// Registration represents a user registration.
type Registration struct {
	MailAddr       string `json:"mail_addr"`
	Name           string `json:"name"`
	PaypalMailAddr string `json:"paypal_mail_addr"`
	Password       string `json:"password"`
}

// Login represents an user's attempt to authenticate himself.
type Login struct {
	MailAddr string `json:"mail_addr"`
	Password string `json:"password"`
}

// PaypalMailAddrSetter is a setter type for the PaypalMailAddr property.
type PaypalMailAddrSetter struct {
	PaypalMailAddr string `json:"paypal_mail_addr"`
}

// PublicUser is the public representation of an user (i. e. profile).
type PublicUser struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
