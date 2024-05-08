package model 


 /* PostTagInput */ 



func MapPostTagInputToPostTag(input PostTagInput, existing PostTag) PostTag { 
 if input.PostId != nil { 
	existing.PostId = *input.PostId 
 } 
 
 if input.TagContent != nil { 
	existing.TagContent = *input.TagContent 
 } 
 
  return existing 
}