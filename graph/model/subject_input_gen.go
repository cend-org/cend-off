package model

func MapSubjectInputToSubject(a SubjectInput, e Subject) Subject { 
 if a.EducationLevelId != nil { 
	e.EducationLevelId = *a.EducationLevelId 
 }
 if a.Name != nil { 
	e.Name = *a.Name 
 }
  return e
}
