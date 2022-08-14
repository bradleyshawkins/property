package moving

import (
	"context"
	"github.com/bradleyshawkins/property"
)

type MoveIn struct {
	propertyRepository PropertyRepository
	residentRepository ResidentRepository
}

func NewMoveIn(p PropertyRepository, r ResidentRepository) *MoveIn {
	return &MoveIn{
		propertyRepository: p,
		residentRepository: r,
	}
}

func (m *MoveIn) MoveIn(ctx context.Context, residentID *property.ResidentID, propertyID *property.PropertyID) error {
	resident, prop, err := m.getResidentAndProperty(ctx, residentID, propertyID)
	if err != nil {
		return err
	}

	err = property.MoveIn(resident, prop)
	if err != nil {
		return err
	}

	err = m.saveResidentAndProperty(ctx, resident, prop)
	if err != nil {
		return err
	}

	return nil
}

func (m *MoveIn) getResidentAndProperty(ctx context.Context, residentID *property.ResidentID, propertyID *property.PropertyID) (*property.Resident, *property.Property, error) {
	resident, err := m.residentRepository.GetResident(ctx, residentID)
	if err != nil {
		return nil, nil, err
	}

	prop, err := m.propertyRepository.GetProperty(ctx, propertyID)
	if err != nil {
		return nil, nil, err
	}

	return resident, prop, nil
}

func (m *MoveIn) saveResidentAndProperty(ctx context.Context, resident *property.Resident, prop *property.Property) error {
	err := m.residentRepository.SaveResident(ctx, resident)
	if err != nil {
		return err
	}

	err = m.propertyRepository.SaveProperty(ctx, prop)
	if err != nil {
		return err
	}
	return nil
}
