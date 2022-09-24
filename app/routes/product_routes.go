// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/pkg/server"

	"github.com/gorilla/mux"
)

func productRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/product", svr.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/product", svr.DeleteProduct).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/product", svr.EditProduct).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/product", svr.GetProduct).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/products", svr.GetAllProducts).Methods(http.MethodGet)
}
