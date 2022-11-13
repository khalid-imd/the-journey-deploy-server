package journeydto

type JourneyResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	UserId      int    `json:"user_id"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
