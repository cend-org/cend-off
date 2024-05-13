package model

func MapUserInputToUser(a UserInput, e User) User { 
 if a.Name != nil { 
	e.Name = *a.Name 
 }
 if a.FamilyName != nil { 
	e.FamilyName = *a.FamilyName 
 }
 if a.NickName != nil { 
	e.NickName = *a.NickName 
 }
 if a.Email != nil { 
	e.Email = *a.Email 
 }
 if a.BirthDate != nil { 
	e.BirthDate = *a.BirthDate 
 }
 if a.Sex != nil { 
	e.Sex = *a.Sex 
 }
 if a.Lang != nil { 
	e.Lang = *a.Lang 
 }
 if a.Description != nil { 
	e.Description = *a.Description 
 }
 if a.CoverText != nil { 
	e.CoverText = *a.CoverText 
 }
 if a.Profile != nil { 
	e.Profile = *a.Profile 
 }
 if a.ExperienceDetail != nil { 
	e.ExperienceDetail = *a.ExperienceDetail 
 }
 if a.AdditionalDescription != nil { 
	e.AdditionalDescription = *a.AdditionalDescription 
 }
 if a.AddOnTitle != nil { 
	e.AddOnTitle = *a.AddOnTitle 
 }
  return e
}
