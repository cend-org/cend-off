package model 


 /* PasswordInput */ 



func MapPasswordInputToPassword(input PasswordInput, existing Password) Password { 
 if input.Hash != nil { 
	existing.Hash = *input.Hash 
 } 
 
  return existing 
}