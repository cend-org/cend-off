package model

func MapContractTimesheetDetailInputToContractTimesheetDetail(a ContractTimesheetDetailInput, e ContractTimesheetDetail) ContractTimesheetDetail { 
 if a.ContractId != nil { 
	e.ContractId = *a.ContractId 
 }
 if a.Date != nil { 
	e.Date = *a.Date 
 }
 if a.Hours != nil { 
	e.Hours = *a.Hours 
 }
  return e
}
