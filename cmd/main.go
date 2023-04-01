package main

import (
	"fmt"
	"github.com/NurfitraPujo/finance-tracker/build"
	"github.com/NurfitraPujo/finance-tracker/config"
)

func main() {
	fmt.Println("build.Environment:\t", build.Environment)
	fmt.Println("build.Version:\t", build.Version)
	fmt.Println("build.Date:\t", build.Date)
	fmt.Println("build.GitCommit:\t", build.GitCommit)

	config.LoadConfig()
}
