package handler

import (
	"encoding/json"
	"net/http"
	"time"
)



type Handler struct {
	
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request)  {
    // Mendapatkan alamat IP client
    clientIP := r.RemoteAddr

    // Mendapatkan user agent client
    userAgent := r.UserAgent()

    // Mendapatkan metode permintaan HTTP
    requestMethod := r.Method

    // Mendapatkan header permintaan HTTP
    requestHeaders := r.Header

    // Mendapatkan path permintaan HTTP
    requestPath := r.URL.Path

    // Mendapatkan query string permintaan HTTP
    requestQuery := r.URL.Query()

    // Mendapatkan waktu permintaan HTTP
    requestTime := time.Now()

    // Menyimpan informasi yang ingin Anda sertakan
    responseData := map[string]interface{}{
        "status":         "OK",
        "message":        "SERVER 1",
        "client":         map[string]interface{}{
            "ip":         clientIP,
            "user_agent": userAgent,
            // Anda bisa menambahkan lebih banyak informasi tentang client di sini
        },
        "request_info":   map[string]interface{}{
            "method":     requestMethod,
            "headers":    requestHeaders,
            "path":       requestPath,
            "query":      requestQuery,
            "request_time": requestTime.Format(time.RFC3339),
            // Anda bisa menambahkan lebih banyak informasi tentang permintaan di sini
        },
    }

    // Mengirim respons dalam format JSON
    h.writeJSON(w, http.StatusOK, responseData)
    return
}

func (h *Handler) writeJSON(w http.ResponseWriter, code int, data any) {
	dataJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		http.Error(w, "error converting data to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(dataJSON)
}