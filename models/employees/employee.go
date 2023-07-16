package employees

type Employee struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Telepon string `json:"telepon" validate:"required"`
}
