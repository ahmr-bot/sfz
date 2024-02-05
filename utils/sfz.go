package utils

import (
	"regexp"
)

const IDCardPattern = `[1-9]\d{14}(\d{2}[0-9X])?`

func IsValidIDCard(idCard string) bool {
	matched, err := regexp.MatchString(IDCardPattern, idCard)
	if err != nil || !matched {
		return false
	}

	var sum int
	for i, c := range idCard[:17] {
		n := int(c - '0')
		sum += n * []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}[i]
	}

	ckCodes := "10X98765432"
	return ckCodes[sum%11] == idCard[17]
}
