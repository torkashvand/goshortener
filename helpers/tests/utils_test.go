package tests

import (
	"testing"

	"github.com/torkashvand/goshortener/helpers"
)

func TestConvertBase(t *testing.T) {
	inputSlice := [5]uint{1, 500, 125612, 12, 121213424564546}
	expectedResults := [5]string{"1", "84", "wG0", "c", "yq1SwCc2"}
	var retunedResult [5]string

	for i, num := range inputSlice {
		result := helpers.ConvertBase(num)
		retunedResult[i] = result
	}

	if expectedResults != retunedResult {
		t.Errorf("TestConvertBase FAILED, expected value %v but got %v", expectedResults, retunedResult)
	}
}
