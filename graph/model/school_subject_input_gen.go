package model

/* SchoolSubjectInput */

func MapSchoolSubjectInputToSchoolSubject(input SchoolSubjectInput, existing SchoolSubject) SchoolSubject {
	if input.SchoolNumber != nil {
		existing.SchoolNumber = *input.SchoolNumber
	}

	if input.Name != nil {
		existing.Name = *input.Name
	}

	return existing
}
