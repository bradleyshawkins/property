package property

type PropertyResidents map[ResidentID]*Resident

func (r PropertyResidents) AddResident(resident *Resident) error {
	if r.alreadyResident(resident) {
		return ErrorAlreadyResident
	}

	r.addResident(resident)

	return nil
}

func (r PropertyResidents) alreadyResident(resident *Resident) bool {
	_, ok := r[resident.ID]
	return ok
}

func (r PropertyResidents) addResident(resident *Resident) {
	r[resident.ID] = resident
}
