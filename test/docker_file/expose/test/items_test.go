package test

import (
	"fmt"
	"go_expose/internal/handler"
	"io"
	"net/http/httptest"
	"testing"
)

var handl = handler.NewItemHandler()

func PrintRecorderJSON(res *httptest.ResponseRecorder) {
	body , _ := io.ReadAll(res.Result().Body)
	fmt.Println(string(body))
}
func TestGetAll(t *testing.T) {
	request := httptest.NewRequest("GET","http://localhost:8081",nil)
	recorder := httptest.NewRecorder()
	handl.GetAll(recorder,request)
	PrintRecorderJSON(recorder)
}

