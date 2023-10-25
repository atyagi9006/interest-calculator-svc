package handlers

import (
	v1SIGrp "github.com/atyagi9006/interest-calculator-svc/app/services/interest-cal-api/handlers/v1/sigrp"
	simpleinterest "github.com/atyagi9006/interest-calculator-svc/business/core/simpleInterest"
	"github.com/atyagi9006/interest-calculator-svc/business/core/simpleInterest/repo"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// v1 binds all the version 1 routes.

func V1(app *fiber.App, db *mongo.Database, collection string) {
	siCollection := db.Collection(collection)
	siRepo := repo.NewRepo(siCollection)
	// Register user management and authentication endpoints.
	siService := simpleinterest.NewService(siRepo)
	v1SIGrp.NewSIHandler(app.Group("/api/v1/si"), siService)

}
