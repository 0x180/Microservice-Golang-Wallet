package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"

	"github.com/Marlos-Rodriguez/go-postgres-wallet-back/user/handlers"
)

func routes(DB *gorm.DB) *fiber.App {
	app := fiber.New()

	handler := handlers.NewUserhandlerService(DB)

	user := app.Group("/user")

	user.Use(cors.New())

	user.Get("/:id", handler.GetUser)                     //Get Basic user info
	user.Get("/all/:id", handler.GetProfileUser)          //Get Profile User Info
	user.Get("/relation/:id/:page", handler.GetRelations) //Get relations of user
	user.Put("/:id", handler.ModifyUser)                  //Modify the user info
	user.Post("/add", handler.CreateRelation)             //Create a new relation

	return app
}