package main

import "maojainwei/simple_test/hootTest"

func main() {
	method := &hootTest.DogActivityMethod{
		Species: "nsds"}
	hootTest.SetDogBark(method)

	hootTest.DogActivity("0", method)
}
