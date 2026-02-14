package main

import (
	"fmt"
	"os"
	"strings"

	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

const fileName = "data.vault"

var menu = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": searchAccountByUrl,
	"3": searchAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Поиск аккаунта по URL",
	"3. Поиск аккаунта по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите вариант",
}

func main() {

	fmt.Println("__Менеджер паролей__")
	err := godotenv.Load()
	if err != nil {
		output.PrintErrorSwitch("Error: fail load .env file")
	}
	err = os.Setenv("MY_VAR", "value")
	if err != nil {
		output.PrintErrorSwitch(err)
	}
	fmt.Println(os.Getenv("MY_VAR"))
	vault := account.NewVault(files.NewJSONDB(fileName), *encrypter.NewEncrypter()) // Парсим json файл каждый раз чтобы работать с актуальной версией файла
	// vault := account.NewVault(cloud.NewCloudDb("data.json")) // Парсим json файл каждый раз чтобы работать с актуальной версией файла

breakMenu:
	for {
		variant := promtData(menuVariants...)
		if menuFunc := menu[variant]; menuFunc != nil {
			menuFunc(vault)
		} else {
			break breakMenu
		}
	}
}

func searchAccountByLogin(vault *account.VaultWithDB) {
	login := promtData("Введите логин для поиска")
	accounts := vault.FindAccount(login, func(a account.Account, s string) bool {
		return strings.Contains(a.Login, s)
	})
	ouptupResult(&accounts)
}

func searchAccountByUrl(vault *account.VaultWithDB) {
	url := promtData("Введите url для поиска")
	accounts := vault.FindAccount(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	ouptupResult(&accounts)
}

func ouptupResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		output.PrintErrorSwitch("Аккаунтов не найдено!")
	}
	fmt.Printf("Найдено записей: %d\n", len(*accounts))
	for _, value := range *accounts {
		// fmt.Printf("login: %s, password: %s\n", key, value)
		value.Output()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	url := promtData("Введите URL для удаления")
	account, isDelete := vault.DeleteAccount(url)
	if !isDelete {
		output.PrintErrorSwitch("Аккаунт не найден")
	} else {
		color.Green("Аккаунт удален:")
		account.Output()
	}
}

func createAccount(vault *account.VaultWithDB) {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите url")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintErrorSwitch(err)
		return
	}
	vault.AddAccount(*myAccount) // Добавляем myAccount, парсим и пушим в файл
}

func promtData(message ...string) string {
	for i, v := range message {
		if i == len(message)-1 {
			fmt.Print(v, ": ")
		} else {
			fmt.Println(v)
		}
	}
	var res string
	_, err := fmt.Scanln(&res)
	if err != nil {
		output.PrintErrorSwitch(err)
		return ""
	}
	return res
}
