package clock

import (
	"time"
)

func Fake(now time.Time) Contract {
	return &fakeClock{
		now: now,
	}
}

type fakeClock struct {
	now time.Time
}

func (f fakeClock) Now() time.Time {
	return f.now
}

func (f fakeClock) TwNow() time.Time {
	return f.now
}
