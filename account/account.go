// Package account работает с генерацией, выводом пароля
package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var lettersRune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+[]{}|\\;:'\",.<>/?~`")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) ToByte() ([]byte, error) {
	content, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (acc *Account) generatePassword(n int) {
	newSlice := make([]rune, n) // n длинна пароля
	for i := range newSlice {
		newSlice[i] = lettersRune[rand.IntN(len(lettersRune))]
	}
	acc.Password = string(newSlice)
}

func (acc *Account) Output() {
	fmt.Println("URL:", color.CyanString(acc.Url))
	fmt.Println("login:", color.CyanString(acc.Login))
	fmt.Println("password:", color.CyanString(acc.Password))
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     login,
		Password:  password,
		Url:       urlString,
	}
	if password == "" {
		newAcc.generatePassword(15)
	}
	return newAcc, nil
}
