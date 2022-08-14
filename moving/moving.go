package moving

import (
	"context"
	"github.com/bradleyshawkins/property"
)

type ResidentRepository interface {
	GetResident(ctx context.Context, residentID *property.ResidentID) (*property.Resident, error)
	SaveResident(ctx context.Context, resident *property.Resident) error
}

type PropertyRepository interface {
	GetProperty(ctx context.Context, propertyID *property.PropertyID) (*property.Property, error)
	SaveProperty(ctx context.Context, property *property.Property) error
}
