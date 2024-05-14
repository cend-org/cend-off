package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type PhoneNumberCollector struct {} 
func (c *PhoneNumberCollector) MobilePhoneNumber(a *string) (r PhoneNumber, err error) { 
		err = database.Get(&r, `SELECT * FROM phone_number WHERE mobile_phone_number = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *PhoneNumberCollector) IsUrgency(a *bool) (r PhoneNumber, err error) { 
		err = database.Get(&r, `SELECT * FROM phone_number WHERE is_urgency = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *PhoneNumberCollector) _b(){ 
		_ = time.Now()
}

func MapPhoneNumberInputToPhoneNumber(a PhoneNumberInput, e PhoneNumber) PhoneNumber { 
 if a.MobilePhoneNumber != nil { 
	e.MobilePhoneNumber = *a.MobilePhoneNumber 
 }
 if a.IsUrgency != nil { 
	e.IsUrgency = *a.IsUrgency 
 }
  return e
}
