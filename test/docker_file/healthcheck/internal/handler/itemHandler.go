package handler

import (
	"bytes"
	"encoding/json"
	"go_wd/internal/memory"
	helperJson "go_wd/pkg/json"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/google/uuid"
)

type generalResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ItemHandler struct {
	*memory.Stores
}

func NewItemHandler() *ItemHandler {
	return &ItemHandler{Stores: memory.NewStores()}
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

func (i *ItemHandler) GetAll(w http.ResponseWriter, r *http.Request)  {
	items := i.Stores.FindAll()
	helperJson.WriteJSON(w,http.StatusOK,&generalResponse{Data: items, Message: "successfully get all"})
}

func (i *ItemHandler) CreateItems(w http.ResponseWriter , r *http.Request) {
	var item memory.Item
	err := helperJson.DecodeJSONBody(w, r, &item)
	if m,ok := err.(*helperJson.MalformedRequest);err != nil && ok {
		helperJson.WriteJSON(w,m.GetStatus(),&generalResponse{Data: nil,Message: m.GetMessage()})
		return
	}
	i.Save(&item)
	helperJson.WriteJSON(w,http.StatusOK,&generalResponse{Data: nil , Message:  "successfully create"})
}

func (i *ItemHandler) HealthItems(w http.ResponseWriter,r *http.Request)  {
	
}


func (i *ItemHandler) seed(wg *sync.WaitGroup,total int) {
	item := &memory.Item{Name: "Test",Qtt: 10,Price: &memory.Currency{Code: "US",Amount: 40.0,Name: "Dollar",Symbol: "$"}}
	payload , _ := json.MarshalIndent(item,"", "    ")

	for n := 0; n < total; n++ {
		wg.Add(1)
		go func() {
			request := httptest.NewRequest("POST", "http://localhost:8081", bytes.NewBuffer(payload))
			recorder := httptest.NewRecorder()
			i.CreateItems(recorder,request)
			defer wg.Done()
		}()
	}
}