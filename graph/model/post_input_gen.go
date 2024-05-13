package model

func MapPostInputToPost(a PostInput, e Post) Post { 
 if a.PublisherId != nil { 
	e.PublisherId = *a.PublisherId 
 }
 if a.Content != nil { 
	e.Content = *a.Content 
 }
 if a.ExpirationDate != nil { 
	e.ExpirationDate = *a.ExpirationDate 
 }
  return e
}
