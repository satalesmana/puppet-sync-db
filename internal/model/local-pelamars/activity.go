package local_pelamar

type (
	Activity struct {
		ID        string `json:"id"`
		PelamarId string `json:"pelamarId"`
		SendEmail int64  `json:"sendMail"`
		SendWa    int64  `json:"sendWa"`
	}
)
