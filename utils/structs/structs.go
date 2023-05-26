package structs

type Contact struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	RealAddress string `json:"realAddress"`
	Departement string `json:"department"`
	Country     string `json:"country"`
	Tel         string `json:"tel"`
	Email       string `json:"email"`
}
