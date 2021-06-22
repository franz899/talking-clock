package clock

var hours = []string{"midnight", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"}

// ConvertHoursToWord returns the word equivalent to the passed integer. Only works for numbers from 0 to 24
func ConvertHoursToWord(h int) string {
	if h < 13 {
		return hours[h]
	}
	return hours[h-12]
}

var minutes = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty", "twenty one", "twenty two", "twenty three", "twenty four", "twenty five", "twenty six", "twenty seven", "twenty eight", "twenty nine"}

// ConvertMinutesToWord
var halfAMinute = 30

func ConvertMinutesToWord(m int) string {
	if m > halfAMinute {
		return minutes[60-m]
	}
	return minutes[m]
}

func GetNextHour(h int) int {
	if h == 23 {
		return 0
	}
	return h + 1
}
