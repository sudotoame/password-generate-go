package main

import (
	"fmt"

	"demo/password/account"
	"demo/password/files"
	"demo/password/output"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault(files.NewJSONDB("data.json")) // Парсим json файл каждый раз чтобы работать с актуальной версией файла
	// vault := account.NewVault(cloud.NewCloudDb("data.json")) // Парсим json файл каждый раз чтобы работать с актуальной версией файла
	fmt.Println("__Менеджер паролей__")
breakMenu:
	for {
		switch promtData([]string{
			"1. Создать аккаунт",
			"2. Поиск аккаунта",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		}) {
		case "1":
			createAccount(vault)
		case "2":
			searchAccount(vault)
		case "3":
			deleteAccount(vault)
		case "4":
			break breakMenu
		default:
			fmt.Println("Выберите пункт 1-4")
		}
	}
}

func searchAccount(vault *account.VaultWithDB) {
	url := promtData([]string{"Введите url для поиска"})
	accounts := vault.FindAccount(url)
	if len(accounts) == 0 {
		output.PrintErrorSwitch("Аккаунтов не найдено!")
	}
	fmt.Printf("Найдено записей: %d\n", len(accounts))
	for _, value := range accounts {
		// fmt.Printf("login: %s, password: %s\n", key, value)
		value.Output()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	url := promtData([]string{"Введите URL для удаления"})
	account, isDelete := vault.DeleteAccount(url)
	if !isDelete {
		output.PrintErrorSwitch("Аккаунт не найден")
	} else {
		color.Green("Аккаунт удален:")
		account.Output()
	}
}

func createAccount(vault *account.VaultWithDB) {
	login := promtData([]string{"Введите логин"})
	password := promtData([]string{"Введите пароль"})
	url := promtData([]string{"Введите url"})
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintErrorSwitch(err)
		return
	}
	vault.AddAccount(*myAccount) // Добавляем myAccount, парсим и пушим в файл
}

func promtData[T any](message []T) string {
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
