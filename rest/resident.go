package rest

import (
	"encoding/json"
	"github.com/bradleyshawkins/property"
	"github.com/bradleyshawkins/property/moving"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type moveInRequest struct {
	PropertyID string `json:"propertyID"`
}

func (s *Server) MoveIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	residentID, propertyID, err := getResidentAndPropertyID(r)
	if err != nil {
		return
	}

	tx, err := s.database.BeginTransaction()
	if err != nil {
		return
	}

	defer tx.RollbackFunc(err)

	moveIn := moving.NewMoveIn(tx, tx)

	err = moveIn.MoveIn(ctx, residentID, propertyID)
	if err != nil {
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}
}

func getResidentAndPropertyID(r *http.Request) (*property.ResidentID, *property.PropertyID, error) {
	residentID, err := getResidentID(r)
	if err != nil {
		return nil, nil, err
	}

	propertyID, err := getPropertyID(r)
	if err != nil {
		return nil, nil, err
	}

	return residentID, propertyID, nil
}

func getResidentID(r *http.Request) (*property.ResidentID, error) {
	residentIDParam := chi.URLParam(r, "residentID")

	residentID, err := uuid.FromString(residentIDParam)
	if err != nil {
		return nil, err
	}

	return &property.ResidentID{
		UUID: residentID,
	}, nil
}

func getPropertyID(r *http.Request) (*property.PropertyID, error) {
	var moveInRequest moveInRequest
	err := json.NewDecoder(r.Body).Decode(&moveInRequest)
	if err != nil {
		return nil, err
	}

	propertyID, err := uuid.FromString(moveInRequest.PropertyID)
	if err != nil {
		return nil, err
	}

	return &property.PropertyID{UUID: propertyID}, nil
}
