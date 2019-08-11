package main

import "fmt"

func getCurrentDay() {

}

func generateOutput() {

}

func archive() {
	fmt.Println("not implemented yet")
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err)
	}
}
