package main

import (
	"context"
	"fmt"

	proto "gomicro-samples/helloworld/proto"

	micro "github.com/micro/go-micro/v2"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	var words string
	for {

		fmt.Printf("Please input string : ")
		n, err := fmt.Scanf("%s\n", &words)
		if err != nil {
			fmt.Println(n, err)
		}

		rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: words})
		if err != nil {
			fmt.Println(err)
		}

		// Print response
		fmt.Println(rsp.Greeting)
	}

}
