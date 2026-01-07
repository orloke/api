package models

import (
	"errors"
)

type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}

func (p *Password) Validate() error {
	if p.New == "" {
		return errors.New("new password is required")
	}
	if p.Current == "" {
		return errors.New("current password is required")
	}
	return nil
}
