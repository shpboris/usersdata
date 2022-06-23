package usersdata

type User struct {
	Id     string `json:"userid"`
	Name   string `json:"username"`
	Unit   string `json:"unit"`
	Salary int    `json:"salary"`
}
