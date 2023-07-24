package serializer

import (
	"fmt"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ValidatePinReq struct {
	Pin string `json:"pin"`
}

func (r ValidatePinReq) ValidateRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Pin, validation.Required, validation.By(securePINs)),
	)
}

func securePINs(value interface{}) error {
	pin, ok := value.(string)
	if !ok {
		return fmt.Errorf("pin must be string")
	}
	if len(pin) < 6 {
		return fmt.Errorf("length must more than 6")
	}

	if isAllDigitDuplicate(pin) {
		return fmt.Errorf("all digit was duplicated")
	}

	if isTwoDuplicatesInOrder(pin) {
		return fmt.Errorf("2 digit was duplicated")
	}

	if isDifferenceLessThanOrEqualThreeDigits(pin) {
		return fmt.Errorf("pin must difference more than 3 digits")
	}

	if isNumberOrdered(pin) {
		return fmt.Errorf("pin must not order")
	}

	return nil
}

func isNumberOrdered(pin string) bool {
	var increasing, decreasing bool

	for i := 1; i <= len(pin)-2; i++ {
		prevInt, _ := strconv.Atoi(string(pin[i-1]))
		currInt, _ := strconv.Atoi(string(pin[i]))
		nextInt, _ := strconv.Atoi(string(pin[i+1]))

		//321
		if prevInt-1 == currInt && currInt-1 == nextInt {
			decreasing = true
			break
		}

		//123
		if prevInt+1 == currInt && currInt+1 == nextInt {
			increasing = true
			break
		}

	}

	return increasing || decreasing
}

// isTwoDuplicatesInOrder ex. 1112233
func isTwoDuplicatesInOrder(pin string) bool {
	var isDuplicateTwoDigitInOrder bool
	var count int

	for i := 1; i <= len(pin)-1; i++ {
		if pin[i] == pin[i-1] {
			count++
		}

	}

	if count > 2 {
		isDuplicateTwoDigitInOrder = true
	}

	return isDuplicateTwoDigitInOrder
}

func isAllDigitDuplicate(pin string) bool {
	var digits = make(map[rune]int)
	for _, c := range pin {
		digits[c]++
	}
	return len(digits) == 1
}

// isDifferenceLessThanThreeDigits ex. 111115
func isDifferenceLessThanOrEqualThreeDigits(pin string) bool {
	var digits = make(map[rune]int)
	for _, c := range pin {
		digits[c]++
	}
	return len(digits) < 3
}
