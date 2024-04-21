package model

/* AssetInput */

func MapAssetInputToAsset(input AssetInput, existing Asset) Asset {
	if input.Description != nil {
		existing.Description = *input.Description
	}

	return existing
}

/* PasswordInput */

func MapPasswordInputToPassword(input PasswordInput, existing Password) Password {
	if input.Hash != nil {
		existing.Hash = *input.Hash
	}

	return existing
}

/* SchoolInput */

func MapSchoolInputToSchool(input SchoolInput, existing School) School {
	if input.Name != nil {
		existing.Name = *input.Name
	}

	return existing
}

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

/* UserInput */

func MapUserInputToUser(input UserInput, existing User) User {
	if input.Name != nil {
		existing.Name = *input.Name
	}

	if input.LastName != nil {
		existing.LastName = *input.LastName
	}

	if input.Email != nil {
		existing.Email = *input.Email
	}

	return existing
}
