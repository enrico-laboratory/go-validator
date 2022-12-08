package validator

import (
	"bytes"
	"fmt"
	"regexp"
	"unicode"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Errors interface {
	Error() string
}

type ErrorMap map[string][]string

// Add adds an error message for a given form field
func (e ErrorMap) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns first error message for a field
func (e ErrorMap) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}

// Validator defines a new Validator type which contains a map of validation errors.
type Validator struct {
	ErrorMap ErrorMap
}

// New is a helper which creates a new Validator instance with an empty errors map.
func New() *Validator {
	return &Validator{ErrorMap: ErrorMap{}}
}

// Valid returns true if the errors map doesn't contain any entries.
func (v *Validator) Valid() bool {
	return len(v.ErrorMap) == 0
}

// AddError adds an error message to the map (so long as no entry already exists for the given key).
func (v *Validator) AddError(key, message string) {
	if _, exists := v.ErrorMap[key]; !exists {
		v.ErrorMap.Add(key, message)
	}
}

// Check adds an error message to the map only if a validation check is not 'ok'.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// In returns true if a specific value is in a list of strings.
func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

// Matches returns true if a string value matches a specific regexp pattern.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Unique returns true if all string values in a slice are unique.
func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}

// Error returns the Errors map as string
func (v *Validator) Error() string {
	b := new(bytes.Buffer)
	for key, value := range v.ErrorMap {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

// ValidatePassword with minimum amount characters, at least one uppercase letter, one lowercase letter and one numbe
func ValidatePassword(s string, minimumCharacters int) bool {
	characterSum := 0
	number := false
	upper := false
	special := false
	hasMinimumCharacters := false
	for _, c := range s {
		characterSum++
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
		default:
			return false
		}
	}
	hasMinimumCharacters = characterSum >= minimumCharacters
	if !number || !upper || !special || !hasMinimumCharacters {
		return false
	}
	return true
}
