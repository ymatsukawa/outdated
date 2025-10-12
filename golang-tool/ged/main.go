package main

import (
	"fmt"

	"github.com/ymatsukawa/ged/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
