package utils

import (
	"math"
	"strconv"
	"strings"
)

var tenToAny =  [62]string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P", "A", "S", "D", "F", "G", "H", "J", "K", "L", "Z", "X", "C", "V", "B", "N", "M"};




// 10进制转任意进制
var modNum int64= 62
func Ten_To_62(num int64) string {
	new_num_str := ""
	var remainder int64
	var remainder_string string
	for num != 0 {
		remainder = num%modNum
		if 76 > remainder && remainder > 9 {
			remainder_string = tenToAny[remainder]
		} else {
			remainder_string = strconv.FormatInt(remainder,10)
		}
		new_num_str = remainder_string + new_num_str
		num = num / modNum
	}
	return new_num_str
}

func findKey(in string) int {
	result := -1
	for k, v := range tenToAny {
		if in == v {
			result = k
			break
		}
	}
	return result
}

// 62进制转10进制
func Sixty_two_To_10(num string) int {
	var new_num float64
	new_num = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(findKey(value))
		if tmp != -1 {
			new_num = new_num + tmp * math.Pow(float64(62), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(new_num)
}
