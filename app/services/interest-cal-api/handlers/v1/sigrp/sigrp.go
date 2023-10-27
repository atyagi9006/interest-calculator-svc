package sigrp

import (
	"context"

	"github.com/atyagi9006/interest-calculator-svc/business/core/entities"
	si "github.com/atyagi9006/interest-calculator-svc/business/core/simpleInterest"
	"github.com/gofiber/fiber/v2"
)

type SIHandler struct {
	simpleInterestSVC *si.SimpleinterestSVC
}

func NewSIHandler(siRoute fiber.Router, siSVC *si.SimpleinterestSVC) {
	// Create a handler based on our created service / use-case.
	handler := &SIHandler{
		simpleInterestSVC: siSVC,
	}

	// Declare routing endpoints for general routes.
	siRoute.Post("", handler.createSimpleInterest)

	// Declare routing endpoints for specific routes.
	siRoute.Get("/:id", handler.getSimpleInterest)
	siRoute.Delete("/:id", handler.deleteSimpleInterest)
}

func (h *SIHandler) getSimpleInterest(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Fetch parameter.
	param := c.Params("id")
	
	// Get one SimpleInterest.
	SimpleInterest, err := h.simpleInterestSVC.GetSimpleInterest(customContext,param)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   SimpleInterest,
	})
}
func (h *SIHandler) createSimpleInterest(c *fiber.Ctx) error {
	// Initialize variables.
	simpleInterest := &entities.SimpleInterest{}

	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Parse request body.
	if err := c.BodyParser(simpleInterest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Create one SimpleInterest.
	err := h.simpleInterestSVC.CreateSimpleInterest(customContext, simpleInterest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return result.
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"result":  simpleInterest,
		"status":  "success",
		"message": "SimpleInterest has been created successfully!",
	})
}

// Deletes a single SimpleInterest.
func (h *SIHandler) deleteSimpleInterest(c *fiber.Ctx) error {
	// Initialize previous SimpleInterest ID.

	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	params := c.Params("id")

	// Delete one SimpleInterest.
	err := h.simpleInterestSVC.DeleteSimpleInterest(customContext, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return 204 No Content.
	return c.SendStatus(fiber.StatusNoContent)
}
