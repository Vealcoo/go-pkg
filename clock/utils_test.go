package clock

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestEndOfWeek(t *testing.T) {
	// Monday 4/13
	// Friday 4/17
	// Sunday 4/19
	type args struct {
		t time.Time
	}
	endOfWeek := time.Date(2020, 4, 19, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Friday",
			args: args{
				t: time.Date(2020, 4, 17, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Monday",
			args: args{
				t: time.Date(2020, 4, 13, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Sunday",
			args: args{t: time.Date(2020, 4, 19, 0, 0, 0, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndOfWeek(tt.args.t); !reflect.DeepEqual(got, endOfWeek) {
				t.Errorf("EndOfWeek() = %v, want %v", got, endOfWeek)
			}
		})
	}
}

func TestBeginningOfWeek(t *testing.T) {
	// Monday 4/13
	// Friday 4/17
	// Sunday 4/19
	type args struct {
		t time.Time
	}
	beginOfWeek := time.Date(2020, 4, 13, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Friday",
			args: args{
				t: time.Date(2020, 4, 17, 12, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Monday",
			args: args{
				t: time.Date(2020, 4, 13, 11, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Sunday",
			args: args{t: time.Date(2020, 4, 19, 0, 0, 0, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BeginningOfWeek(tt.args.t, time.Monday); !reflect.DeepEqual(got, beginOfWeek) {
				t.Errorf("BeginningOfWeek() = %v, want %v", got, beginOfWeek)
			}
		})
	}
}

func TestModifyMonthsWithoutOverflow(t *testing.T) {
	type args struct {
		input  time.Time
		output time.Time
		gap    int
	}

	tests := []args{
		{
			input:  time.Date(2020, 1, 31, 0, 5, 0, 0, time.UTC).In(TW),
			output: time.Date(2020, 2, 29, 0, 5, 0, 0, time.UTC).In(TW),
			gap:    1,
		},
		{
			input:  time.Date(2020, 2, 29, 0, 5, 0, 0, time.UTC).In(TW),
			output: time.Date(2020, 1, 29, 0, 5, 0, 0, time.UTC).In(TW),
			gap:    -1,
		},
		{
			input:  time.Date(2020, 7, 31, 0, 5, 0, 0, time.UTC).In(TW),
			output: time.Date(2020, 6, 30, 0, 5, 0, 0, time.UTC).In(TW),
			gap:    -1,
		},
		{
			input:  time.Date(2020, 1, 31, 0, 5, 0, 0, time.UTC).In(TW),
			output: time.Date(2019, 11, 30, 0, 5, 0, 0, time.UTC).In(TW),
			gap:    -2,
		},
		{
			input:  time.Date(2019, 11, 30, 0, 5, 0, 0, time.UTC).In(TW),
			output: time.Date(2020, 1, 30, 0, 5, 0, 0, time.UTC).In(TW),
			gap:    2,
		},
	}
	for k, tt := range tests {
		t.Run(fmt.Sprintf("TestSixMonthsAgoWithoutDate %d", k+1), func(t *testing.T) {
			if got := ModifyMonthsWithoutOverflow(tt.input, tt.gap); !reflect.DeepEqual(got, tt.output) {
				t.Errorf("SixMonthsAgoWithoutDate() = %v, want %v", got, tt.output)
			}
		})
	}
}
