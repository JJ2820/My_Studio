package models

import "time"

type Booking struct {
	MemberName string `json:"member_name"`
	Date       string `json:"date"`
}

func (b *Booking) IsValid() bool {
	_, err := time.Parse("2006-01-02", b.Date)
	return b.MemberName != "" && err == nil
}
