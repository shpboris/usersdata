package test

type TestUser struct {
	Id             string `json:"userid"`
	Name           string `json:"username"`
	Department     string `json:"department"`
	Salary         int    `json:"salary"`
	DepartmentCode string `json:"departmentcode"`
}
