package model

/* SubjectInput */

func MapSubjectInputToSubject(input SubjectInput, existing Subject) Subject {
	if input.EducationLevelId != nil {
		existing.EducationLevelId = *input.EducationLevelId
	}

	if input.Name != nil {
		existing.Name = *input.Name
	}

	return existing
}
