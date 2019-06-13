package main

import (
	"log"
	"net/rpc"
)

// Product Object responsible to wrap product info
type Product struct {
	Name, ID string
}

func main() {

	var reply Product
	var slice []Product

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	checkError(err)

	// Register Some Products on Storage Remote Service
	greenBean := Product{"Green Beans", "vegie01"}
	brownBean := Product{"Brown Beans", "vegie02"}
	greenApple := Product{"Green Apple", "fruit01"}
	redApple := Product{"Red Apple", "fruit02"}

	err = client.Call("Storage.AddProduct", greenBean, &reply)
	checkError(err)
	log.Println("Product Added: ", reply)

	err = client.Call("Storage.AddProduct", brownBean, &reply)
	checkError(err)
	log.Println("Product Added: ", reply)

	err = client.Call("Storage.AddProduct", greenApple, &reply)
	checkError(err)
	log.Println("Product Added: ", reply)

	err = client.Call("Storage.AddProduct", redApple, &reply)
	checkError(err)
	log.Println("Product Added: ", reply)

	// Get all products
	client.Call("Storage.GetProducts", "", &slice)
	log.Println("Stored Products: ", slice)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
