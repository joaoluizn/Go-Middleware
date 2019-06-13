package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Storage Storage Service Type
type Storage int

// Product Object responsible to wrap product info
type Product struct {
	Name, ID string
}

// Our Storage
var productSlice []Product

// (AddProduct, GetProducts) are methods, it follows the RPC documentation
// Instead of returning a value, we use a pointer to the reply param and return nil to indicate no errors.
// If we do want to return an error we just return it and the reply parameter will not be sent back to the client.

// AddProduct A Storage Operation that append a new product to the Storage
func (s *Storage) AddProduct(product Product, reply *Product) error {
	productSlice = append(productSlice, product)
	*reply = product
	return nil
}

// GetProducts A Storage Operation that returns all Products stored
func (s *Storage) GetProducts(title string, reply *[]Product) error {
	*reply = productSlice
	return nil
}

func main() {
	// Object storage will be used to call our methods
	storage := new(Storage)

	// Register service on rpc
	err := rpc.Register(storage)
	checkError(err)

	// Handle HTTP is responsible for routing this remote service
	rpc.HandleHTTP()

	// Out listener to intercept calls
	listener, err := net.Listen("tcp", ":1234")
	checkError(err)

	log.Printf("Listening RPC server on port %d", 1234)

	err = http.Serve(listener, nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
