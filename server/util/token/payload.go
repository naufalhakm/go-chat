package token

import "time"

type Token struct {
	AuthID  int
	Expired time.Time
}
