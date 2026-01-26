package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var lettersRune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+[]{}|\\;:'\",.<>/?~`")

type account struct {
	password string
	login    string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func (acc *account) generatePassword(n int) {
	newSlice := make([]rune, n) // n длинна пароля
	for i := range newSlice {
		newSlice[i] = lettersRune[rand.IntN(len(lettersRune))]
	}
	acc.password = string(newSlice)
}

func (acc *accountWithTimeStamp) outputPassword() {
	fmt.Println(acc)
	fmt.Println("createdAt:", acc.createdAt)
	fmt.Println("updatedAt:", acc.updatedAt)
	fmt.Println("login:", acc.login)
	fmt.Println("password:", acc.password)
	fmt.Println("url:", acc.url)
}

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}
	if password == "" {
		newAcc.generatePassword(15)
	}
	return newAcc, nil
}

// func newAccount(login, password, urlString string) (*account, error) {
// 	if login == "" {
// 		return nil, errors.New("EMPTY_LOGIN")
// 	}
// 	_, err := url.ParseRequestURI(urlString)
// 	if err != nil {
// 		return nil, errors.New("INVALID_URL")
// 	}
// 	newAcc := &account{
// 		login:    login,
// 		password: password,
// 		url:      urlString,
// 	}
// 	if password == "" {
// 		newAcc.generatePassword(15)
// 	}
// 	return newAcc, nil
// }

func main() {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	// var password int
	// fmt.Print("Введите длину пароля: ")
	// fmt.Scan(&password)
	url := promtData("Введите url")

	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAccount.outputPassword()
}

func promtData(message string) string {
	fmt.Print(message + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
