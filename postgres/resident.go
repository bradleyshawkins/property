package postgres

import (
	"context"
	"github.com/bradleyshawkins/property"
)

func (t *Transaction) GetResident(ctx context.Context, residentID *property.ResidentID) (*property.Resident, error) {

	return nil, nil
}

func (t *Transaction) SaveResident(ctx context.Context, resident *property.Resident) error {
	return nil
}
