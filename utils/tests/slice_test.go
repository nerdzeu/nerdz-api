package slice_test

import (
	"github.com/nerdzeu/nerdz-api/utils"
	"testing"
)


func TestValidSlice(t *testing.T) {
    letters := []string{"a", "b", "c", "d"}
    letters = utils.ReverseSlice(letters).([]string)
    if letters[len(letters)-1] != "a" {
        t.Errorf("Last letter should be 'a' but got: %v", letters[len(letters)-1])
    }
}
