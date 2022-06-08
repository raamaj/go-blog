package controllers

import (
	"net/http"

	"github.com/raamaj/go-blog/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome Home!")
}
