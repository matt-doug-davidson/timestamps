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
	fmt.Println(utcTime)
	tStr := utcTime.Format("2006-01-02T15:04:05.000Z")
	fmt.Println(tStr)
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

func (t *Timestamp) toGoTimestring() string {
	unixTime := time.Unix(t.seconds, t.nanoSeconds) //gives unix time stamp in utc
	fmt.Println(unixTime)
	return ""
}
func (t *SecondTimestamp) ToGoTimestring() string {
	return t.ts.toGoTimestring()
}

func (t *MillisecondTimestamp) ToGoTimestring() string {
	return t.ts.toGoTimestring()
}

func (t *NanosecondTimestamp) ToGoTimestring() string {
	return t.ts.toGoTimestring()
}
