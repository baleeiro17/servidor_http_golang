package models

type Reserva struct {
	Id       string   `json:"_id"`
	Template Template `json:"template"`
	Sshkeys  []Sshkey `json:"sshkeys"`
	Status   string   `json:"status"`
	Logs     []Log    `json:"logs"`
}

type Template struct {
	Metadata Metadata `json:"metadata"`
}

type Log struct {
	Error_information string `json:"error_information"`
	Category          string `json:"category"`
	Host              string `json:"host"`
	Time              string `json:"time"`
}

type Sshkey struct {
	Ssh_public_key string `json:"ssh_public_key"`
}

type Metadata struct {
	Baremetal []Baremetal `json:"baremetal"`
}

type Baremetal struct {
	Hostname string `json:"hostname"`
	Image    string `json:"image"`
}
