package clock

import "time"

type Contract interface {
	Now() time.Time
	TwNow() time.Time
}
