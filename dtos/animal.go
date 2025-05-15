package dtos

type Animal interface {
	Speak() string
}

type Dog struct{
	Name string
}

type Cat struct {
	Name string
}

func (dog Dog) Speak() string{
	return dog.Name + " " + "barks."
}
func (cat Cat) Speak() string{
	return cat.Name + " " + "meows."
}