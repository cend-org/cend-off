package model 


 /* MarkInput */ 



func MapMarkInputToMark(input MarkInput, existing Mark) Mark { 
 if input.UserId != nil { 
	existing.UserId = *input.UserId 
 } 
 
 if input.AuthorId != nil { 
	existing.AuthorId = *input.AuthorId 
 } 
 
 if input.AuthorComment != nil { 
	existing.AuthorComment = *input.AuthorComment 
 } 
 
 if input.AuthorMark != nil { 
	existing.AuthorMark = *input.AuthorMark 
 } 
 
  return existing 
}