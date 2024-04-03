package json
import (
	"net/http"
	"encoding/json"
)

func WriteJSON(w http.ResponseWriter, code int, data any) {
	dataJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		http.Error(w, "error converting data to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(dataJSON)
}