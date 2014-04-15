package slice_test

import (
	"fmt"
	"github.com/nerdzeu/nerdz-api/utils"
	"testing"
)

type Amazing struct {
	Animal string
}

type ComplexData struct {
	One   int
	Two   string
	Horse Amazing
}

func TestValidSlice(t *testing.T) {
	letters := []string{"a", "b", "c", "d"}
	letters = utils.ReverseSlice(letters).([]string)
	if letters[len(letters)-1] != "a" {
		t.Errorf("Last letter should be 'a' but got: %v", letters[len(letters)-1])
	}

	fmt.Println("Letter test ok")

	var horse, nope Amazing
	horse.Animal = "weebl"
	nope.Animal = "nerdz"

	complexData := []ComplexData{
		ComplexData{One: 1, Two: "lol", Horse: horse},
		ComplexData{One: 2, Two: "asd", Horse: nope}}

	fmt.Printf("Before: %+v\n", complexData)

	complexData = utils.ReverseSlice(complexData).([]ComplexData)

	if complexData[0].Horse.Animal != "nerdz" {
		t.Errorf("Animal should be nerdz, but got: %v", complexData[0].Horse.Animal)
	}

	fmt.Printf("After: %+v\n", complexData)

}
