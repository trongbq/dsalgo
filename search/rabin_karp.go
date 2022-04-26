package search

import (
	"math"
	"strings"
)

// RabinKarp returns the index in t where s matchs a part of t
// It returns -1 if s is not in t
func RabinKarp(s, t string) int {
	if len(s) > len(t) {
		return -1
	}

	base := 26
	pow := int(math.Pow(float64(base), float64(len(s)-1)))
	hashFunc := func(str string) int {
		// hash(abc) = a*26^2 + b*26 + c
		v := 0
		for _, rune := range str {
			v = v*base + int(rune)
		}
		return v
	}
	isEqual := func(sv, tv int, ts string) bool {
		// Check against two hash code value and their substring to see
		// if they are indeed equal, to protect from hash collision
		return sv == tv && strings.Compare(s, ts) == 0
	}

	// Calculate hash value of s and first substring of t
	// then compare them to see if a match
	sValue := hashFunc(s)
	tSubValue := hashFunc(t[:len(s)])
	if isEqual(sValue, tSubValue, t[0:len(s)]) {
		return 0
	}

	// Otherwise, start rolling hash value and continue compare to find a match
	for i := 1; i < len(t)-len(s)+1; i++ {
		tSubValue = (tSubValue-int(t[i-1])*pow)*base + int(t[i+len(s)-1])
		if isEqual(sValue, tSubValue, t[i:i+len(s)]) {
			return i
		}
	}

	return -1
}
