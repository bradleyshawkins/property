package property

func MoveIn(resident *Resident, property *Property) error {
	err := property.AddResident(resident)
	if err != nil {
		return err
	}

	err = resident.MoveIn(property)
	if err != nil {
		return err
	}

	return nil
}
