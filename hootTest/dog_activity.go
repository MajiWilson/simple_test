package hootTest

import "fmt"

type DogActivityMethod struct {
	Species  string
	DogBarkM DogBark
}

type DogBark interface {
	Bark()
}

type ShibaDogBark struct {
}

func (sdb *ShibaDogBark) Bark() {
	fmt.Printf("xixiixxixixiix")
}

type NormalDogBark struct {
}

func (sdb *NormalDogBark) Bark() {
	fmt.Printf("wanwgang")
}

func SetDogBark(dam *DogActivityMethod) {
	if dam.Species == "shiba" {
		dam.DogBarkM = &ShibaDogBark{}
	} else {
		dam.DogBarkM = &NormalDogBark{}
	}
}
