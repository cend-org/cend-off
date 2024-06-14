package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type AppointmentCollector struct {} 
func (c *AppointmentCollector) Availability(a *time.Time) (r Appointment, err error) { 
		err = database.Get(&r, `SELECT * FROM appointment WHERE availability = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *AppointmentCollector) _b(){ 
		_ = time.Now()
}

func MapAppointmentInputToAppointment(a AppointmentInput, e Appointment) Appointment { 
 if a.Availability != nil { 
	e.Availability = *a.Availability 
 }
  return e
}
