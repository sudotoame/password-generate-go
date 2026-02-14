package account

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"demo/password/encrypter"
	"demo/password/output"

	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type DB interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDB struct {
	Vault
	db  DB
	enc encrypter.Encrypter
}

func NewVault(db DB, enc encrypter.Encrypter) *VaultWithDB {
	file, err := db.Read() // Читаем файл, если нету то возвращаем пустой
	if err != nil {
		output.PrintErrorSwitch(err)
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	data := enc.Decrypt(file)
	var vault Vault
	err = json.Unmarshal(data, &vault) // Если файл сущестует то мы парсим json файл и возвращаем уже с распарсенный vault
	if err != nil {
		output.PrintErrorSwitch(err)
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	color.Green("Найдено аккаунтов: %d", len(vault.Accounts))
	return &VaultWithDB{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDB) DeleteAccount(u string) (Account, bool) {
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

func (vault *VaultWithDB) FindAccount(str string, checker func(Account, string) bool) []Account {
	// findAcc := map[string]string{}
	var accounts []Account
	for _, value := range vault.Accounts {
		if checker(value, str) {
			accounts = append(accounts, value)
		}
	}
	return accounts
}

func (vault *VaultWithDB) AddAccount(acc Account) {
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

func (vault *VaultWithDB) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToByte() // Парсим данные из созданного аккаунта
	if err != nil {
		color.Red(err.Error())
	}
	encData := vault.enc.Encrypt(data)
	fmt.Printf("encData len=%d, first 20 bytes: %x\n", len(encData), encData[:20])
	vault.db.Write(encData)
}
