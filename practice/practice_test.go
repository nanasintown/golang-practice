package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
    target := 32
    arrTest := []int{11, 312, 44, 9, 18, 23, 85}
    expected := []int{3, 5}

    result := TwoSum(arrTest, target)

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expect %v but got %v", expected, result)
    }
}
