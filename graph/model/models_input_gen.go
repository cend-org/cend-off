package model

/* AddressInput */

func MapAddressInputToAddress(input AddressInput, existing Address) Address {
	if input.Country != nil {
		existing.Country = *input.Country
	}

	if input.City != nil {
		existing.City = *input.City
	}

	if input.Latitude != nil {
		existing.Latitude = *input.Latitude
	}

	if input.Longitude != nil {
		existing.Longitude = *input.Longitude
	}

	if input.Street != nil {
		existing.Street = *input.Street
	}

	if input.FullAddress != nil {
		existing.FullAddress = *input.FullAddress
	}

	return existing
}

/* AssetInput */

func MapAssetInputToAsset(input AssetInput, existing Asset) Asset {
	if input.Description != nil {
		existing.Description = *input.Description
	}

	return existing
}

/* CalendarPlanningInput */

func MapCalendarPlanningInputToCalendarPlanning(input CalendarPlanningInput, existing CalendarPlanning) CalendarPlanning {
	if input.StartDateTime != nil {
		existing.StartDateTime = *input.StartDateTime
	}

	if input.EndDateTime != nil {
		existing.EndDateTime = *input.EndDateTime
	}

	if input.Description != nil {
		existing.Description = *input.Description
	}

	return existing
}

/* MarkInput */

func MapMarkInputToMark(input MarkInput, existing Mark) Mark {
	if input.UserID != nil {
		existing.UserID = *input.UserID
	}

	if input.AuthorID != nil {
		existing.AuthorID = *input.AuthorID
	}

	if input.AuthorComment != nil {
		existing.AuthorComment = *input.AuthorComment
	}

	if input.AuthorMark != nil {
		existing.AuthorMark = *input.AuthorMark
	}

	return existing
}

/* MessageInput */

func MapMessageInputToMessage(input MessageInput, existing Message) Message {
	if input.ResourceType != nil {
		existing.ResourceType = *input.ResourceType
	}

	if input.ResourceValue != nil {
		existing.ResourceValue = *input.ResourceValue
	}

	if input.ResourceLabel != nil {
		existing.ResourceLabel = *input.ResourceLabel
	}

	if input.ResourceLanguage != nil {
		existing.ResourceLanguage = *input.ResourceLanguage
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

/* PhoneNumberInput */

func MapPhoneNumberInputToPhoneNumber(input PhoneNumberInput, existing PhoneNumber) PhoneNumber {
	if input.MobilePhoneNumber != nil {
		existing.MobilePhoneNumber = *input.MobilePhoneNumber
	}

	if input.IsUrgency != nil {
		existing.IsUrgency = *input.IsUrgency
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

/* SubjectInput */

func MapSubjectInputToSubject(input SubjectInput, existing Subject) Subject {
	if input.EducationLevelID != nil {
		existing.EducationLevelID = *input.EducationLevelID
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
