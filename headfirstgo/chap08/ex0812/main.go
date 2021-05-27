package main

import (
	"fmt"

	"../magazine"
)

func main() {
	address := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Address = address
	fmt.Println(subscriber.Address)
	fmt.Printf("%#v\n", subscriber.Address)
}
