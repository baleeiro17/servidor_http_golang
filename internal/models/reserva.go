package models

type Reserva struct {
	Id              string      `json:"id"`
	Host_images     []Baremetal `json:"host_images"`
	Ssh_public_keys []string    `json:"ssh_public_keys"`
	Status          string      `json:"status"`
	Logs            []Log       `json:"logs"`
}

type Log struct {
	Information string `json:"information"`
	Category    string `json:"category"`
	Host        string `json:"host"`
	Time        string `json:"time"`
}

type Baremetal struct {
	Host  string `json:"host"`
	Image string `json:"image"`
}
