package clock

import (
	"time"
)

func Real() Contract {
	return &realClock{}
}

type realClock struct {
}

func (r realClock) Now() time.Time {
	return time.Now()
}

func (r realClock) TwNow() time.Time {
	return time.Now().In(TW)
}
