package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"github.com/jim-minter/origin-template-service-broker/pkg/broker"

	kerrors "k8s.io/kubernetes/pkg/api/errors"
)

func readRequest(r *http.Request, obj interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("error: invalid content-type")
	}

	return json.NewDecoder(r.Body).Decode(&obj)
}

func writeResponse(w http.ResponseWriter, code int, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	// return json.NewEncoder(w).Encode(obj)

	// pretty-print for easier debugging
	b, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	i := bytes.Buffer{}
	json.Indent(&i, b, "", "  ")
	i.WriteString("\n")
	_, err = w.Write(i.Bytes())
	return err
}

func writeDefaultResponse(w http.ResponseWriter, code int, resp interface{}, err error) error {
	if err == nil {
		return writeResponse(w, code, resp)
	} else if statusErr, ok := err.(*kerrors.StatusError); ok {
		return writeResponse(w, int(statusErr.ErrStatus.Code), broker.ErrorResponse{Description: err.Error()})
	} else {
		return writeResponse(w, http.StatusInternalServerError, broker.ErrorResponse{Description: err.Error()})
	}
}
