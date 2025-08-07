package main

import (
	"fmt"
	"main/internal/database"
	"main/internal/handlers"
	"main/internal/server"
)

func main() {

	Serv := server.ServerInit("8080", "localhost")

	db, err := database.InitBD()
	if err != nil {
		fmt.Println(err)
		return
	}
	handlers.HandlersInit(Serv.Engine, db)
	if err := Serv.Run(); err != nil {
		fmt.Println(err)
		return
	}

}

/*

апи онлайн доски объявлений мероприятий
алгоритм действий на сайте
я регистрируюсь создаю мероприятие далее эта доска появляеться у всех без обновления страницы нужно создать страницу каждой доски
*/
