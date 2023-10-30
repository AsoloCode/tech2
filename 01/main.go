package main

import "fmt"

type Human struct {
	firstName string
	lastName  string
	age       int
	height    int
	weight    int
}

type Activity struct {
	profession string
	jobTitle   string
}

func (h *Human) Speak() {
	fmt.Printf("%s говорит: Привет, моя фамилия %s, мне %d. \n А мой рост составлят %d см и вес %d кг.", h.firstName, h.lastName, h.age, h.height, h.weight)
}
func (a *Activity) MyActivity() {
	fmt.Printf("Я работаю %s \n На должности: %s", a.profession, a.jobTitle)
}

type Action struct {
	Human
	Activity
}

func main() {
	a := Action{
		Human: Human{
			firstName: "Artem",
			lastName:  "Solovev",
			age:       23,
			height:    170,
			weight:    75,
		},
		Activity: Activity{
			profession: "Golang developer",
			jobTitle:   "Middle Backend",
		},
	}

	a.Speak()
	a.MyActivity()

}
