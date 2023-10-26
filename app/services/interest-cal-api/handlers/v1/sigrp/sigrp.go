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
	siRoute.Get("", handler.getSimpleInterest)
	siRoute.Post("", handler.createSimpleInterest)

	// Declare routing endpoints for specific routes.
	siRoute.Get("/:SimpleInterestID", handler.getSimpleInterest)
	siRoute.Delete("/:SimpleInterestID", handler.deleteSimpleInterest)
}

func (h *SIHandler) getSimpleInterest(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Fetch parameter.

	// TODO: fix it for SIID

	// Get one SimpleInterest.
	SimpleInterest, err := h.simpleInterestSVC.GetSimpleInterest(customContext, "targetedSimpleInterestID")
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

	h.calculateSI(simpleInterest)
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

	//TODO : fix it for SIID
	// Delete one SimpleInterest.
	err := h.simpleInterestSVC.DeleteSimpleInterest(customContext, "targetedSimpleInterestID")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return 204 No Content.
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *SIHandler) calculateSI(si *entities.SimpleInterest) {
	si.InterestAmount = (si.Princpal * si.ROI * si.TimePeriod) / 100
	si.FinalAmount = si.Princpal + si.InterestAmount
}
