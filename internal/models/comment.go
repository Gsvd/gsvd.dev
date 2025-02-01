package models

import "time"

type Comment struct {
	Id                 int
	Username           string
	PostId             int
	Comment            string
	Approved           bool
	CreatedAtFormatted string
	CreatedAt          time.Time
}
