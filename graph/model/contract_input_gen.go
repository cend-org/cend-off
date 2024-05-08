package model 


 /* ContractInput */ 



func MapContractInputToContract(input ContractInput, existing Contract) Contract { 
 if input.TutorId != nil { 
	existing.TutorId = *input.TutorId 
 } 
 
 if input.ParentId != nil { 
	existing.ParentId = *input.ParentId 
 } 
 
 if input.StudentId != nil { 
	existing.StudentId = *input.StudentId 
 } 
 
 if input.StartDate != nil { 
	existing.StartDate = *input.StartDate 
 } 
 
 if input.EndDate != nil { 
	existing.EndDate = *input.EndDate 
 } 
 
 if input.PaymentType != nil { 
	existing.PaymentType = *input.PaymentType 
 } 
 
 if input.SalaryValue != nil { 
	existing.SalaryValue = *input.SalaryValue 
 } 
 
 if input.PaymentMethod != nil { 
	existing.PaymentMethod = *input.PaymentMethod 
 } 
 
  return existing 
}