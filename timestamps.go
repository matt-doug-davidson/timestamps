package timestamps

import (
	"fmt"
	"strconv"
	"time"
)

// Timestamp structure is the underlying structure of the other
// timestamp types
type Timestamp struct {
	seconds     int64
	nanoSeconds int64
}

// SecondTimestamp is a timestamp that contains seconds
type SecondTimestamp struct {
	ts Timestamp
}

// MillisecondTimestamp is a timestamp that contains seconds
type MillisecondTimestamp struct {
	ts Timestamp
}

// NanosecondTimestamp is a timestamp that contains seconds
type NanosecondTimestamp struct {
	ts Timestamp
}

// Set the underlying Timestamp with a value input as seconds.
func (t *SecondTimestamp) Set(seconds int64) {
	t.ts.seconds = seconds
	t.ts.nanoSeconds = 0
}

// Set the underlying Timestamp with a value input as milliseconds.
func (t *MillisecondTimestamp) Set(milliseconds int64) {
	t.ts.seconds = milliseconds / 1000
	t.ts.nanoSeconds = (milliseconds % 1000) * 1000000
}

// Set the underlying Timestamp with a value input as nanoseconds.
func (t *NanosecondTimestamp) Set(nanoseconds int64) {
	t.ts.seconds = nanoseconds / 1000000000
	t.ts.nanoSeconds = (nanoseconds % 1000000000)
}

// FromString Set the underlying Timestamp with a stream input as seconds.
func (t *SecondTimestamp) FromString(value string) {
	t.ts.seconds, _ = strconv.ParseInt(value, 10, 64)
	t.ts.nanoSeconds = 0
}

// FromString Set the underlying Timestamp with a stream input as milliseconds.
func (t *MillisecondTimestamp) FromString(value string) {
	milliseconds, _ := strconv.ParseInt(value, 10, 64)
	t.Set(milliseconds)
}

// FromString Set the underlying Timestamp with a stream input as nanoseconds.
func (t *NanosecondTimestamp) FromString(value string) {
	nanoseconds, _ := strconv.ParseInt(value, 10, 64)
	t.Set(nanoseconds)
}

func (t *Timestamp) print() {
	fmt.Printf("Timestamp seconds: %d\n", t.seconds)
	fmt.Printf("Timestamp nanoseconds: %d\n", t.nanoSeconds)
	fmt.Printf("%d.%09d\n", t.seconds, t.nanoSeconds)
}

// Print outputs the underlying Timestamp
func (t *SecondTimestamp) Print() {
	t.ts.print()
}

// Print outputs the underlying Timestamp
func (t *MillisecondTimestamp) Print() {
	t.ts.print()
}

// Print outputs the underlying Timestamp
func (t *NanosecondTimestamp) Print() {
	t.ts.print()
}

// Convert the underlying Timestamp to a UTCZ string.
func (t *Timestamp) toUTCZTimestamp() string {
	got := time.Unix(t.seconds, t.nanoSeconds)
	utcLoc, _ := time.LoadLocation("UTC")
	utcTime := got.In(utcLoc)
	tStr := utcTime.Format("2006-01-02T15:04:05.000Z")
	return tStr
}

// ToUTCZ outputs a string in UTCZ format. The timestamp must already
// be set.
func (t *SecondTimestamp) ToUTCZ() string {
	return t.ts.toUTCZTimestamp()
}

// ToUTCZ outputs a string in UTCZ format. The timestamp must already
// be set.
func (t *MillisecondTimestamp) ToUTCZ() string {
	return t.ts.toUTCZTimestamp()
}

// ToUTCZ outputs a string in UTCZ format. The timestamp must already
// be set.
func (t *NanosecondTimestamp) ToUTCZ() string {
	return t.ts.toUTCZTimestamp()
}

// ConvertUTCZ outputs a string in UTCZ format. It sets the
// value based on value string.
func (t *SecondTimestamp) ConvertUTCZ(value string) string {
	t.FromString(value)
	return t.ts.toUTCZTimestamp()
}

// ConvertUTCZ outputs a string in UTCZ format. It sets the
// value based on value string.
func (t *MillisecondTimestamp) ConvertUTCZ(value string) string {
	t.FromString(value)
	return t.ts.toUTCZTimestamp()
}

// ConvertUTCZ outputs a string in UTCZ format. It sets the
// value based on value string.
func (t *NanosecondTimestamp) ConvertUTCZ(value string) string {
	t.FromString(value)
	return t.ts.toUTCZTimestamp()
}

// Convert the timestamp to a time.Time object
func (t *Timestamp) ToGoTime() time.Time {
	unixTime := time.Unix(t.seconds, t.nanoSeconds) //gives unix time stamp in utc
	return unixTime
}

func (t *SecondTimestamp) ToGoTime() time.Time {
	return t.ts.ToGoTime()
}

func (t *MillisecondTimestamp) ToGoTime() time.Time {
	return t.ts.ToGoTime()
}

func (t *NanosecondTimestamp) ToGoTime() time.Time {
	return t.ts.ToGoTime()
}

// NextHourTimestamp returns the next hour timestamp in nanoseconds
func NextHourTimestamp() int64 {
	now := time.Now().UnixNano()
	// Get the current hour in timestamp
	hour := int64(60 * 60 * 1000000000)
	thisHour := (now / int64(hour))
	// Determin next hour
	nextHour := thisHour + 1
	// Convert back to nanoseconds
	nextHourNs := nextHour * hour
	return nextHourNs
}

// NextDayTimestamp returns the next hour timestamp in nanoseconds
func NextDayTimestamp() int64 {
	now := time.Now()
	// Convert to to just date part
	dateT := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	// Add one day to create a new time object
	future := dateT.AddDate(0, 0, 1)
	// Convert to timestamp
	timestamp := future.UnixNano()
	return timestamp
}

// NextMinuteTimestamp returns the next minute mark timestamp in nanoseconds
func NextMinuteTimestamp(interval int64) int64 {
	// Get current times
	seconds := time.Now().Unix()
	// Current mark
	currentMinute := seconds / 60                             // minutes
	minuteOfHour := currentMinute % 60                        // minutes
	currentMarkOfHour := (minuteOfHour / interval) * interval // minutes
	// Next Mark
	nextMarkOfHour := currentMarkOfHour + interval
	nextMarkSeconds := currentMinute + nextMarkOfHour - minuteOfHour
	nextMarkNanoSecs := nextMarkSeconds * 60000000000
	return nextMarkNanoSecs
}

// UTCZtoLocalTimestamp convert a UTCZ string to a local timestamp.
func UTCZtoLocalTimestamp(utczString string) int64 {
	t, err := time.Parse("2006-01-02T15:04:05.000Z", utczString)
	if err != nil {
		return -1
	}

	lt := t.Local()

	return lt.UnixNano()
}

// UTCZToUTCTimestamp converts UTCZ timestring to UTC timestamp.
func UTCZToUTCTimestamp(utczString string) int64 {
	t, err := time.Parse("2006-01-02T15:04:05.000Z", utczString)

	if err != nil {
		return -1
	}
	return t.UnixNano()
}

func TimestampToLocalTimestring(ts int64) string {

	nsts := NanosecondTimestamp{}
	nsts.Set(ts)
	t := nsts.ToGoTime()
	lt := t.Local()
	return lt.Format("2006-01-02T15:04:05.000")
}

func TimestampToTimestringNoMilli(ts int64) string {
	nsts := NanosecondTimestamp{}
	nsts.Set(ts)
	t := nsts.ToGoTime()
	lt := t.Local()
	return lt.Format("2006-01-02T15:04:05")
}

func TimestampToUTCZTimestring(ts int64) string {
	nsts := NanosecondTimestamp{}
	nsts.Set(ts)
	return nsts.ToUTCZ()
}

// UTCTimestamp returns the UTC timestamp in nanoseconds
func UTCTimestamp() int64 {
	return time.Now().UTC().UnixNano()
}

func LocalTimestamp() int64 {
	return time.Now().UnixNano()
}

// Nanoseconds converts timestamp in seconds to one in nanoseconds
func Nanoseconds(tSec float64) int64 {
	t := tSec * 1000000000
	return int64(t)
}

// RoundDownMinutes rounds a timestamp to the minute
func RoundDownMinutes(tNSec int64) int64 {
	return (tNSec / 60000000000) * 60000000000
}

func MinutesEarlier(tNSec int64, minutes int64) int64 {
	// Minutes to nanoseconds
	minTs := minutes * 1000000000 * 60
	return tNSec - minTs
}

type TimestringPair struct {
	Begin, End string
}

// GetTimespansFromTimestrings
func GetTimespansFromTimestrings(start string, stop string, spanDurationMinutes int64) []TimestringPair {
	span := spanDurationMinutes * 60 * 1000000000
	//span += 1000000
	fmt.Println(span)

	// Convert stop to stop to timestamps
	startTs := UTCZToUTCTimestamp(start)
	if startTs == -1 {
		fmt.Println("Error: start time ", start, " conversion failed.")
		return nil
	}
	stopTs := UTCZToUTCTimestamp(stop)
	if stopTs == -1 {
		fmt.Println("Error: stop time ", stop, " conversion failed.")
		return nil
	}
	fmt.Println(startTs, stopTs)
	spanEnd := startTs + span
	fmt.Println(TimestampToUTCZTimestring(spanEnd))
	timespans := []TimestringPair{}
	if startTs+span < stopTs {
		for beginTs := startTs; beginTs < stopTs; beginTs += span {
			endTs := beginTs + span
			if endTs != stopTs {
				endTs -= 1000000
				if endTs > stopTs {
					endTs = stopTs
				}
			}
			fmt.Println(beginTs, endTs)
			fmt.Println(TimestampToUTCZTimestring(beginTs))
			fmt.Println(TimestampToUTCZTimestring(endTs))
			timespans = append(timespans,
				TimestringPair{TimestampToUTCZTimestring(beginTs),
					TimestampToUTCZTimestring(endTs)})
		}
	} else {

	}
	return timespans
}

// IsTimeFirstMinuteOfHour determines if the time provided falls in the first
// minute of the hour.
func IsTimeFirstMinuteOfHour(t time.Time) bool {
	return t.Minute() == 0
}

// IsNanoTimestampFirstMinuteOfHour determines if the nano timestamp provided
// falls in the first minute of the hour.
func IsNanoTimestampFirstMinuteOfHour(t int64) bool {
	minutes := t / 60000000000
	return minutes%60 == 0
}

// AddMinutes adds specified minutes to a timestamp.
func AddMinutes(ts int64, minutes int64) int64 {
	return ts + minutes*60000000000

}

// SubtractMinutes subtracts specfied minutes from a timestamp.
func SubtractMinutes(ts int64, minutes int64) int64 {
	return ts - minutes*60000000000
}
