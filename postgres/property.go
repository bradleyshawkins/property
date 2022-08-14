package postgres

import (
	"context"
	"github.com/bradleyshawkins/property"
)

func (t *Transaction) GetProperty(ctx context.Context, propertyID *property.PropertyID) (*property.Property, error) {
	return nil, nil
}

func (t *Transaction) SaveProperty(ctx context.Context, property *property.Property) error {
	return nil
}
