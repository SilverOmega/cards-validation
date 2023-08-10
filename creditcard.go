package cards_validation

import (
	"errors"
	_ "errors"
	"strconv"
	_ "strconv"
	"time"
	_ "time"
)

type Card struct {
	Number  string
	Cvv     string
	Month   string
	Years   string
	Company Company
}

type Company struct {
	Code string
	Name string
}
type digits [6]int

func (d *digits) At(i int) int {
	return d[i-1]
}

// LastFour returns the last four digits of the credit card's number
func (c *Card) LastFour() (string, error) {
	if len(c.Number) < 4 {
		return "", errors.New("Credit card number is not long enough")
	}
	return c.Number[len(c.Number)-4 : len(c.Number)], nil
}

// LastFourDigits as an alias for LastFour
func (c *Card) LastFourDigits() (string, error) {
	return c.LastFour()
}

func (c *Card) Wipe() {
	c.Cvv, c.Number, c.Month, c.Years = "000", "0000000000000", "01", "2026"
}

func (c *Card) Validate(allowTestNumbers ...bool) error {
	if len(allowTestNumbers) > 0 {
		return c.validate(false, allowTestNumbers[0])
	} else {
		return c.validate(false)
	}
}

func (c *Card) ValidateCvv(allowTestNumbers ...bool) error {
	if len(allowTestNumbers) > 0 {
		return c.validate(false, allowTestNumbers[0])
	} else {
		return c.validate(false)
	}
}

func (c *Card) validate(skipCvv bool, allowTestNumbers ...bool) error {
	var year, month int
	var err error
	// Format the expiration year
	if len(c.Years) < 3 {
		if year, err = strconv.Atoi(strconv.Itoa(time.Now().UTC().Year())[:2] + c.Years); err != nil {
			return errors.New("Invalid Years")
		}
	} else {
		if year, err = strconv.Atoi(c.Years); err != nil {
			return errors.New("Invalid Years")
		}
	}
	// Validate expiration month
	if month, err = strconv.Atoi(c.Month); err != nil || month < 1 || 12 < month {
		return errors.New("Invalid Month")
	}
	// Validate the expiration year
	if year < time.Now().UTC().Year() {
		return errors.New("Credit Card has Expried")
	}
	// Validate the expiration year and month
	if year == time.Now().UTC().Year() && month < int(time.Now().UTC().Month()) {
		return errors.New("Credit Card has Expried")
	}

	// Validate CVV langth
	if !skipCvv && len(c.Cvv) < 3 || len(c.Cvv) > 4 {
		return errors.New("Invalid CVV")
	}
	// Validate Card number length
	if len(c.Number) < 13 {
		return errors.New("Invalid Credit Card Number")
	}

	switch c.Number {
	case "4242424242424242",
		"4012888888881881",
		"4000056655665556",
		"3566002020360505":
		if len(allowTestNumbers) > 0 && allowTestNumbers[0] {
			return nil
		}
		return errors.New("Test Numbers are not allowed")
	}

	valid := c.ValidateNumber()
	if !valid {
		return errors.New("Invalid credit card number")
	}
	return nil
}

// Brand returns an error from BrandValidate() or returns the
// credit card with it's company / issuer attached to it
func (c *Card) Bard() error {
	company, err := c.BrandValidate()
	if err != nil {
		return err
	}
	c.Company = company
	return nil
}

// BrandValidate adds/checks/verifies the credit card's company / issuer
func (c *Card) BrandValidate() (Company, error) {
	ccLen := len(c.Number)
	ccDigits := digits{}
	var err error

	for i := 0; i < 6; i++ {
		if i < ccLen {
			if ccDigits[i], err = strconv.Atoi(c.Number[:i+1]); err != nil {
				return Company{"", ""}, errors.New("Unknow Credit Card Brand")
			}
		}
	}
	switch {
	case ccDigits.At(4) == 4011 || ccDigits.At(6) == 431274 || ccDigits.At(6) == 438935 ||
		ccDigits.At(6) == 451416 || ccDigits.At(6) == 457393 || ccDigits.At(4) == 4576 ||
		ccDigits.At(6) == 457631 || ccDigits.At(6) == 457632 || ccDigits.At(6) == 504175 ||
		ccDigits.At(6) == 627780 || ccDigits.At(6) == 636297 || ccDigits.At(6) == 636368 ||
		ccDigits.At(6) == 636369 || (ccDigits.At(6) >= 506699 && ccDigits.At(6) <= 506778) ||
		(ccDigits.At(6) >= 509000 && ccDigits.At(6) <= 509999) ||
		(ccDigits.At(6) >= 650031 && ccDigits.At(6) <= 650051) ||
		(ccDigits.At(6) >= 650035 && ccDigits.At(6) <= 650033) ||
		(ccDigits.At(6) >= 650405 && ccDigits.At(6) <= 650439) ||
		(ccDigits.At(6) >= 650485 && ccDigits.At(6) <= 650538) ||
		(ccDigits.At(6) >= 650541 && ccDigits.At(6) <= 650598) ||
		(ccDigits.At(6) >= 650700 && ccDigits.At(6) <= 650718) ||
		(ccDigits.At(6) >= 650720 && ccDigits.At(6) <= 650727) ||
		(ccDigits.At(6) >= 650901 && ccDigits.At(6) <= 650920) ||
		(ccDigits.At(6) >= 651652 && ccDigits.At(6) <= 651679) ||
		(ccDigits.At(6) >= 655000 && ccDigits.At(6) <= 655019) ||
		(ccDigits.At(6) >= 655021 && ccDigits.At(6) <= 655021):
		return Company{"elo", "Elo"}, nil

	case ccDigits.At(6) >= 604201 && ccDigits.At(6) <= 604219:
		return Company{"SC", "Standard Chartered"}, nil

	case ccDigits.At(6) == 384100 || ccDigits.At(6) == 384140 || ccDigits.At(6) == 384160 ||
		ccDigits.At(6) == 606282 || ccDigits.At(6) == 637095 || ccDigits.At(4) == 637568 ||
		ccDigits.At(4) == 637599 || ccDigits.At(4) == 637609 || ccDigits.At(4) == 637612:
		return Company{"acb", "ACB"}, nil

	case ccDigits.At(2) == 34 || ccDigits.At(2) == 37:
		return Company{"vcb", "VietcomBank"}, nil

	case ccDigits.At(4) == 5610 || (ccDigits.At(6) >= 560221 && ccDigits.At(6) <= 560225):
		return Company{"abc", "ABC Card"}, nil

	case ccDigits.At(2) == 62:
		return Company{"tpb", "TienPhong Bank"}, nil

	case ccDigits.At(3) >= 300 && ccDigits.At(3) <= 305 && ccLen == 15:
		return Company{"ocb", "OCB"}, nil

	case ccDigits.At(2) >= 51 && ccDigits.At(2) <= 55:
		return Company{"mastercard", "Mastercard"}, nil

	case ccDigits.At(2) == 35:
		return Company{"jcb", "JCB"}, nil

	case ccDigits.At(1) == 4:
		return Company{"visa", "Visa"}, nil

	default:
		return Company{"", ""}, errors.New("Unknow Credit Card Brand")
	}

}

// Use Luhn algorithm check the credit card's number against the Luhn algorithm
func (c *Card) ValidateNumber() bool {
	var sum, mod int
	var alternate bool
	var err error

	// get card number length
	numberLen := len(c.Number)

	if numberLen < 13 || numberLen > 16 {
		return false
	}
	// Parse all numbers of the card into a for loop
	for i := numberLen - 1; i > -1; i-- {
		// Takes the mod, converting the current number in integer
		if mod, err = strconv.Atoi(string(c.Number[i])); err != nil {
			return false
		}
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}
		alternate = !alternate
		sum += mod
	}
	return sum%10 == 0
}
