package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SimpleInterest struct {
	ID             primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Princpal       float64            `json:"princpal" bson:"princpal"`
	ROI            float64            `json:"roi" bson:"roi"`
	TimePeriod     float64            `json:"timePeriod" bson:"timePeroid"`
	InterestAmount float64            `json:"interestAmount" bson:"interestAmount"`
	FinalAmount    float64            `json:"finalAmount" bson:"finalAmount"`
	CreatedAt      time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt" bson:"updatedAt"`
}
