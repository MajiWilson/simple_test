package hootTest

import "fmt"

func DogActivity(dogId string, activityMethod *DogActivityMethod) {

	fmt.Println("start******")

	activityMethod.DogBarkM.Bark()

	fmt.Println("end *******")

}
