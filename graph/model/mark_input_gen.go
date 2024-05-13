package model

func MapMarkInputToMark(a MarkInput, e Mark) Mark { 
 if a.UserId != nil { 
	e.UserId = *a.UserId 
 }
 if a.AuthorId != nil { 
	e.AuthorId = *a.AuthorId 
 }
 if a.AuthorComment != nil { 
	e.AuthorComment = *a.AuthorComment 
 }
 if a.AuthorMark != nil { 
	e.AuthorMark = *a.AuthorMark 
 }
  return e
}
