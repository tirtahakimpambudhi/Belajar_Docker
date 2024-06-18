package json


import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "strings"
)

type MalformedRequest struct {
    status int
    msg    string
}

func (mr *MalformedRequest) GetStatus() int {
    return mr.status
}

func (mr *MalformedRequest) GetMessage() string {
    return mr.msg
}

func (mr *MalformedRequest) Error() string {
    return mr.msg
}

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
    ct := r.Header.Get("Content-Type")
    if ct != "" {
        mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
        if mediaType != "application/json" {
            msg := "Content-Type header is not application/json"
            return &MalformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
        }
    }

    r.Body = http.MaxBytesReader(w, r.Body, 1048576)

    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()

    err := dec.Decode(&dst)
    if err != nil {
        var syntaxError *json.SyntaxError
        var unmarshalTypeError *json.UnmarshalTypeError

        switch {
        case errors.As(err, &syntaxError):
            msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
            return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

        case errors.Is(err, io.ErrUnexpectedEOF):
            msg := fmt.Sprint("Request body contains badly-formed JSON")
            return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

        case errors.As(err, &unmarshalTypeError):
            msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
            return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

        case strings.HasPrefix(err.Error(), "json: unknown field "):
            fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
            msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
            return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

        case errors.Is(err, io.EOF):
            msg := "Request body must not be empty"
            return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

        case err.Error() == "http: request body too large":
            msg := "Request body must not be larger than 1MB"
            return &MalformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}

        default:
            return err
        }
    }

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
        msg := "Request body must only contain a single JSON object"
        return &MalformedRequest{status: http.StatusBadRequest, msg: msg}
    }

    return nil
}