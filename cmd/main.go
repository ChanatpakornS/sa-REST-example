package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ChanatpakornS/sa-REST-example/internal/handlers"
	model "github.com/ChanatpakornS/sa-REST-example/internal/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "restdb"
	password = "restdb"
	dbname   = "restdb"
)

func main() {
	app := fiber.New()
	db := setUpDatabase(host, port, user, password, dbname)

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("concerts/:id", func(c fiber.Ctx) error {
		return handlers.GetConcert(c, db)
	})

	app.Post("/concerts", func(c fiber.Ctx) error {
		return handlers.CreateConcert(c, db)
	})

	app.Put("/concerts/:id", func(c fiber.Ctx) error {
		return handlers.UpdateConcert(c, db)
	})

	app.Delete("/concerts/:id", func(c fiber.Ctx) error {
		return handlers.DeleteConcert(c, db)
	})

	log.Fatal(app.Listen(":3000"))
}

func setUpDatabase(host string, port int32, user string, password string, dbname string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&model.Concert{})

	return db
}
