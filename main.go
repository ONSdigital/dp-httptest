package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ONSdigital/dp-mocking/httpmocks"
)

func main() {
	input := []string{"1", "2", "3"}
	b, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}

	readCloser := httpmocks.NewReadCloserMock(b, nil)

	bOut, err := ioutil.ReadAll(readCloser)
	if err != nil {
		panic(err)
	}

	var output []string
	err = json.Unmarshal(bOut, &output)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

	// Output
	// [1 2 3]
	// 200
}
