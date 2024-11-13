package dto

import "admin/microservice/helpers"

type Rank struct {
	RankID    string  `json:"rank_id" bson:"rank_id"`
	RankName  string  `json:"rank_name" bson:"rank_name"`
	RankImage string  `json:"rank_image" bson:"rank_image"`
	Min       float64 `json:"min" bson:"min"`
	Max       float64 `json:"max" bson:"max"`
	TimeAt    TimeAt  `json:"-" bson:"time_at"`
}

func NewRank() Rank {
	return Rank{
		RankID:    helpers.GeneratedUUID(),
		RankName:  "",
		RankImage: "",
		Min:       0,
		Max:       0,
		TimeAt: TimeAt{
			CreatedAt:              helpers.GetCurrentUnix(),
			HumanReadableCreatedAt: helpers.TimeNowStr(),
			UpdatedAt:              0,
			HumanReadableUpdatedAt: "",
		},
	}
}
