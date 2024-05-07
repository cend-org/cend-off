package model

/* SchoolInput */

func MapSchoolInputToSchool(input SchoolInput, existing School) School {
	if input.Name != nil {
		existing.Name = *input.Name
	}

	return existing
}
