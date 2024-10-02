package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateIn time.Time `json:"CreateIn,omitempty"`
}

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()
	return nil
}

// validate mandatory fields
func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("the name field is mandatory")
	}

	if user.Nick == "" {
		return errors.New("the nick field is mandatory")
	}

	if user.Email == "" {
		return errors.New("the e-mail field is mandatory")
	}

	if user.Password == "" {
		return errors.New("the password field is mandatory")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

}
