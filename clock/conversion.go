package clock

var hours = []string{"midnight", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"}

func convertHoursToWord(h int) string {
	if h < 13 {
		return hours[h]
	}
	return hours[h-12]
}

var minutes = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty", "twenty one", "twenty two", "twenty three", "twenty four", "twenty five", "twenty six", "twenty seven", "twenty eight", "twenty nine"}

var halfAMinute = 30

func convertMinutesToWord(m int) string {
	if m > halfAMinute {
		return minutes[60-m]
	}
	return minutes[m]
}

func getNextHour(h int) int {
	if h == 23 {
		return 0
	}
	return h + 1
}
