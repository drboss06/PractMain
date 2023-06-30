package authPract

type Team struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UserTeam struct {
	Id     int
	UserId int
	TeamId int
}
