package models

type Bookmark struct {
	ID        int             `json:"id" gorm:"primary_key:auto_increment"`
	User      UserResponse    `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId    int             `json:"user_id" gorm:"int"`
	JourneyId int             `json:"journey_id" gorm:"int"`
	Journey   JourneyResponse `json:"journey" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type BookmarkResponse struct {
	ID        int             `json:"id" gorm:"primary_key:auto_increment"`
	JourneyId int             `json:"journey_id" gorm:"int"`
	Journey   JourneyResponse `json:"journey"`
	User      UserResponse    `json:"user"`
	UserID    int             `json:"user_id" grom:"int"`
}

func (BookmarkResponse) TableName() string {
	return "bookmarks"
}
