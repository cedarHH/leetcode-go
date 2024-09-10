package __strings

import (
	"fmt"
)

func replaceNumber(strByte []byte) string {
	numCount, oldSize := 0, len(strByte)
	for i := 0; i < len(strByte); i++ {
		if (strByte[i] <= '9') && (strByte[i] >= '0') {
			numCount++
		}
	}

	for i := 0; i < numCount; i++ {
		strByte = append(strByte, []byte("     ")...)
	}
	tmpBytes := []byte("number")

	leftP, rightP := oldSize-1, len(strByte)-1
	for leftP < rightP {
		rightShift := 1

		if (strByte[leftP] <= '9') && (strByte[leftP] >= '0') {
			for i, tmpByte := range tmpBytes {
				strByte[rightP-len(tmpBytes)+i+1] = tmpByte
			}
			rightShift = len(tmpBytes)
		} else {
			strByte[rightP] = strByte[leftP]
		}
		rightP -= rightShift
		leftP -= 1
	}
	return string(strByte)
}

func ReplaceNumber() {
	strByte := []byte("q1q1q")

	newString := replaceNumber(strByte)

	fmt.Println(newString)
}
