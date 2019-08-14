// Package http provides HTTP handlers, functioning as
// interfaces between the API server and the use cases.
package http

import (
	"net/http"

	"github.com/go-chi/render"
)

// PostOffer creates a new offer.
func PostOffer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "PostOffer")
}

// GetOffer renders a given offer.
func GetOffer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "GetOffer")
}

// PutOffer updates a given offer.
func PutOffer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "PutOffer")
}

// DeleteOffer deletes a given offer.
func DeleteOffer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "DeleteOffer")
}

// GetOrders renders all orders for the given offer.
func GetOrders(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "GetOrders")
}

// PostOrder creates an order associated with the current user.
func PostOrder(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "PostOrder")
}

// PutOrder updates an order associated with the current user.
func PutOrder(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "PutOrder")
}

func F(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "Needs to be implemented")
}
