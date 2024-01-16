package main

import (
	"fmt"
	"sql_go/pkg/config"
	"sql_go/pkg/database"
	"sql_go/pkg/repository"
	"sql_go/pkg/service"
)

func main() {

	config := config.NewConfig()

	db := database.NewDatabase(config.DatabaseURL)

	repo := repository.NewRepository(db)

	svc := service.NewService(repo)

	for {
		var choice int
		fmt.Println("1. Логин")
		fmt.Println("2. Забыл пароль")
		fmt.Println("3. Забыл пароль и получил код")
		fmt.Println("4. Получить код")
		fmt.Println("5. Выйти")

		fmt.Scan(&choice)

		if choice == 1 {
			svc.Login()
		} else if choice == 2 {
			svc.ForgotPassword()
		} else if choice == 3 {
			svc.ChangePassword()
		} else if choice == 4 {
			svc.GetCodeByUsername()
		} else if choice == 5 {
			return
		}
	}
}

// Я забыл пароль, и отправляю ему username
// Успешно, позвоните по номеру 544 и получите номер

// Получает username, code
// Введите новый пароль
// Успешно

// Неправильный код
