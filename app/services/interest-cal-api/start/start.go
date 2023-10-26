package start

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/atyagi9006/interest-calculator-svc/app/services/interest-cal-api/handlers"
	"github.com/atyagi9006/interest-calculator-svc/business/data"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
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

	defer cancel()

	app := setupApp(db)
	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":3500"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")

}

func setupApp(db *mongo.Database) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	handlers.V1(app, db, collection)
	return app
}
