package dto

import "admin/microservice/helpers"

type NavigationBar struct {
	NavigationBarID   string `json:"navigation_bar_id" bson:"navigation_bar_id"`
	NavigationBarName string `json:"navigation_bar_name" bson:"navigation_bar_name"`
	TimeAt            TimeAt `json:"time_at" bson:"time_at"`
}

func NewNavigationBar() NavigationBar {
	return NavigationBar{
		NavigationBarID:   helpers.GeneratedUUID(),
		NavigationBarName: "",
		TimeAt: TimeAt{
			CreatedAt:              helpers.GetCurrentUnix(),
			HumanReadableCreatedAt: helpers.TimeNowStr(),
		},
	}
}
