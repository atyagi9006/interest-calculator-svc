package simpleinterest

import (
	"context"
	"errors"
	"time"

	"github.com/atyagi9006/interest-calculator-svc/business/core/entities"
	"github.com/atyagi9006/interest-calculator-svc/business/core/simpleInterest/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	errInvalidPrincipal  = "invalid principal"
	errInvalidROI        = "invalid roi"
	errInvalidTimePeriod = "invalid timeperiod"
	errInvalidId         = "invalid input"
)

type SimpleinterestSVC struct {
	repo repo.Repository
}

func NewService(r repo.Repository) *SimpleinterestSVC {
	return &SimpleinterestSVC{
		repo: r,
	}
}

func (siSVC *SimpleinterestSVC) CreateSimpleInterest(ctx context.Context, si *entities.SimpleInterest) error {
	err := siSVC.calculateSI(si)
	if err != nil {
		return err
	}
	si.ID = primitive.NewObjectID()
	si.CreatedAt = time.Now()
	si.UpdatedAt = time.Now()
	_, err = siSVC.repo.Create(ctx, si)
	if err != nil {
		return err
	}

	return nil

}

func (siSVC *SimpleinterestSVC) GetSimpleInterest(ctx context.Context, id string) (*entities.SimpleInterest, error) {
	if id == "" {
		return nil, errors.New(errInvalidId)
	}

	val, err := siSVC.repo.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return val, nil

}

func (siSVC *SimpleinterestSVC) DeleteSimpleInterest(ctx context.Context, id string) error {
	if id == "" {
		return errors.New(errInvalidId)
	}
	err := siSVC.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil

}
func (siSVC *SimpleinterestSVC) calculateSI(si *entities.SimpleInterest) error {
	if si.Princpal <= 0.0 {
		return errors.New(errInvalidPrincipal)
	}
	if si.ROI <= 0.0 {
		return errors.New(errInvalidROI)
	}
	if si.TimePeriod <= 0.0 {
		return errors.New(errInvalidTimePeriod)
	}
	si.InterestAmount = (si.Princpal * si.ROI * si.TimePeriod) / 100
	si.FinalAmount = si.Princpal + si.InterestAmount
	return nil
}
