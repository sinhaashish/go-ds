package strings

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParseString(dateInput string) (inputtime string, err error) {

	var validID = regexp.MustCompile(`^(\d+d)?\s*((?:[01]?\d|2[0-3])h)?\s*((?:[0-5]?\d)m)?$`)
	fmt.Println(validID.MatchString(inputtime))

	// var err error
	fmt.Println("\n The input string is ", dateInput)
	//var hours := ""
	message := dateInput
	var numberOfDays, numberOfHours, numberOfMinutes int
	if days := strings.IndexByte(message, 'd'); days >= 0 {
		numberOfDays, err = GetStringInBetween(message, "", "d")
		fmt.Println(numberOfDays)
		if err != nil {
			fmt.Print(" A ", err)
			return "", err
		}
	}
	if hours := strings.IndexByte(message, 'h'); hours >= 0 {
		startchar := ""
		if days := strings.IndexByte(message, 'd'); days >= 0 {
			startchar = "d"
		}
		numberOfHours, err = GetStringInBetween(message, startchar, "h")
		fmt.Println(numberOfHours)
		if err != nil {
			fmt.Print(" B ", err)
			return "", err
		}
	}
	// if minutes := strings.IndexByte(message, 'm'); minutes >= 0 {
	// 	startchar := ""
	// 	if days := strings.IndexByte(message, 'd'); days >= 0 {
	// 		if hours := strings.IndexByte(message, 'h'); hours >= 0 {
	// 			startchar = "h"
	// 		} else {
	// 			startchar = "d"
	// 		}
	// 	} else {
	// 		if hours := strings.IndexByte(message, 'h'); hours >= 0 {
	// 			startchar = "h"
	// 		}
	// 	}
	// 	numberOfMinutes, err = GetStringInBetween(message, startchar, "m")
	// 	fmt.Println(numberOfMinutes)
	// 	if err != nil {
	// 		fmt.Print(" C ", err)
	// 		return "", err
	// 	}
	// }
	fmt.Printf(" \n numberOfDays :", numberOfDays)
	fmt.Printf(" \n numberOfHours :", numberOfHours)
	// fmt.Printf(" \n numberOfMinutes :", numberOfMinutes)
	test := strconv.Itoa(24*numberOfDays + numberOfHours)
	fmt.Print(test)

	s := []string{strconv.Itoa(24*numberOfDays + numberOfHours), "h", strconv.Itoa(numberOfMinutes), "m"}
	outputTime := strings.Join(s, "")

	//outputTime := string(24*numberOfDays+numberOfHours) + "h" + string(numberOfMinutes) + "m"
	fmt.Printf("\n The output time is %#v \n ", outputTime)

	return outputTime, nil

}

func GetStringInBetween(input, start, end string) (result int, err error) {
	var output string
	startIndex := strings.Index(input, start)

	if startIndex == -1 {
		return
	}
	startIndex += len(start)
	fmt.Println(" startIndex ", startIndex)

	endIndex := strings.Index(input, end)
	fmt.Println(" endIndex ", endIndex)

	if endIndex > startIndex {
		output = input[startIndex:endIndex]
		err = nil
	} else {
		output = ""
		err = fmt.Errorf(" Error in parsing ")
	}
	fmt.Print(" output", output)
	fmt.Print(" error ", err)
	value, err := strconv.Atoi(output)
	return value, err
}

type Duration int64

var errLeadingInt = errors.New("time: bad [0-9]*") // never printed

// leadingInt consumes the leading [0-9]* from s.
func leadingInt(s string) (x int64, rem string, err error) {
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		if x > (1<<63-1)/10 {
			// overflow
			return 0, "", errLeadingInt
		}
		x = x*10 + int64(c) - '0'
		if x < 0 {
			// overflow
			return 0, "", errLeadingInt
		}
	}
	return x, s[i:], nil
}

var unitMap = map[string]int64{
	"ns": int64(Nanosecond),
	"us": int64(Microsecond),
	"µs": int64(Microsecond), // U+00B5 = micro symbol
	"μs": int64(Microsecond), // U+03BC = Greek letter mu
	"ms": int64(Millisecond),
	"s":  int64(Second),
	"m":  int64(Minute),
	"h":  int64(Hour),
	"d":  int64(Day),
}

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
	Day                  = 24 * Hour
)

// leadingFraction consumes the leading [0-9]* from s.
// It is used only for fractions, so does not return an error on overflow,
// it just stops accumulating precision.
func leadingFraction(s string) (x int64, scale float64, rem string) {
	i := 0
	scale = 1
	overflow := false
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		if overflow {
			continue
		}
		if x > (1<<63-1)/10 {
			// It's possible for overflow to give a positive number, so take care.
			overflow = true
			continue
		}
		y := x*10 + int64(c) - '0'
		if y < 0 {
			overflow = true
			continue
		}
		x = y
		scale *= 10
	}
	return x, scale, s[i:]
}

// ParseDuration parses a duration string.
// A duration string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix,
// such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func ParseDuration(s string) (Duration, error) {
	// [-+]?([0-9]*(\.[0-9]*)?[a-z]+)+
	orig := s
	var d int64
	neg := false

	// Consume [-+]?
	if s != "" {
		c := s[0]
		if c == '-' || c == '+' {
			neg = c == '-'
			s = s[1:]
		}
	}
	// Special case: if all that is left is "0", this is zero.
	if s == "0" {
		return 0, nil
	}
	if s == "" {
		return 0, errors.New("time: invalid duration " + orig)
	}
	for s != "" {
		var (
			v, f  int64       // integers before, after decimal point
			scale float64 = 1 // value = v + f/scale
		)

		var err error

		// The next character must be [0-9.]
		if !(s[0] == '.' || '0' <= s[0] && s[0] <= '9') {
			return 0, errors.New("time: invalid duration " + orig)
		}
		// Consume [0-9]*
		pl := len(s)
		v, s, err = leadingInt(s)
		if err != nil {
			return 0, errors.New("time: invalid duration " + orig)
		}
		pre := pl != len(s) // whether we consumed anything before a period

		// Consume (\.[0-9]*)?
		post := false
		if s != "" && s[0] == '.' {
			s = s[1:]
			pl := len(s)
			f, scale, s = leadingFraction(s)
			post = pl != len(s)
		}
		if !pre && !post {
			// no digits (e.g. ".s" or "-.s")
			return 0, errors.New("time: invalid duration " + orig)
		}

		// Consume unit.
		i := 0
		for ; i < len(s); i++ {
			c := s[i]
			if c == '.' || '0' <= c && c <= '9' {
				break
			}
		}
		if i == 0 {
			return 0, errors.New("time: missing unit in duration " + orig)
		}
		u := s[:i]
		s = s[i:]
		unit, ok := unitMap[u]
		if !ok {
			return 0, errors.New("time: unknown unit " + u + " in duration " + orig)
		}
		if v > (1<<63-1)/unit {
			// overflow
			return 0, errors.New("time: invalid duration " + orig)
		}
		v *= unit
		if f > 0 {
			// float64 is needed to be nanosecond accurate for fractions of hours.
			// v >= 0 && (f*unit/scale) <= 3.6e+12 (ns/h, h is the largest unit)
			v += int64(float64(f) * (float64(unit) / scale))
			if v < 0 {
				// overflow
				return 0, errors.New("time: invalid duration " + orig)
			}
		}
		d += v
		if d < 0 {
			// overflow
			return 0, errors.New("time: invalid duration " + orig)
		}
	}

	if neg {
		d = -d
	}
	return Duration(d), nil
}
