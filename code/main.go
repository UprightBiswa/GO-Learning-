// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// var originalCount int = 10
// 	// var eatenCount int = 4
// 	// fmt.Println("I started with", originalCount, "apples.")
// 	// fmt.Println("Some jerk ate", eatenCount, "apples.")
// 	// fmt.Println("There are", originalCount-eatenCount, "apples left.")
// 	// fmt.Print("Enter a number: ")
// 	// reader := bufio.NewReader(os.Stdin)
// 	// // input, _ := reader.ReadString('\n')
// 	// input, err := reader.ReadString('\n')
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// fmt.Println("You entered:", input)

// 	// if false {
// 	// 	fmt.Println("1 is less than 2")
// 	// }
// 	// ifelse()

// 	// fmt.Print("Enter a grade: ")
// 	// reader := bufio.NewReader(os.Stdin)
// 	// input, err := reader.ReadString('\n')
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// grade, err := strconv.Atoi(strings.TrimSpace(input))
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// if grade == 100 {
// 	// 	fmt.Println("Perfect!")
// 	// } else if grade >= 60 {
// 	// 	fmt.Println(calculateAverage([]float64{1, 2, 3, 4, 5}))

// 	// 	fmt.Println("You passed!")
// 	// } else {
// 	// 	fmt.Println("You failed!")
// 	// }

// 	// target := rand.Intn(100) +1
// 	// fmt.Println(target)
// 	// seconds := time.Now().Unix()
// 	// rand.NewSource(seconds)
// 	// target := rand.Intn(100)+1
// 	// fmt.Println("I've chosen a random number between 1 and 100.")
// 	// fmt.Println("Can you guess it?")
// 	// fmt.Println(target)

// 	for x:= 4; x<=6; x++{
// 		fmt.Println(x)
// 	}

// }

// // func calculateAverage(numbers []float64) float64 {
// // 	var sum float64
// // 	for _, num := range numbers {
// // 		sum += num
// // 	}
// // 	return sum / float64(len(numbers))
// // }

// // func ifelse() {
// // 	if 1 == 1 {
// // 		fmt.Println("1 is equal to 1")
// // 	}
// // 	if 1 == 0 {
// // 		fmt.Println("1 is equal to 0")
// // 	}
// // 	if 1 != 0 {
// // 		fmt.Println("1 is not equal to 0")
// // 	}
// // 	if 1 > 0 {
// // 		fmt.Println("1 is greater than 0")
// // 	}
// // 	if true && true{
// // 		fmt.Println("true and true")
// // 	}

// // 	if true{
// // 		fmt.Println("true")
// // 	}
// // 	if false{
// // 		fmt.Println("false")
// // 	}
// // 	if(12 == 12 || 5.9 ==6.4){
// // 		fmt.Println("12 is equal to 12 or 5.9 is equal to 6.4")
// // 	}
// // }

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"math/rand"
)
func main(){
	seconds := time.Now().Unix()
	rand.NewSource(seconds)
	target := rand.Intn(100)+1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses< 10;guesses++{
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		} 
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess< target{
			fmt.Println("Oops. Your guess was LOW.")
		}else if guess> target{
			fmt.Println("Oops. Your guess was HIGH.")
		}else{
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}
	if !success{
		fmt.Println("Sorry. You didn't guess my number. It was:", target)
	}
}