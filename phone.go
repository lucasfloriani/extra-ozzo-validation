package isbr

import (
	"regexp"

	"github.com/go-ozzo/ozzo-validation"
)

var (
	// Phone validate if a string is a brazilian phone format or not
	Phone = validation.NewStringRule(isPhone, "must be a valid phone")
	// CellPhone validate if a string is a brazilian cellphone format or not
	CellPhone = validation.NewStringRule(isCellPhone, "must be a valid cellphone")
	// PhoneOrCellPhone validate if a string is a brazilian phone/cellphone format or not
	PhoneOrCellPhone = validation.NewStringRule(isPhoneOrCellPhone, "must be a valid phone or cellphone")
)

func isPhone(value string) bool {
	rePhone := regexp.MustCompile(`^\([0-9]{2}\) [0-9]{4}-[0-9]{4}$`)
	return rePhone.MatchString(value)
}

func isCellPhone(value string) bool {
	reCellPhone := regexp.MustCompile(`^\([0-9]{2}\) [0-9]{5}-[0-9]{4}$`)
	return reCellPhone.MatchString(value)
}

func isPhoneOrCellPhone(value string) bool {
	regex := regexp.MustCompile(`^\([0-9]{2}\) [0-9]{4,5}-[0-9]{4}$`)
	return regex.MatchString(value)
}
