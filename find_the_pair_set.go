package main

import (
	"sort"
)

type pairSet struct {
	x int
	y int
}

type FindThePairSetInpunt struct {
	NumberList []int
	Target     int
}

func GetFindThePairSetInput() FindThePairSetInpunt {
	return FindThePairSetInpunt{
		[]int{1, 2, 3, 4, 5, 65, 23, 12, 5, 6, 45, 123, 764, 8, 5, 3, 45, 876, 234, 1233, 6, 46, 34, 12, 34, 3, 221, 12, 1, 2, 23, 4, 8, 9, 0, 3, 1},
		4,
	}
}

func FindThePairSet(numberList []int, target int) (pairSets []pairSet) {
	// in order to avoid changes on the original list
	list := make([]int, len(numberList))
	copy(numberList, list)

	// the list must be sorted in order to be processed efficiently
	sort.Ints(list)

	// we need control over 2 indexes of the list to compare
	indexA := 0
	indexB := len(list) - 1

	for indexA < indexB {
		sum := list[indexA] + list[indexB]
		if sum > target {
			indexB--
			continue
		}
		if sum < target {
			indexA++
			continue
		}
		if list[indexA] == list[indexB] {
			// TODO if the rest are equals we have to add all the combinations of each other
			pairSets = append(pairSets, pairSet{list[indexA], list[indexB]})
			break
		}
		pairSets = append(pairSets, pairSet{list[indexA], list[indexB]})
	}
	return
}
