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

func (i *ItemHandler) GetItemByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_, err := uuid.Parse(id)
	if err != nil {
		helperJson.WriteJSON(w, http.StatusBadRequest, &generalResponse{Data: nil, Message: err.Error()})
		return
	}
	item, err := i.Stores.FindByID(id)
	if err != nil {
		helperJson.WriteJSON(w, http.StatusNotFound, &generalResponse{Data: nil, Message: err.Error()})
		return
	}
	helperJson.WriteJSON(w, http.StatusOK, &generalResponse{Data: item, Message: "successfully get id"})
}

func (i *ItemHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	items := i.Stores.FindAll()
	helperJson.WriteJSON(w, http.StatusOK, &generalResponse{Data: items, Message: "successfully get all"})
}

func (i *ItemHandler) CreateItems(w http.ResponseWriter, r *http.Request) {
	var item memory.Item
	err := helperJson.DecodeJSONBody(w, r, &item)
	if m, ok := err.(*helperJson.MalformedRequest); err != nil && ok {
		helperJson.WriteJSON(w, m.GetStatus(), &generalResponse{Data: nil, Message: m.GetMessage()})
		return
	}
	i.Save(&item)
	helperJson.WriteJSON(w, http.StatusOK, &generalResponse{Data: nil, Message: "successfully create"})
}

func (i *ItemHandler) HealthItems(w http.ResponseWriter, r *http.Request) {
	const total = 10
	var wg sync.WaitGroup
	codes, err := i.seed(&wg, total)
	wg.Wait()
	for _, code := range codes {
		if code == http.StatusInternalServerError {
			helperJson.WriteJSON(w, code, &generalResponse{Message: "unhealthy", Data: map[string]any{
				"code":  code,
				"error": err.Error(),
			}})
			return
		}
	}
	request := httptest.NewRequest("GET", "http://localhost:8081/items", nil)
	recorder := httptest.NewRecorder()
	i.GetAll(recorder, request)
	if recorder.Result().StatusCode != http.StatusOK {
		helperJson.WriteJSON(w, http.StatusInternalServerError, &generalResponse{Message: "unhealthy", Data: nil})
		return
	}
	helperJson.WriteJSON(w, recorder.Result().StatusCode, &generalResponse{Message: "healthy", Data: nil})
}

func (i *ItemHandler) seed(wg *sync.WaitGroup, total int) ([]int, error) {
	var actualyCodesStatus []int
	item := &memory.Item{Name: "Test", Qtt: 10, Price: &memory.Currency{Code: "US", Amount: 40.0, Name: "Dollar", Symbol: "$"}}
	payload, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		return []int{http.StatusInternalServerError}, err
	}
	for n := 0; n < total; n++ {
		wg.Add(1)
		go func() {
			request := httptest.NewRequest("POST", "http://localhost:8081", bytes.NewBuffer(payload))
			recorder := httptest.NewRecorder()
			i.CreateItems(recorder, request)
			actualyCodesStatus = append(actualyCodesStatus, recorder.Result().StatusCode)
			defer wg.Done()
		}()
	}
	return actualyCodesStatus, nil
}
