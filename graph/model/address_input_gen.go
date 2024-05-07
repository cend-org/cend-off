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
