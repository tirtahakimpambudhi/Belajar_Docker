package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_env/internal/handler"
	"go_env/internal/memory"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	_ "time"

	"github.com/stretchr/testify/require"
)

var handl = handler.NewItemHandler()

func PrintRecorderJSON(res *httptest.ResponseRecorder) {
	body, _ := io.ReadAll(res.Result().Body)
	fmt.Println(string(body))
}

func ItemSeed(total int) {

}

func TestGetAll(t *testing.T) {
	var wg sync.WaitGroup
	Create(t,&wg,10)
	wg.Wait()
	request := httptest.NewRequest("GET", "http://localhost:8081", nil)
	recorder := httptest.NewRecorder()
	handl.GetAll(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	PrintRecorderJSON(recorder)
}

func Create(t *testing.T,wg *sync.WaitGroup,total int) {
	item := &memory.Item{Name: "Test",Qtt: 10,Price: &memory.Currency{Code: "US",Amount: 40.0,Name: "Dollar",Symbol: "$"}}
	payload , err := json.MarshalIndent(item,"", "    ")
	require.NoError(t,err)
	for i := 0; i < total; i++ {
		wg.Add(1)
		go func() {
			request := httptest.NewRequest("POST", "http://localhost:8081", bytes.NewBuffer(payload))
			recorder := httptest.NewRecorder()
			handl.CreateItems(recorder,request)
			require.Equal(t, http.StatusOK, recorder.Result().StatusCode)
			defer wg.Done()
		}()
	}
}



