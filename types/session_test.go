package types

import (
	"testing"
	"time"
)

func TestUserSession(t *testing.T) {
	u := &User{
		ID:   "djawjdiajdawdji",
		Name: "aleksander",
	}

	s := u.GenerateSession("192.168.0.1", 500*time.Millisecond)

	time.Sleep(200 * time.Millisecond)

	if s.IsExpired() {
		t.FailNow()
	}
	time.Sleep(200 * time.Millisecond)

	if s.IsExpired() {
		t.FailNow()
	}
	time.Sleep(200 * time.Millisecond)

	if !s.IsExpired() {
		t.FailNow()
	}
}
