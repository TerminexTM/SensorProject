package util

import "time"

type IDateChecker interface {
	CheckDateTimeDuration(from, to time.Time, duration time.Duration) bool
	CheckDateBeforeThresold(date time.Time, duration time.Duration) bool
}

type dateChecker struct{}

func NewDateChecker() IDateChecker {
	return dateChecker{}
}

func (d dateChecker) CheckDateTimeDuration(from, to time.Time, duration time.Duration) bool {
	return to.Sub(from) <= duration
}

func (d dateChecker) CheckDateBeforeThresold(date time.Time, duration time.Duration) bool {
	return time.Since(date) <= duration
}
