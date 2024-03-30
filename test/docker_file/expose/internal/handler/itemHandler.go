package handler

import (
	"go_expose/internal/memory"
	"net/http"
	helperJson "go_expose/pkg/json"
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
		helperJson.WriteJSON(w,http.StatusBadRequest,&generalResponse{Data: nil , Message: err.Error()})
		return
	}
	item, err := i.Stores.FindByID(id)
	if err != nil {
		helperJson.WriteJSON(w,http.StatusNotFound,&generalResponse{Data: nil , Message: err.Error()})
		return
	}
	helperJson.WriteJSON(w,http.StatusOK,&generalResponse{Data: item, Message: "successfully get id"})
}

func (i *ItemHandler) GetAll()  {
	
}