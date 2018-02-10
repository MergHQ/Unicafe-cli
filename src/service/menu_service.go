package service

import (
	"github.com/dghubble/sling"
	"../domain"
	"strconv"
	"fmt"
	"strings"
	"time"
)

type MenuResponse struct {
	Data []domain.Menu `json:"data,omitempty"`
	Information domain.RestaurantInfoLong
	Status string `json:"status,omitempty"`
}
/*
	Menu service
*/
type MenuService struct {
	sling* sling.Sling
}

func NewMenuService() *MenuService {
	return &MenuService {
		sling: sling.New().Base("https://messi.hyyravintolat.fi/publicapi"),
	}
}

func (ms* MenuService) GetTodaysMenuForRestaurant(restaurant int) (domain.RestaurantInfoLong, domain.Menu) {
	response := new(MenuResponse)
	path := fmt.Sprintf("https://messi.hyyravintolat.fi/publicapi/restaurant/%s", strconv.Itoa(restaurant))
	_, err := ms.sling.New().Get(path).Receive(response, nil)
	if err != nil {
		panic(err)
	}
	
	for i := 0; i < len(response.Data); i++ {
		splitDate := strings.Split(response.Data[i].Date, " ");
		if time.Now().Format("02.01") == splitDate[1] {
			return response.Information, response.Data[i];
		}
	}
	return response.Information, domain.Menu{} 
}