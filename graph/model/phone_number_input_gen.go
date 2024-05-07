package model

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
