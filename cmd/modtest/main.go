package main

import (
	"fmt"
	auth "github.com/vinmazzi/gocmdtest/internal"
	"github.com/vinmazzi/gocmdtest/models"
)

func main() {
	auth.Make()
	fmt.Println("Finishing command execution.")
	models.TestModel()
}
