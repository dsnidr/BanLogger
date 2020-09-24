package strdur

import (
	"testing"
)

func TestGetDuration(t *testing.T) {
	var expected int64

	// Minute tests
	expected = 60
	seconds := GetDuration("1min")

	if seconds != expected {
		t.Errorf("GetDuration call with 1min should have outputted %d seconds, but instead output: %d\n", expected, seconds)
	}

	expected = (7363 * secondsInMinute)
	seconds = GetDuration("7363min")

	if seconds != expected {
		t.Errorf("GetDuration call with 7363min should have outputted %d seconds, but instead output: %d\n", expected, seconds)
	}

	// Hour test
	expected = secondsInHour
	seconds = GetDuration("1h")

	if seconds != expected {
		t.Errorf("GetDuration call with 1h should have outputted %d seconds, but instead output: %d\n", expected, seconds)
	}

	expected = (120 * secondsInHour)
	seconds = GetDuration("120h")

	if seconds != expected {
		t.Errorf("GetDuration call with 120h should have outputted %d seconds, but instead output: %d\n", expected, seconds)
	}

	// Day test
	expected = secondsInDay
	seconds = GetDuration("1d")

	if seconds != expected {
		t.Errorf("GetDuration call with 1d should have outputted %d seconds, but instead output: %d\n", expected, seconds)
	}

	expected = (490 * secondsInDay)
	seconds = GetDuration("490d")

	if seconds != expected {
		t.Errorf("GetDuration call with 490d should have outputted %d seconds, but instead output: %d\n", expected, seconds)
	}

	// Week test
	expected = secondsInWeek
	seconds = GetDuration("1w")

	if seconds != expected {
		t.Errorf("GetDuration call with 1w should have outputted %d seconds, but instead output: %d\n", expected, seconds)
	}

	expected = (283 * secondsInWeek)
	seconds = GetDuration("283w")

	if seconds != expected {
		t.Errorf("GetDuration call with 283w should have outputted %d seconds, but instead output: %d\n", expected, seconds)
	}

	// Month test
	expected = secondsInMonth
	seconds = GetDuration("1m")

	if seconds != expected {
		t.Errorf("GetDuration call with 1m should have outputted %d seconds, but instead output: %d\n", expected, seconds)
	}

	expected = (192 * secondsInMonth)
	seconds = GetDuration("192m")

	if seconds != expected {
		t.Errorf("GetDuration call with 192m should have outputted %d seconds, but instead output: %d\n", expected, seconds)
	}

	// Year test
	expected = secondsInYear
	seconds = GetDuration("1y")

	if seconds != expected {
		t.Errorf("GetDuration call with 1y should have outputted %d seconds, but instead output: %d\n", expected, seconds)
	}

	expected = (420 * secondsInYear)
	seconds = GetDuration("420y")

	if seconds != expected {
		t.Errorf("GetDuration call with 420d should have outputted %d seconds, but instead output: %d\n", expected, seconds)
	}
}
