package main

func main() {}

type menu struct {
	Breakfast []item
	Lunch []item
	Dinner []item
}

type item struct {
	Price string
	Item string
	Description string
}

type menus []menu