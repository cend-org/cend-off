package model

func MapPhoneNumberInputToPhoneNumber(a PhoneNumberInput, e PhoneNumber) PhoneNumber { 
 if a.MobilePhoneNumber != nil { 
	e.MobilePhoneNumber = *a.MobilePhoneNumber 
 }
 if a.IsUrgency != nil { 
	e.IsUrgency = *a.IsUrgency 
 }
  return e
}
