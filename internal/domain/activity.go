package domain

import "time"

type Activity struct {
	ID          int64
	UserID      int64
	SkillID     int64
	Type        string
	Title       string
	Description string
	XP          int64
	CreatedAt   time.Time
}
