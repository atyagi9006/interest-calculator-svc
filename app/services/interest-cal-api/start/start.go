package start

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/atyagi9006/interest-calculator-svc/app/services/interest-cal-api/handlers"
	"github.com/atyagi9006/interest-calculator-svc/business/data"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	//TODO make these configurable
	collection       = "simpleinterest"
	url              = "mongodb://sam:sam@mongodb-container:27017"
	selectionTimeout = 500
	mongoDBName      = "intereststore"
	appPort          = "3500"
)

func Run(svcBuild string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := setupDB(ctx)
	db := client.Database(mongoDBName)
	defer client.Disconnect(ctx)

	app := setupApp(svcBuild, db)
	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":" + appPort); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")

	log.Println("Fiber was successful shutdown.")

}

func setupApp(svcBuild string, db *mongo.Database) *fiber.App {
	app := fiber.New(
		fiber.Config{
			AppName: svcBuild,
		},
	)
	app.Use(cors.New())
	api := app.Group("/api")
	handlers.V1(api, db, collection)
	//if required we can extend  api v2 api version
	return app
}

func setupDB(ctx context.Context) *mongo.Client {
	repoConn := &data.RepoConn{
		MongoURL:         url,
		SelectionTimeOut: selectionTimeout,
		MongoDBName:      mongoDBName,
	}
	client, err := repoConn.NewRepoConn(ctx)
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("Database ping failed $s", err)
	}
	log.Println("Database connection success!")
	return client
}
