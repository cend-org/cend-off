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

	if input.Description != nil {
		existing.Description = *input.Description
	}

	if input.CoverText != nil {
		existing.CoverText = *input.CoverText
	}

	if input.Profile != nil {
		existing.Profile = *input.Profile
	}

	if input.ExperienceDetail != nil {
		existing.ExperienceDetail = *input.ExperienceDetail
	}

	if input.AdditionalDescription != nil {
		existing.AdditionalDescription = *input.AdditionalDescription
	}

	if input.AddOnTitle != nil {
		existing.AddOnTitle = *input.AddOnTitle
	}

	return existing
}
