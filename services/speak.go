package services

import (
	"fmt"
	"go-animal/dtos"
)

func MakeItSpeak(a dtos.Animal){
	fmt.Println(a.Speak())
}