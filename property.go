package property

import uuid "github.com/satori/go.uuid"

type PropertyID struct {
	uuid.UUID
}

type Property struct {
	id        PropertyID
	address   Address
	residents PropertyResidents
}

func (p *Property) AddResident(r *Resident) error {
	return p.residents.AddResident(r)
}
