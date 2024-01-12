package main

import "strconv"

func ConvertToDigitalFormat(hour, minute, second int) string {
	hourStr, minuteStr, secondStr := strconv.Itoa(hour), strconv.Itoa(minute), strconv.Itoa(second)
	if len(hourStr) == 1 {
		hourStr = "0" + hourStr
	}
	if len(minuteStr) == 1 {
		minuteStr = "0" + minuteStr
	}
	if len(secondStr) == 1 {
		secondStr = "0" + secondStr
	}
	return hourStr + ":" + minuteStr + ":" + secondStr
}

func ExtractTimeUnits(seconds int) (int, int, int) {
	var hour, minute, second int
	hour = seconds / 3600
	minute = (seconds - 3600*hour) / 60
	second = seconds % 60
	return hour, minute, second
}
