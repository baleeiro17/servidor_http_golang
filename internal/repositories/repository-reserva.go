package repositories

import "servidor_http_golang/internal/models"

func GetReserva() (models.Reserva, error) {

	data := models.Reserva{}
	data.Id = "1"
	data.Status = "SCHEDULED"
	data.Template.Metadata.Baremetal.Hostname = "host1"
	data.Template.Metadata.Baremetal.Image = "ubuntu_18_04"

	return data, nil
}

func UpdateReserva(id, status string) (models.Reserva, error) {

	data := models.Reserva{}
	data.Id = id
	data.Status = status
	data.Template.Metadata.Baremetal.Hostname = "host1"
	data.Template.Metadata.Baremetal.Image = "ubuntu_18_04"

	return data, nil
}
