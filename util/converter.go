package util

import "strconv"

func StringToInt(value string) int {
    intValue, _ := strconv.Atoi(value)

    return intValue
}
