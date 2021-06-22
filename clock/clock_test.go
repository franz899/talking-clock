package clock

import "testing"

func TestTalk(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"00:00", "Midnight"},
		{"00:10", "Ten past midnight"},
		{"00:30", "Half past midnight"},
		{"1:00", "One o'clock"},
		{"2:00", "Two o'clock"},
		{"11:45", "Fifteen to twelve"},
		{"12:00", "Noon"},
		{"13:00", "One o'clock"},
		{"13:05", "Five past one"},
		{"13:10", "Ten past one"},
		{"13:25", "Twenty five past one"},
		{"13:30", "Half past one"},
		{"13:35", "Twenty five to two"},
		{"13:55", "Five to two"},
		{"14:38", "Twenty two to three"},
		{"17:15", "Fifteen past five"},
		{"23:35", "Twenty five to midnight"},
	}

	for _, c := range cases {
		got, _ := Talk(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
