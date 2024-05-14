package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type ContractCollector struct {} 
func (c *ContractCollector) TutorId(a *int) (r Contract, err error) { 
		err = database.Get(&r, `SELECT * FROM contract WHERE tutor_id = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *ContractCollector) ParentId(a *int) (r Contract, err error) { 
		err = database.Get(&r, `SELECT * FROM contract WHERE parent_id = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *ContractCollector) StudentId(a *int) (r Contract, err error) { 
		err = database.Get(&r, `SELECT * FROM contract WHERE student_id = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *ContractCollector) StartDate(a *time.Time) (r Contract, err error) { 
		err = database.Get(&r, `SELECT * FROM contract WHERE start_date = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *ContractCollector) EndDate(a *time.Time) (r Contract, err error) { 
		err = database.Get(&r, `SELECT * FROM contract WHERE end_date = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *ContractCollector) PaymentType(a *int) (r Contract, err error) { 
		err = database.Get(&r, `SELECT * FROM contract WHERE payment_type = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *ContractCollector) SalaryValue(a *float64) (r Contract, err error) { 
		err = database.Get(&r, `SELECT * FROM contract WHERE salary_value = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *ContractCollector) PaymentMethod(a *int) (r Contract, err error) { 
		err = database.Get(&r, `SELECT * FROM contract WHERE payment_method = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *ContractCollector) _b(){ 
		_ = time.Now()
}

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
