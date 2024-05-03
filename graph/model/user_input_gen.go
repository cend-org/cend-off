package model

/* UserInput */

func MapUserInputToUser(input UserInput, existing User) User {
	if input.Name != nil {
		existing.Name = *input.Name
	}

	if input.FamilyName != nil {
		existing.FamilyName = *input.FamilyName
	}

	if input.NickName != nil {
		existing.NickName = *input.NickName
	}

	if input.Email != nil {
		existing.Email = *input.Email
	}

	if input.BirthDate != nil {
		existing.BirthDate = *input.BirthDate
	}

	if input.Sex != nil {
		existing.Sex = *input.Sex
	}

	if input.Lang != nil {
		existing.Lang = *input.Lang
	}

	return existing
}
