/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package utils_test

import (
	"fmt"
	"testing"

	"github.com/nerdzeu/nerdz-api/utils"
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

	if utils.InSlice("c", letters) == false {
		t.Errorf("Letter c is present in letters slice")
	}

	fmt.Println("Letter test ok")

	var horse, nope Amazing
	horse.Animal = "weebl"
	nope.Animal = "nerdz"

	complexData := []ComplexData{
		{One: 1, Two: "lol", Horse: horse},
		{One: 2, Two: "asd", Horse: nope}}

	fmt.Printf("Before: %+v\n", complexData)

	complexData = utils.ReverseSlice(complexData).([]ComplexData)

	if complexData[0].Horse.Animal != "nerdz" {
		t.Errorf("Animal should be nerdz, but got: %v", complexData[0].Horse.Animal)
	}

	fmt.Printf("After: %+v\n", complexData)

	if utils.InSlice(complexData[1], complexData) == false {
		t.Errorf("This value is present in complexData slice")
	}

	if utils.InSlice("banana", complexData) {
		t.Errorf("Banana is not into complexData slice (and have different type)")
	}

}

func TestParseTag(t *testing.T) {
}

func TestUpperFirst(t *testing.T) {
	if utils.UpperFirst("ciao") != "Ciao" {
		t.Errorf("UpperFirst does not work")
	}
}
