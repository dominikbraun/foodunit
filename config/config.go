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

// Package config provides tools for reading and processing configurations.
package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Reader interface {
	Get(key string) interface{}
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
}

var (
	ErrFileNotFound = errors.New("the configuration file was not found.")
)

func NewFile() (Reader, error) {
	reader := viper.New()

	reader.SetConfigName("config")
	reader.AddConfigPath("$HOME/foodunit")
	reader.AddConfigPath(".")

	if err := reader.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, ErrFileNotFound
		}
		return nil, err
	}

	return reader, nil
}
