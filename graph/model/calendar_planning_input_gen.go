package model

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
