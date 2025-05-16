package main

import (
	"fmt"
	"github.com/vinmazzi/gocmdtest/internal/auth"
)

func main() {
	auth.Make()
	fmt.Println("Finishing command execution.")
}
