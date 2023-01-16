package employeemodel

type Employee struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Phone    string `json:"fullname"`
	Role     string `json:"role"`
}

type EmployeeCreateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Role     string `json:"role"`
}
