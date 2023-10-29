package simpleinterest

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/atyagi9006/interest-calculator-svc/business/core/entities"
	"github.com/atyagi9006/interest-calculator-svc/business/core/simpleInterest/fakerepo/fakeMongo"
	"github.com/stretchr/testify/assert"
)

func TestCreateSimpleInterest(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"create simple interest with zero principal": func() {
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			fakeRepo.CreateReturns(nil, nil)

			ctx := context.Background()
			input := &entities.SimpleInterest{
				ID:             [12]byte{},
				Principal:       0,
				ROI:            0,
				TimePeriod:     0,
				InterestAmount: 0,
				FinalAmount:    0,
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			}

			err := siSvc.CreateSimpleInterest(ctx, input)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), errInvalidPrincipal)
		},
		"create simple interest with negitive principal": func() {
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			fakeRepo.CreateReturns(nil, nil)

			ctx := context.Background()
			input := &entities.SimpleInterest{
				ID:             [12]byte{},
				Principal:       -1,
				ROI:            0,
				TimePeriod:     0,
				InterestAmount: 0,
				FinalAmount:    0,
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			}

			err := siSvc.CreateSimpleInterest(ctx, input)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), errInvalidPrincipal)
		},
		"create simple interest with zero ROI": func() {
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			fakeRepo.CreateReturns(nil, nil)

			ctx := context.Background()
			input := &entities.SimpleInterest{
				ID:             [12]byte{},
				Principal:       100,
				ROI:            0,
				TimePeriod:     0,
				InterestAmount: 0,
				FinalAmount:    0,
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			}

			err := siSvc.CreateSimpleInterest(ctx, input)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), errInvalidROI)
		},
		"create simple interest with negitive ROI": func() {
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			fakeRepo.CreateReturns(nil, nil)

			ctx := context.Background()
			input := &entities.SimpleInterest{
				ID:             [12]byte{},
				Principal:       100,
				ROI:            -1,
				TimePeriod:     0,
				InterestAmount: 0,
				FinalAmount:    0,
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			}

			err := siSvc.CreateSimpleInterest(ctx, input)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), errInvalidROI)
		},
		"create simple interest with zero timeperiod": func() {
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			fakeRepo.CreateReturns(nil, nil)

			ctx := context.Background()
			input := &entities.SimpleInterest{
				ID:             [12]byte{},
				Principal:       100,
				ROI:            5.5,
				TimePeriod:     0,
				InterestAmount: 0,
				FinalAmount:    0,
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			}

			err := siSvc.CreateSimpleInterest(ctx, input)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), errInvalidTimePeriod)
		},
		"create simple interest with negitive timeperiod": func() {
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			fakeRepo.CreateReturns(nil, nil)

			ctx := context.Background()
			input := &entities.SimpleInterest{
				ID:             [12]byte{},
				Principal:       100,
				ROI:            5.5,
				TimePeriod:     -1,
				InterestAmount: 0,
				FinalAmount:    0,
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			}

			err := siSvc.CreateSimpleInterest(ctx, input)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), errInvalidTimePeriod)
		},
		"create simple interest when db return error": func() {
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			testErr := errors.New("test error")
			fakeRepo.CreateReturns(nil, testErr)

			ctx := context.Background()
			input := &entities.SimpleInterest{
				ID:             [12]byte{},
				Principal:       100,
				ROI:            5.5,
				TimePeriod:     1,
				InterestAmount: 0,
				FinalAmount:    0,
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			}

			err := siSvc.CreateSimpleInterest(ctx, input)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), testErr.Error())
		},
		"create simple interest success": func() {
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			fakeRepo.CreateReturns(nil, nil)

			ctx := context.Background()
			input := &entities.SimpleInterest{
				ID:             [12]byte{},
				Principal:       100,
				ROI:            5.5,
				TimePeriod:     1,
				InterestAmount: 0,
				FinalAmount:    0,
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			}

			err := siSvc.CreateSimpleInterest(ctx, input)
			assert.Nil(t, err)
			assert.Greater(t, input.InterestAmount, float64(0))
			assert.Greater(t, input.FinalAmount, float64(0))
			assert.NotEmpty(t, input.ID)
		},
	}

	t.Log("Create SimpleInterest tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}

func TestGetSimpleInterest(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"Get simple interest empty id": func() {
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			fakeRepo.ReadReturns(nil, nil)

			ctx := context.Background()
			inputId := ""
			res, err := siSvc.GetSimpleInterest(ctx, inputId)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), errInvalidId)
			assert.Nil(t, res)
		},
		"get simple interest when db return error": func() {
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			testErr := errors.New("test error")
			fakeRepo.ReadReturns(nil, testErr)

			ctx := context.Background()
			inputId := "test"
			res, err := siSvc.GetSimpleInterest(ctx, inputId)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), testErr.Error())
			assert.Nil(t, res)
		},
		"get simple interest success": func() {
			ctx := context.Background()
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			fakeRepo.CreateReturns(nil, nil)
			input := &entities.SimpleInterest{
				ID:             [12]byte{},
				Principal:       100,
				ROI:            5.5,
				TimePeriod:     1,
				InterestAmount: 0,
				FinalAmount:    0,
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			}

			err := siSvc.CreateSimpleInterest(ctx, input)
			if err != nil {
				t.Fatal(err)
			}

			fakeRepo.ReadReturns(input, nil)
			res, err := siSvc.GetSimpleInterest(ctx, input.ID.String())
			assert.Nil(t, err)
			assert.Equal(t, res, input)
		},
	}

	t.Log("Create GetSimpleInterest tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}

func TestDeleteSimpleInterest(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"delete simple interest empty id": func() {
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			fakeRepo.DeleteReturns(nil)

			ctx := context.Background()
			inputId := ""
			err := siSvc.DeleteSimpleInterest(ctx, inputId)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), errInvalidId)
		},
		"delete simple interest when db return error": func() {
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			testErr := errors.New("test error")
			fakeRepo.DeleteReturns(testErr)

			ctx := context.Background()
			inputId := "test"
			err := siSvc.DeleteSimpleInterest(ctx, inputId)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), testErr.Error())
		},
		"delete simple interest success": func() {
			ctx := context.Background()
			fakeRepo := fakeMongo.FakeRepository{}
			siSvc := NewService(&fakeRepo)
			fakeRepo.CreateReturns(nil, nil)
			input := &entities.SimpleInterest{
				ID:             [12]byte{},
				Principal:       100,
				ROI:            5.5,
				TimePeriod:     1,
				InterestAmount: 0,
				FinalAmount:    0,
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			}

			err := siSvc.CreateSimpleInterest(ctx, input)
			if err != nil {
				t.Fatal(err)
			}

			fakeRepo.DeleteReturns(nil)
			err = siSvc.DeleteSimpleInterest(ctx, input.ID.String())
			assert.Nil(t, err)
		},
	}

	t.Log("Create GetSimpleInterest tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}
