package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// /*
// swap values in the slice using references
// */
// func Swap(first *int32, second *int32) {
//   previousFirstValue := *first
//   *first = *second
//   *second = previousFirstValue
// }

/*
swap values directly in the slice that is AC criteria
*/
func Swap(sli []int32, index int) {
	sli[index], sli[index+1] = sli[index+1], sli[index]
}

/*
sort
*/
func BubbleSort(sli []int32) {
  arrLength := len(sli)
  for i := 0; i < arrLength - 1; i++ {
    for j := 0; j < arrLength - 1 - i; j++ {
      if sli[j] > sli[j+1] {
        //Swap(&sli[j], &sli[j+1])
        Swap(sli, j)
      }
    }
  } 
}

/*
user input serializer
*/
func serializeInput(input *string) []int32 {
  var sli []int32
  inputStrings := strings.Split(*input, ",")
  for _, v := range inputStrings {
    number, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64)

		if err != nil {
			fmt.Println("Got an error", err)
      panic(err)
		}
    sli = append(sli, int32(number))
  }

  return sli
}
 
func main () {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Pls provide comma separated string of integers. (ex. 5,2,4,1,4)")
  scanner.Scan()
	input := scanner.Text()
  intArray := serializeInput(&input)
	BubbleSort(intArray[:])
  fmt.Println("Here is your sorted array:", intArray)
}
