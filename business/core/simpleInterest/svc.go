package simpleinterest

import (
	"context"

	"github.com/atyagi9006/interest-calculator-svc/business/core/entities"
	"github.com/atyagi9006/interest-calculator-svc/business/core/simpleInterest/repo"
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
	siSVC.calculateSI(si)
	_, err := siSVC.repo.Create(ctx, si)
	if err != nil {
		return err
	}

	return nil

}

func (siSVC *SimpleinterestSVC) GetSimpleInterest(ctx context.Context, id string) (*entities.SimpleInterest, error) {
	val, err := siSVC.repo.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return val, nil

}

func (siSVC *SimpleinterestSVC) DeleteSimpleInterest(ctx context.Context, id string) error {

	err := siSVC.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil

}
func (siSVC *SimpleinterestSVC) calculateSI(si *entities.SimpleInterest) {
	si.InterestAmount = (si.Princpal * si.ROI * si.TimePeriod) / 100
	si.FinalAmount = si.Princpal + si.InterestAmount
}
