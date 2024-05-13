package model

func MapPostTagInputToPostTag(a PostTagInput, e PostTag) PostTag { 
 if a.PostId != nil { 
	e.PostId = *a.PostId 
 }
 if a.TagContent != nil { 
	e.TagContent = *a.TagContent 
 }
  return e
}
