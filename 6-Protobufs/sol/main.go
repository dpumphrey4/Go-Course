package main

import (
	"flag"
	"fmt"
	"protobufs/gophers"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func marshalProto(gopher *gophers.Gopher) {
	gopherBytes, err := proto.Marshal(gopher)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully marshalled gopher into protobuf bytes.")
	fmt.Printf("%+v\n", gopherBytes)
	fmt.Printf("Length in bytes: %d\n", len(gopherBytes))

	// Converting back to proto message struct
	message := &gophers.Gopher{}

	err = proto.Unmarshal(gopherBytes, message)
	if err != nil {
		fmt.Println(err)
	}
}

func marshalJson(gopher *gophers.Gopher) {
	gopherBytes, err := protojson.Marshal(gopher)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully marshalled gopher into Json.")
	fmt.Println(string(gopherBytes))
	fmt.Printf("Length in bytes: %d\n", len(gopherBytes))

	// Converting back to proto message struct
	message := &gophers.Gopher{}

	err = protojson.Unmarshal(gopherBytes, message)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	parseType := flag.String("type", "protobuf", "Specifies which serialized format to use.")

	flag.Parse()

	if *parseType != "protobuf" && *parseType != "json" {
		panic("Failed to provide proper parse type.")
	}

	// This is the struct that we shall be serializing or as go calls it "marshal"
	primaryGopher := gophers.Gopher{Name: "Dave", Age: 10000, Burrows: []string{"New York", "Paris", "London", "Berlin", "Montreal", "Kyiv", "Antartica"}}

	if *parseType == "json" {
		marshalJson(&primaryGopher)
	} else {
		marshalProto(&primaryGopher)
	}
}
