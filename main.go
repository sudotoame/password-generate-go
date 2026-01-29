package main

import (
	"fmt"

	"demo/password/account"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault() // Парсим json файл каждый раз чтобы работать с актуальной версией файла
	fmt.Println("__Менеджер паролей__")
breakMenu:
	for {
		switch getMenu() {
		case 1:
			createAccount(vault)
		case 2:
			searchAccount(vault)
		case 3:
			deleteAccount(vault)
		case 4:
			break breakMenu
		default:
			fmt.Println("Выберите пункт 1-4")
		}
	}
}

func getMenu() int {
	var ch int
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Поиск аккаунта")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Print("Выберите вариант: ")
	fmt.Scan(&ch)
	return ch
}

func searchAccount(vault *account.Vault) {
	url := promtData("Введите url для поиска")
	accounts := vault.FindAccount(url)
	if len(accounts) == 0 {
		color.Green("Аккаунтов не найдено!")
	}
	fmt.Printf("Найдено записей: %d\n", len(accounts))
	for _, value := range accounts {
		// fmt.Printf("login: %s, password: %s\n", key, value)
		value.Output()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promtData("Введите URL для удаления")
	account, isDelete := vault.DeleteAccount(url)
	if !isDelete {
		color.Red("Не найдено")
	} else {
		color.Green("Аккаунт удален:")
		account.Output()
	}
}

func createAccount(vault *account.Vault) {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите url")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault.AddAccount(*myAccount) // Добавляем myAccount, парсим и пушим в файл
}

func promtData(message string) string {
	fmt.Print(message + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
