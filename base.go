package gomdbuilder

import (
	"fmt"
)

const (
	LE = "\n"
)

func stringifySlice(slice []interface{}) []string {
	sliceStr := make([]string, len(slice))
	for i, v := range slice {
		sliceStr[i] = fmt.Sprint(v)
	}
	return sliceStr
}

func rangeStrings(f func(int, string)string, strs []string) []string {
	newStrs := make([]string, len(strs))
	for i, v := range strs {
		newStrs[i] = f(i, v)
	}
	return newStrs
}