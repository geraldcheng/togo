package greetings

import (
	"time"
)

func IsMorning(t time.Time) bool {
	if t.Hour() <= 12 || t.Hour() >= 6 {
		return true
	} else {
		return false
	}
}

func IsAfternoon(t time.Time) bool {
	if t.Hour() <= 18 || t.Hour() > 12 {
		return true
	} else {
		return false
	}
}

func IsEvening(t time.Time) bool {
	if t.Hour() <= 24 || t.Hour() > 6 {
		return true
	} else {
		return false
	}
}
