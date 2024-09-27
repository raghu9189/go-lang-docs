package main

import (
	"fmt"
)

func callMeFun(paraName string) int {
	myNameString := paraName
	fmt.Println("Called Me!", paraName, myNameString)
	return 45
}
func main() {
	var myBool bool = true
	var myInteger uint = 2345
	var myFloat float32 = 34250.9800
	var myName string = "Raghu"
	var myArray = [...]int{1, 2, 3}
	mySlice := []int{}
	mySlice = append(mySlice, 20, 23)
	fmt.Print(myBool, "\n", myInteger, myFloat, myName, myArray, mySlice)
	myNames := []string{"Raghu", "Mahesh", "Shilesh"}

	for index, value := range myNames {
		fmt.Println(index, value)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < len(myNames); i++ {
		fmt.Println(myNames[i])
	}
	var myNumReturnded int = callMeFun("Srikanth")
	fmt.Println(myNumReturnded)

	myMap := map[string]int{}
	myMap["code"] = 1
	myMap["id"] = 23

	for id, value := range myMap {
		println(id, value)
	}

	var exMyFloat float32 = 45.89

	println(int(exMyFloat))

}
