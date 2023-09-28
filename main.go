package main

import (
    "encoding/json"
    "fmt"
    "os"
)

// Making a quiz game with login

type Person struct {
    Name  string `json:"name"`
    Score   int    `json:"score"`
    Email string `json:"email"`
}

func main() {

  var name string
  fmt.Println("Set your name: ")

  var email string
  fmt.Println("Set your email: ")
  
  person := Person{
    Name: input,
    Score: 0,
    Email: email,
  }
  
  jsonData, err := json.MarshalIndent(person, "", "  ")

  // To be finished
  fmt.Scanln(&input)
	fmt.Println(name, email)
}
