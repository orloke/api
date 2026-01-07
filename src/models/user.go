package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewUserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Nick  string `json:"nick"`
	Email string `json:"email"`
}

type CreateUserResponse struct {
	Message string          `json:"message"`
	User    NewUserResponse `json:"user"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name is required and cannot be empty")
	}

	if user.Nick == "" {
		return errors.New("nick is required and cannot be empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("the email entered is invalid")
	}

	if step == "register" && user.Password == "" {
		return errors.New("password is required and cannot be empty")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashedPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)
	}

	return nil
}
