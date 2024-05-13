package model

func MapPasswordInputToPassword(a PasswordInput, e Password) Password { 
 if a.Hash != nil { 
	e.Hash = *a.Hash 
 }
  return e
}
