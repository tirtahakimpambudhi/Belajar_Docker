package handler

import (
	"go_expose/internal/memory"
	"net/http"

	"github.com/google/uuid"
)

type generalResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ItemHandler struct {
	*memory.Stores
}

func (i *ItemHandler) GetItemByID(w http.ResponseWriter, r *http.Request)  {
	id := r.PathValue("id")
	_, err := uuid.Parse(id)
	if err != nil {
		
		return
	}
}