package service

import (
	"github.com/dghubble/sling"
	"../domain"
)

type RestaurantResponse struct {
	Data []domain.RestaurantInfoShort `json:"data,omitempty"`
	Status string `json:"status,omitempty"`
}

/*
	Restaurant service
*/
type RestaurantService struct {
	sling* sling.Sling
}

func NewRestaurantService() *RestaurantService {
	return &RestaurantService {
		sling: sling.New().Base("https://messi.hyyravintolat.fi/publicapi"),
	}
}

func (rs* RestaurantService) GetRestaurants() []domain.RestaurantInfoShort {
	response := new(RestaurantResponse)
	_, err := rs.sling.New().Get("https://messi.hyyravintolat.fi/publicapi/restaurants").Receive(response, nil)
	if err != nil {
		println(err.Error())
		return nil
	}
	return response.Data
}