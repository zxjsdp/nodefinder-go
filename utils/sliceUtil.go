package utils

func CheckRuneInRunesV2(runes []rune, r rune) bool {
	set := make(map[rune]bool)
	for _, v := range runes {
		set[v] = true
	}
	return set[r]
}

func CheckRuneInRunesV1(runes []rune, r rune) bool {
	for _, each := range runes {
		if each == r {
			return true
		}
	}
	return false
}

func Reverse(numbers []int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

func FindLongerAndShorterArray(arrayA, arrayB []int) (longerArray, shorterArray []int) {
	if len(arrayA) >= len(arrayB) {
		return arrayA, arrayB
	} else {
		return arrayB, arrayA
	}
}
