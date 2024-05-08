package model 


 /* PostInput */ 



func MapPostInputToPost(input PostInput, existing Post) Post { 
 if input.PublisherId != nil { 
	existing.PublisherId = *input.PublisherId 
 } 
 
 if input.Content != nil { 
	existing.Content = *input.Content 
 } 
 
 if input.ExpirationDate != nil { 
	existing.ExpirationDate = *input.ExpirationDate 
 } 
 
  return existing 
}