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
	"github.com/pkg/errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// mountRoutes makes the file server to either directly serve the files or use a
// reverse proxy to prevent cross-origin errors.
func (fs *FileServer) mountRoutes() {
	dir := "ui/dist"

	fileServer(fs.router, "/js", http.Dir(dir+"/js"))
	fileServer(fs.router, "/css", http.Dir(dir+"/css"))
	fileServer(fs.router, "/img", http.Dir(dir+"/img"))

	// Redirect all other requests to index.html
	fs.router.NotFound(func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, dir+"/index.html")
	})
}

// useReverseProxy sets up a single host reverse proxy to
func (fs *FileServer) useReverseProxy(clientURL string) error {
	client, err := url.Parse(clientURL)
	if err != nil {
		return errors.Wrap(err, "client URL is not valid")
	}

	fs.router.NotFound(func(res http.ResponseWriter, req *http.Request) {
		httputil.NewSingleHostReverseProxy(client).ServeHTTP(res, req)
	})

	return nil
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}
