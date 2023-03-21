package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type IAnimal interface {
	Eat()
	Move()
	Speak()
}

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

/*Animals enum*/
const (
	cow   string = "cow"
	bird  string = "bird"
	snake string = "snake"
)

/*Animal actions enum*/
const (
	eat   string = "eat"
	move  string = "move"
	speak string = "speak"
)

/* Cow factory*/
type Cow struct {
	Animal
}

func createCow() IAnimal {
	return &Cow{
		Animal: Animal{"grass", "walk", "moo"},
	}
}

/* Bird factory*/
type Bird struct {
	Animal
}

func createBird() IAnimal {
	return &Bird{
		Animal: Animal{"worms", "fly", "peep"},
	}
}

/* Snake factory*/
type Snake struct {
	Animal
}

func createSnake() IAnimal {
	return &Snake{
		Animal: Animal{"mice", "slither", "hsss"},
	}
}

/* animal dictionary */
var dict = map[string]func() IAnimal{
	cow:   createCow,
	bird:  createBird,
	snake: createSnake,
}

func getAnimalAction(a IAnimal, action *string) (act func(), err error) {
	switch *action {
	case eat:
		return a.Eat, nil
	case move:
		return a.Move, nil
	case speak:
		return a.Speak, nil
	}
	err = fmt.Errorf("Shouldn't reach here")
	return
}

func IsValidAnimal(i *string) bool {
	switch *i {
	case cow, bird, snake:
		return true
	}
	return false
}

func IsValidAnimalAction(i *string) bool {
	switch *i {
	case eat, move, speak:
		return true
	}
	return false
}

func validateInput(input *string) (animalName string, action string, e error) {
	inputStrings := strings.Split(*input, " ")

	if !IsValidAnimal(&inputStrings[0]) {
		e = fmt.Errorf("invalid animal")
	}

	if !IsValidAnimalAction(&inputStrings[1]) {
		e = fmt.Errorf("invalid animal action")
	}

	animalName = inputStrings[0]
	action = inputStrings[1]

	return
}

/* prompt for user input infinite as required in AC*/
func prompt() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(">")
	scanner.Scan()
	input := scanner.Text()
	animalName, action, err := validateInput(&input)
	if err != nil {
		fmt.Println(err)
		prompt()
	}

	act, _ := getAnimalAction(dict[animalName](), &action)
	act()
	prompt()
}

func main() {
	fmt.Println("Pls input animal (cow, bird or snake) and one of the actions (eat, move or speak). For example `cow speak`.")
	prompt()
}
