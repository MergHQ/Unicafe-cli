package main

import (
	"./service"
	"fmt"
	"strconv"
	"os"
)

func main() {
	if len(os.Args) >= 3 {
		if (os.Args[1] == "TodaysMenu") {
      res, err := strconv.Atoi(os.Args[2])
      if err != nil {
        println("Invlaid argument.")
        return
      }
			PostTodaysMenu(res)	
			return
		} else if len(os.Args) == 4 && os.Args[1] == "ShowMenu" {
      restaurant, err := strconv.Atoi(os.Args[3])
      if err != nil {
        println("Invalid argument")
        return
      }
      PostMenuForDate(os.Args[2], restaurant)
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
	restaurantInformation, menu := menuService.GetTodaysMenuForRestaurant(restaurant)
	println(restaurantInformation.Name)
	println(restaurantInformation.Address)
	print("\n\n\n")
	for i := 0; i < len(menu.Data); i++ {
		str := fmt.Sprintf("%s | Student price: %s", menu.Data[i].Name, menu.Data[i].Price.Value.Student)
		println(str)
	}
}

func PostMenuForDate(date string, restaurant int) {
  menuService := service.NewMenuService()
	restaurantInformation, menu := menuService.GetMenuForRestaurant(date, restaurant)
	println(restaurantInformation.Name)
	println(restaurantInformation.Address)
	print("\n\n\n")
	for i := 0; i < len(menu.Data); i++ {
		str := fmt.Sprintf("%s | Student price: %s", menu.Data[i].Name, menu.Data[i].Price.Value.Student)
		println(str)
	}
}

