package account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (vault *Vault) DeleteAccount(u string) (Account, bool) {
	var accounts Account
	for i, account := range vault.Accounts {
		if strings.Contains(account.Url, u) {
			accounts = account
			vault.Accounts = append(vault.Accounts[:i], vault.Accounts[i+1:]...)
			vault.save()
			return accounts, true
		}
	}
	return accounts, false
}

func (vault *Vault) FindAccount(url string) []Account {
	// findAcc := map[string]string{}
	var accounts []Account
	for _, value := range vault.Accounts {
		if strings.Contains(value.Url, url) {
			accounts = append(accounts, value)
		}
	}
	return accounts
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json") // Читаем файл, если нету то возвращаем пустой
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault) // Если файл сущестует то мы парсим json файл и возвращаем уже с распарсенный vault
	if err != nil {
		color.Red(err.Error())
	}
	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc) // Добавляем созданный аккаунт
	vault.save()
}

func (vault *Vault) ToByte() ([]byte, error) {
	content, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToByte() // Парсим данные из созданного аккаунта
	if err != nil {
		color.Red(err.Error())
	}
	files.WriteFile(data, "data.json") // Добавляем распарсенные данные
}
