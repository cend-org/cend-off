package model 


 /* MessageInput */ 



func MapMessageInputToMessage(input MessageInput, existing Message) Message { 
 if input.ResourceType != nil { 
	existing.ResourceType = *input.ResourceType 
 } 
 
 if input.ResourceValue != nil { 
	existing.ResourceValue = *input.ResourceValue 
 } 
 
 if input.ResourceNumber != nil { 
	existing.ResourceNumber = *input.ResourceNumber 
 } 
 
 if input.ResourceLabel != nil { 
	existing.ResourceLabel = *input.ResourceLabel 
 } 
 
 if input.ResourceLanguage != nil { 
	existing.ResourceLanguage = *input.ResourceLanguage 
 } 
 
  return existing 
}