package cognito

import (
	"fmt"
	"regexp"
)

func validateEmail(v interface{}, k string) (ws []string, es []error) {
	value := v.(string)
	if !regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString(value) {
		es = append(es, fmt.Errorf("%q must be a valid email address", k))
	}
	return
}

func validatePassword(v interface{}, k string) (ws []string, es []error) {
	value := v.(string)
	if !regexp.MustCompile(`^(?=[^A-Z]*[A-Z])(?=(?:[^a-z]*[a-z]){3})(?=(?:[^0-9]*[0-9]){‌​2})(?=(?:[^!?@*#&$]*‌​[!?@*#&$]){1})(?!.*(‌​.)\\1{2})[A-Za-z].{7,‌​11}$`).MatchString(value) {
		es = append(es, fmt.Errorf("%q (%s) is an invalid password: %s)", k))
	}
	return
}