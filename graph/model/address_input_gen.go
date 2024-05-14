package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type AddressCollector struct {} 
func (c *AddressCollector) Country(a *string) (r Address, err error) { 
		err = database.Get(&r, `SELECT * FROM address WHERE country = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *AddressCollector) City(a *string) (r Address, err error) { 
		err = database.Get(&r, `SELECT * FROM address WHERE city = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *AddressCollector) Latitude(a *float64) (r Address, err error) { 
		err = database.Get(&r, `SELECT * FROM address WHERE latitude = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *AddressCollector) Longitude(a *float64) (r Address, err error) { 
		err = database.Get(&r, `SELECT * FROM address WHERE longitude = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *AddressCollector) Street(a *string) (r Address, err error) { 
		err = database.Get(&r, `SELECT * FROM address WHERE street = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *AddressCollector) FullAddress(a *string) (r Address, err error) { 
		err = database.Get(&r, `SELECT * FROM address WHERE full_address = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *AddressCollector) _b(){ 
		_ = time.Now()
}

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
