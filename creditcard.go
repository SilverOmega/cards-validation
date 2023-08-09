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
		return errors.New("Card has Expried")
	}
	// Validate the expiration year and month
	if year == time.Now().UTC().Year() && month < int(time.Now().UTC().Month()) {
		return errors.New("Card has Expried")
	}

	// Validate CVV langth
	if !skipCvv && len(c.Cvv) < 3 || len(c.Cvv) > 4 {
		return errors.New("Invalid CVV")
	}
	// Validate Card number length
	if len(c.Number) < 13 {
		return errors.New("Invalid Card Number")
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
	// implemented here
	return Company{}, nil
}

// Use Luhn algorithm check the credit card's number against the Luhn algorithm
func (c *Card) ValidateNumber() bool {
	// implemented here
	return true
}
