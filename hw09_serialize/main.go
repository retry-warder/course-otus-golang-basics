package main

import (
	"encoding/json"
	"fmt"

	"github.com/retry-warder/course-otus-golang-basics/hw09_serialize/types"
	"github.com/retry-warder/course-otus-golang-basics/hw09_serialize/types_pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	b1 := types.NewBook(1, "Test 1", "Test.T.T", 2010, 1000, 10.10)
	j1, err := json.Marshal(&b1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b1 json %s\n", j1)

	var b1u types.Book
	err = json.Unmarshal(j1, &b1u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", b1)
	fmt.Printf("%v\n", b1u)
	fmt.Printf("%v\n", "------------------------------------------------------------")
	b2 := types.NewBook(1, "Test 2", "Test.X.Y", 2000, 1500, 10.90)
	j2, err := json.Marshal(b2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b2 json %s\n", j2)
	var b2u types.Book
	err = json.Unmarshal(j2, &b2u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", b2)
	fmt.Printf("%v\n", b2u)
	fmt.Printf("%v\n", "------------------------------------------------------------")
	var books []types.Book
	books = append(books, b1u, b2u)
	j := types.EnCodeJSON(books)
	fmt.Printf("books1 json %s\n", j)
	var booksu []types.Book
	err = json.Unmarshal(j, &booksu)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", books)
	fmt.Printf("%v\n", booksu)
	fmt.Printf("%v\n", "------------------------------------------------------------")

	b3 := types_pb.NewBook(1, "Test 1", "Test.T.T", 2010, 1000, 10.10)
	marshaled1, err := proto.Marshal(b3)
	if err != nil {
		panic(err)
	}
	fmt.Println(b3)
	fmt.Println(marshaled1)
	fmt.Println(len(marshaled1))
	b3u := &types_pb.Book{}
	err = proto.Unmarshal(marshaled1, b3u)
	if err != nil {
		panic(err)
	}
	fmt.Println(b3u)
	fmt.Printf("%v\n", "------------------------------------------------------------")

	b4 := types_pb.NewBook(1, "Test 2", "Тест Т.Т.", 2000, 1500, 10.90)
	marshaled2, err := proto.Marshal(b4)
	if err != nil {
		panic(err)
	}
	fmt.Println(b4)
	fmt.Println(marshaled2)
	fmt.Println(len(marshaled2))
	b4u := &types_pb.Book{}
	err = proto.Unmarshal(marshaled2, b4u)
	if err != nil {
		panic(err)
	}
	fmt.Println(b4u)
	fmt.Printf("%v\n", "------------------------------------------------------------")
}
