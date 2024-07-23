package local_pelamar

type (
	Activity struct {
		Id         string `json:"_id"`
		DatabaseId string `json:"database_id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		Phone      string `json:"phone"`
		IsMail     string `gorm:"-" json:"isMail"`
		IsPhone    string `gorm:"-" json:"isPhone"`
	}
)
