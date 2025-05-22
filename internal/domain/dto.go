package domain

type AddActivityInput struct {
	UserID      int64
	SkillID     int64
	Type        string
	Title       string
	Description string
	XP          int64
}
