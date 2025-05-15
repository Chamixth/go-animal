package main

import (
	"go-animal/dtos"
	"go-animal/services"
)

func main() {
	dog := dtos.Dog{Name: "Chamith"}
	cat := dtos.Cat{Name: "Chamika"}

	services.MakeItSpeak(dog)
	services.MakeItSpeak(cat)
}