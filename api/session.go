package main

import "time"

type registerEntry struct {
	Token  string
	Expiry time.Time
}
