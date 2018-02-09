package main

import (
	"./service"
	"fmt"
	"strconv"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		res, err := strconv.Atoi(os.Args[2])
		if err != nil {
			println("Invlaid argument.")
			return
		}
		if (os.Args[1] == "TodaysMenu") {
			PostTodaysMenu(res)	
			return
		}
		println("Invlaid argument.")
		return
	}

	println("Welcome to the UniCafe command line client!");
	println("Fetching restaurants...")
	restaurantService := service.NewRestaurantService()
	restaurants := restaurantService.GetRestaurants()
	println(fmt.Sprintf("%s restaurants found:", strconv.Itoa(len(restaurants))))
	for i := 0; i < len(restaurants); i++ {
		str := fmt.Sprintf("%s -> %s", restaurants[i].Name, strconv.Itoa(int(restaurants[i].ID)))
		println(str)
	}

	println("Use unicafe TodaysMenu {restaurant ID} for menus")
}

func PostTodaysMenu(restaurant int) {
	menuService := service.NewMenuService()
	menu := menuService.GetTodaysMenuForRestaurant(restaurant)
	for i := 0; i < len(menu.Data); i++ {
		println(menu.Data[i].Name)
	}
}

