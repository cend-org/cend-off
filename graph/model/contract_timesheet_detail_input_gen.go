package model 


 /* ContractTimesheetDetailInput */ 



func MapContractTimesheetDetailInputToContractTimesheetDetail(input ContractTimesheetDetailInput, existing ContractTimesheetDetail) ContractTimesheetDetail { 
 if input.ContractId != nil { 
	existing.ContractId = *input.ContractId 
 } 
 
 if input.Date != nil { 
	existing.Date = *input.Date 
 } 
 
 if input.Hours != nil { 
	existing.Hours = *input.Hours 
 } 
 
  return existing 
}