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

// Package fileserver provides a server which serves all static frontend files.
package fileserver

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"os/signal"
)

// FileServer represents a server providing frontend files for the client.
// It runs as an independent standalone server and will be treated correspondingly.
type FileServer struct {
	router    *chi.Mux
	interrupt chan os.Signal
}

// Setup creates a new FileServer instance and mounts its own router.
func Setup(clientURL string) (*FileServer, error) {
	router := provideRouter()

	fs := FileServer{
		router:    router,
		interrupt: make(chan os.Signal),
	}

	signal.Notify(fs.interrupt, os.Interrupt)

	if clientURL != "" {
		if err := fs.useReverseProxy(clientURL); err != nil {
			return nil, err
		}
		return &fs, nil
	}
	fs.mountRoutes()

	return &fs, nil
}

// Run makes the FileServer listen on the specified port, serving static files.
func (fs *FileServer) Run() {
	go func() {
		log.Fatal(http.ListenAndServe(":4040", fs.router))
	}()

	<-fs.interrupt
}
