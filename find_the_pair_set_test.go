package main

import (
	"reflect"
	"testing"
)

func TestFindThePairSet(t *testing.T) {
	testingTable := []struct {
		title          string
		input          FindThePairSetInpunt
		expectedResult []PairSet
	}{
		{"basic", FindThePairSetInpunt{[]int{7, 1, 3, 5, 4, 2, 6}, 8}, []PairSet{{1, 7}, {2, 6}, {3, 5}}},
		{"contains_0", FindThePairSetInpunt{[]int{7, 1, 3, 5, 0, 4, 2, 6, 8}, 8}, []PairSet{{0, 8}, {1, 7}, {2, 6}, {3, 5}}},
		{"no_pair_sets_higher", FindThePairSetInpunt{[]int{7, 1, 3, 5, 0, 2, 6, 8}, 44}, []PairSet{}},
		{"no_pair_sets_lower", FindThePairSetInpunt{[]int{7, 11, 13, 5, 10, 22, 6, 8}, 3}, []PairSet{}},
		{"empty_list", FindThePairSetInpunt{[]int{}, 3}, []PairSet{}},
		{"list_one_number", FindThePairSetInpunt{[]int{9}, 3}, []PairSet{}},
		{"list_only_target", FindThePairSetInpunt{[]int{3}, 3}, []PairSet{}},
		{"list_only_target_pair", FindThePairSetInpunt{[]int{3, 0}, 3}, []PairSet{{0, 3}}},
	}

	for _, testingCase := range testingTable {
		t.Run(testingCase.title, func(t *testing.T) {
			result := FindThePairSet(testingCase.input.NumberList, testingCase.input.Target)
			if !reflect.DeepEqual(result, testingCase.expectedResult) {
				t.Errorf("Test case failed with input: %v, expected: %v; got: %v", testingCase.input, testingCase.expectedResult, result)
			}
		})
	}
}
