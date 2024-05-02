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
	if input.UserId != nil {
		existing.UserId = *input.UserId
	}

	if input.AuthorId != nil {
		existing.AuthorId = *input.AuthorId
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

	if input.ResourceNumber != nil {
		existing.ResourceNumber = *input.ResourceNumber
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

/* PostInput */

func MapPostInputToPost(input PostInput, existing Post) Post {
	if input.PublisherId != nil {
		existing.PublisherId = *input.PublisherId
	}

	if input.Content != nil {
		existing.Content = *input.Content
	}

	if input.ExpirationDate != nil {
		existing.ExpirationDate = *input.ExpirationDate
	}

	return existing
}

/* PostTagInput */

func MapPostTagInputToPostTag(input PostTagInput, existing PostTag) PostTag {
	if input.PostId != nil {
		existing.PostId = *input.PostId
	}

	if input.TagContent != nil {
		existing.TagContent = *input.TagContent
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
	if input.EducationLevelId != nil {
		existing.EducationLevelId = *input.EducationLevelId
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
