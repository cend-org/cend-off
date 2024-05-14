package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type CalendarPlanningCollector struct {} 
func (c *CalendarPlanningCollector) StartDateTime(a *time.Time) (r CalendarPlanning, err error) { 
		err = database.Get(&r, `SELECT * FROM calendar_planning WHERE start_date_time = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *CalendarPlanningCollector) EndDateTime(a *time.Time) (r CalendarPlanning, err error) { 
		err = database.Get(&r, `SELECT * FROM calendar_planning WHERE end_date_time = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *CalendarPlanningCollector) Description(a *string) (r CalendarPlanning, err error) { 
		err = database.Get(&r, `SELECT * FROM calendar_planning WHERE description = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *CalendarPlanningCollector) _b(){ 
		_ = time.Now()
}

func MapCalendarPlanningInputToCalendarPlanning(a CalendarPlanningInput, e CalendarPlanning) CalendarPlanning { 
 if a.StartDateTime != nil { 
	e.StartDateTime = *a.StartDateTime 
 }
 if a.EndDateTime != nil { 
	e.EndDateTime = *a.EndDateTime 
 }
 if a.Description != nil { 
	e.Description = *a.Description 
 }
  return e
}
