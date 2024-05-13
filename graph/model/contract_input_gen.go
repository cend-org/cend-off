package model

func MapContractInputToContract(a ContractInput, e Contract) Contract { 
 if a.TutorId != nil { 
	e.TutorId = *a.TutorId 
 }
 if a.ParentId != nil { 
	e.ParentId = *a.ParentId 
 }
 if a.StudentId != nil { 
	e.StudentId = *a.StudentId 
 }
 if a.StartDate != nil { 
	e.StartDate = *a.StartDate 
 }
 if a.EndDate != nil { 
	e.EndDate = *a.EndDate 
 }
 if a.PaymentType != nil { 
	e.PaymentType = *a.PaymentType 
 }
 if a.SalaryValue != nil { 
	e.SalaryValue = *a.SalaryValue 
 }
 if a.PaymentMethod != nil { 
	e.PaymentMethod = *a.PaymentMethod 
 }
  return e
}
