package property

import (
	uuid "github.com/satori/go.uuid"
)

type ResidentID struct {
	uuid.UUID
}

type Resident struct {
	ID       ResidentID
	Address  Address
	property *Property
}

func (r *Resident) MoveIn(property *Property) error {
	if r.hasResidence() {
		return ErrorHasResidence
	}

	r.moveIn(property)

	return nil
}

func (r *Resident) hasResidence() bool {
	return r.property != nil
}

func (r *Resident) moveIn(property *Property) {
	r.property = property
}
