package model

func MapMessageInputToMessage(a MessageInput, e Message) Message { 
 if a.ResourceType != nil { 
	e.ResourceType = *a.ResourceType 
 }
 if a.ResourceValue != nil { 
	e.ResourceValue = *a.ResourceValue 
 }
 if a.ResourceNumber != nil { 
	e.ResourceNumber = *a.ResourceNumber 
 }
 if a.ResourceLabel != nil { 
	e.ResourceLabel = *a.ResourceLabel 
 }
 if a.ResourceLanguage != nil { 
	e.ResourceLanguage = *a.ResourceLanguage 
 }
  return e
}
