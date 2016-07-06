package main

import (
	"fmt"
	"github.com/adriendomoison/ActivityTrackerAPI/repository"
	"github.com/adriendomoison/ActivityTrackerAPI/api"
)

func main() {
	fmt.Println("Hello World!")
	repository.InitDBConnection()
	api.InitApi()
}
