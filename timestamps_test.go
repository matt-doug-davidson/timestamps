package timestamps

import (
	"fmt"
	"testing"
)

func TestUTCZtoLocalTimestamp(t *testing.T) {
	ts := UTCZtoLocalTimestamp("2020-06-11T20:21:22.335Z")
	if ts != 1591906882335000000 {
		t.Errorf("Failed. Expected 1591906882335000000, Got %d", ts)
	}
}

func TestUTCZToUTCTimestamp(t *testing.T) {
	ts := UTCZToUTCTimestamp("2020-06-11T20:21:22.335Z")
	if ts != 1591906882335000000 {
		t.Errorf("Failed. Expected 1591906882335000000, Got %d", ts)
	}
}

func TestTimestampToUTCZTimestring(t *testing.T) {
	ts := UTCZToUTCTimestamp("2020-06-11T20:21:22.335Z")
	tStr := TimestampToUTCZTimestring(ts)
	if tStr != "2020-06-11T20:21:22.335Z" {
		t.Errorf("Expected 2020-06-11T20:21:22.335Z, got %s", tStr)
	}
}

func TestTimestampToLocalTimestring(t *testing.T) {
	ts := UTCZToUTCTimestamp("2020-06-11T20:21:22.335Z")
	tStr := TimestampToLocalTimestring(ts)
	if tStr != "2020-06-11T16:21:22.335" {
		t.Errorf("Expected 2020-06-11T16:21:22.335, got %s", tStr)
	}
}

func TestNextMinute(t *testing.T) {
	ts := NextMinuteTimestamp(5)
	tStr := TimestampToLocalTimestring(ts)
	fmt.Println("Next minute time ", tStr)
}

func TestNextHourTimestamp(t *testing.T) {
	ts := NextHourTimestamp()
	tStr := TimestampToLocalTimestring(ts)
	fmt.Println("Next  hour  time ", tStr)
}

func TestNextDayTimestamp(t *testing.T) {
	ts := NextDayTimestamp()
	tStr := TimestampToLocalTimestring(ts)
	fmt.Println("Next  day   time ", tStr)
}

func TestNanoseconds(t *testing.T) {
	ts := Nanoseconds(2.555333444)
	fmt.Println(ts)
}
