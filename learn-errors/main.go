package main

import (
  "fmt"
  "time"
)


type MyError struct {
  What time.Time  
  When string
}

func (e *MyError) Error() string {
  return fmt.Sprintf("at %v %s", e.When, e.What)
}

func run() error {
  return &MyError(
    time.Now(),
    "it didn't work",
  )
}

func main() {
  fmt.Println("Errors main")

  if err := run(); err != nil {
    fmt.Println(err)
  }
}
