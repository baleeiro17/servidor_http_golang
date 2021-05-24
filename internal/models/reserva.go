package models

type Reserva struct {
	Id       string   `json:"_id"`
	Template Template `json:"template"`
	Status   string   `json:"status"`
}

type Template struct {
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Baremetal []Baremetal `json:"baremetal"`
}

type Baremetal struct {
	Hostname string `json:"hostname"`
	Image    string `json:"image"`
}
