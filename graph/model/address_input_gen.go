package model

func MapAddressInputToAddress(a AddressInput, e Address) Address { 
 if a.Country != nil { 
	e.Country = *a.Country 
 }
 if a.City != nil { 
	e.City = *a.City 
 }
 if a.Latitude != nil { 
	e.Latitude = *a.Latitude 
 }
 if a.Longitude != nil { 
	e.Longitude = *a.Longitude 
 }
 if a.Street != nil { 
	e.Street = *a.Street 
 }
 if a.FullAddress != nil { 
	e.FullAddress = *a.FullAddress 
 }
  return e
}
