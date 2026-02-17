package main

import (
	"fmt"
	"log"
)

type App struct {
	*UserHandler
	*TemplateRender
}

func main() {

	db, err := connectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return
	}
	defer db.Close()
	userRepo := NewUserRepostory(db)
	userService := NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	app := &App{
		UserHandler:    userHandler,
		TemplateRender: NewTemplateRender(false, "templates"),
	}
	err = userRepo.CreateTable()
	if err != nil {
		log.Fatal("Failed to create users table:", err)
		return
	}
	err = app.serve()
	if err != nil {
		panic(err)
	}

	fmt.Println("Server is running on http://localhost:8080")

}
