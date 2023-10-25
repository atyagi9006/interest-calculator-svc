package start

import (
	"fmt"
	"log"

	"github.com/atyagi9006/interest-calculator-svc/app/services/interest-cal-api/handlers"
	"github.com/atyagi9006/interest-calculator-svc/business/data"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const (
	collection       = "simpleinterest"
	url              = "mongodb://username:password@localhost:27017/fiber"
	selectionTimeout = 5
	mongoDBName      = "books"
)

func Run(build string) {

	repoConn := &data.RepoConn{
		MongoURL:         url,
		SelectionTimeOut: selectionTimeout,
		MongoDBName:      mongoDBName,
	}
	db, cancel, err := repoConn.NewRepoConn()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	fmt.Println("Database connection success!")

	app := fiber.New()
	app.Use(cors.New())
	handlers.V1(app, db, collection)
	defer cancel()

	err = app.Listen(":3500")
	if err != nil {
		log.Fatal(err)
	}
}
