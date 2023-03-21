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
	GetName() string
}

type Animal struct {
	food, locomotion, noise, name string
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

func (a *Animal) GetName() string {
	return a.name
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

/*actions*/
const (
	newanimal string = "newanimal"
	query     string = "query"
)

/* Cow factory*/
type Cow struct {
	Animal
}

func createCow(name string) IAnimal {
	return &Cow{
		Animal: Animal{"grass", "walk", "moo", name},
	}
}

/* Bird factory*/
type Bird struct {
	Animal
}

func createBird(name string) IAnimal {
	return &Bird{
		Animal: Animal{"worms", "fly", "peep", name},
	}
}

/* Snake factory*/
type Snake struct {
	Animal
}

func createSnake(name string) IAnimal {
	return &Snake{
		Animal: Animal{"mice", "slither", "hsss", name},
	}
}

var mem = make(map[string]IAnimal)

/* animal dictionary */
var dict = map[string]func(n string) IAnimal{
	cow:   createCow,
	bird:  createBird,
	snake: createSnake,
}

func getAnimalAction(name, action *string) (func(), error) {
	a, ok := mem[*name]
	if !ok {
		return nil, fmt.Errorf("animal not found: %s", *name)
	}

	validActions := map[string]func(){"eat": a.Eat, "move": a.Move, "speak": a.Speak}
	act, ok := validActions[*action]
	if !ok {
		return nil, fmt.Errorf("invalid action: %s", *action)
	}

	return act, nil
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

func IsValidAction(i *string) bool {
	switch *i {
	case newanimal, query:
		return true
	}
	return false
}

func validateNewAnimal(name, animalType *string) (string, string, error) {
	if len(strings.TrimSpace(*name)) < 1 {
		return "", "", fmt.Errorf("invalid name")
	}

	if !IsValidAnimal(animalType) {
		return "", "", fmt.Errorf("invalid animal")
	}
	return *name, *animalType, nil

}
func validateQuery(name, animalAction *string) (string, string, error) {
	if len(strings.TrimSpace(*name)) < 1 {
		return "", "", fmt.Errorf("invalid name")
	}

	if !IsValidAnimalAction(animalAction) {
		return "", "", fmt.Errorf("invalid animal action")
	}
	return *name, *animalAction, nil
}

func validateInput(input *string) (string, string, string, error) {
	inputStrings := strings.Split(strings.TrimSpace(strings.ToLower(*input)), " ")

	if len(inputStrings) != 3 {
		return "", "", "", fmt.Errorf("invalid input")
	}

	if !IsValidAction(&inputStrings[0]) {
		return "", "", "", fmt.Errorf("invalid action")
	}

	switch inputStrings[0] {
	case newanimal:
		name, animalType, err := validateNewAnimal(&inputStrings[1], &inputStrings[2])
		if err != nil {
			return "", "", "", fmt.Errorf("invalid `newanimal` action")
		}
		return name, animalType, "", nil
	case query:
		name, action, err := validateQuery(&inputStrings[1], &inputStrings[2])
		if err != nil {
			return "", "", "", fmt.Errorf("invalid `query` action. err: %w", err)
		}
		return name, "", action, nil
	}
	return "", "", "", fmt.Errorf("invalid input")
}

/* prompt for user input infinite as required in AC*/
func prompt(s *bufio.Scanner) {
	fmt.Printf(">")
	if !s.Scan() {
		return
	}
	input := s.Text()
	animalName, animalType, action, err := validateInput(&input)
	if err != nil {
		fmt.Println(err)
		prompt(s)
	}
	if animalType != "" {
		mem[animalName] = dict[animalType](animalName)
		fmt.Println("Created it!")
		prompt(s)
	}
	act, err := getAnimalAction(&animalName, &action)
	if err != nil {
		fmt.Println(err)
		prompt(s)
	}
	act()
	prompt(s)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt(scanner)
}
