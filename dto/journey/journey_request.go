package journeydto

type CreateJourneyRequest struct {
	Title        string `json:"title"  gorm:"type: varchar(255)" form:"title" validate:"required"`
	UserId       int    `json:"user_id"  gorm:"type: int" form:"user_id"`
	Descriptions string `json:"descriptions"  gorm:"type: text" form:"descriptions" `
	Image        string `json:"image" form:"image" gorm:"type: varchar(255)" `
}
