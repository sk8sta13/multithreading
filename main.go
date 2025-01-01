package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {
	if len(os.Args) != 2 || !validateCep(os.Args[1]) {
		fmt.Println("Enter a valid zip code")
		return
	}

	url1 := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", os.Args[1])
	a1 := make(chan Address)

	url2 := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", os.Args[1])
	a2 := make(chan Address)

	go func() {
		a1 <- getAddress(url1)
	}()

	go func() {
		a2 <- getAddress(url2)
	}()

	select {
	case address1 := <-a1:
		showAddress(address1, url1)
	case address2 := <-a2:
		showAddress(address2, url2)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout")
	}
}

func showAddress(address Address, origin string) {
	fmt.Println("The first API to respond before timeout is", origin)
	fmt.Println("Zip Code:", address.Code)
	fmt.Println("State:", address.State)
	fmt.Println("City:", address.City)
	fmt.Println("Neighborhood:", address.Neighborhood)
	fmt.Println("Street:", address.Street)
}

func validateCep(code string) bool {
	re := regexp.MustCompile("^[0-9]{8}$")
	return re.MatchString(code)
}

func getAddress(url string) Address {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var a Address
	err = json.NewDecoder(resp.Body).Decode(&a)
	if err != nil {
		log.Fatal(err)
	}

	return a
}
