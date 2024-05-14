package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type ContractTimesheetDetailCollector struct {} 
func (c *ContractTimesheetDetailCollector) ContractId(a *int) (r ContractTimesheetDetail, err error) { 
		err = database.Get(&r, `SELECT * FROM contract_timesheet_detail WHERE contract_id = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *ContractTimesheetDetailCollector) Date(a *time.Time) (r ContractTimesheetDetail, err error) { 
		err = database.Get(&r, `SELECT * FROM contract_timesheet_detail WHERE date = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *ContractTimesheetDetailCollector) Hours(a *float64) (r ContractTimesheetDetail, err error) { 
		err = database.Get(&r, `SELECT * FROM contract_timesheet_detail WHERE hours = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *ContractTimesheetDetailCollector) _b(){ 
		_ = time.Now()
}

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
