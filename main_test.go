package main

import (
	"time"
	"testing"
)

func SetNow(t time.Time){
	nowFunc = func() time.Time { return t }
}

func TestNewNow(t *testing.T) {
	cases := []struct {
		t time.Time
	}{
		{ time.Date(2018, time.January, 10, 18, 0, 0, 0, time.UTC)},
		{ time.Date(2018, time.January, 10, 19, 0, 0, 0, time.UTC)},
	}

	for _, c := range cases {
		SetNow(c.t)
		now := NewNow()
		if c.t != now {
			t.Errorf("got %s, want %s", now, c.t)
		}
	}
}