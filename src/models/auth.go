package models

import "errors"

type LoginData struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (l *LoginData) Validate() error {
	if l.User == "" || l.Password == "" {
		return errors.New("invalid user or password")
	}
	return nil
}