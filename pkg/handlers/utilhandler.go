package handlers

import (
	"encoding/json"
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func ValidateAndGetPaymentFromBody(w http.ResponseWriter, r *http.Request) (*payments.Payment, error) {
	if r.Body == nil {
		WriteJSONMessage(w, http.StatusBadRequest, "body of request must be a json object")
		return nil, errors.New("invalid json")
	}

	decoder := json.NewDecoder(r.Body)
	var requestBody payments.Payment
	err := decoder.Decode(&requestBody)

	if err != nil {
		WriteJSONMessage(w, http.StatusBadRequest, "invalid json")
		return nil, errors.New("invalid json")
	}

	return &requestBody, nil
}

func ValidateAndGetUUIDFromParams(w http.ResponseWriter, r *http.Request) (uuid.UUID, error){
	vars := mux.Vars(r)
	id := vars["uuid"]

	uid, err := uuid.Parse(id)
	if err != nil {
		WriteJSONMessage(w, http.StatusBadRequest, "invalid UUID")
		return uuid.UUID{}, errors.New("invalid UUID")
	}

	return uid, nil
}

func ValidateExistence(exists bool, err error, w http.ResponseWriter) error{
	if err != nil {
		log.Printf("[ERROR] cannot delete payment: %s\n", err)
		Write500(w)
		return err
	}

	if !exists {
		WriteJSONMessage(w, http.StatusNotFound, "payment doesnt exist")
		return errors.New("payment doesnt exist")
	}

	return nil
}
