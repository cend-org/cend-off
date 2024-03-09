package utils

import "regexp"


/*
 CHECK IF MOBILE PHONE NUMBER IS NUMERIC AND A LENGHT OF TEN 
*/
func PhoneValidator(phone string) bool {
	re := regexp.MustCompile(`^\d{10}$`)
	return re.MatchString(phone)
}
