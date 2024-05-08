package model 


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