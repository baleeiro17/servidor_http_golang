package repositories

import (
	"fmt"
	"servidor_http_golang/internal/models"
)

func GetReserva2() ([]models.Reserva, error) {

	data := make([]models.Reserva, 1)

	data[0].Id = "1"
	data[0].Status = "SCHEDULED"

	// baremetal tasks
	data[0].Host_images = make([]models.Baremetal, 2)
	data[0].Host_images[0].Host = "host1"
	data[0].Host_images[0].Image = "ubuntu_origin"
	data[0].Host_images[1].Host = "host2"
	data[0].Host_images[1].Image = "ubuntu_origin"

	// ssh tasks
	data[0].Ssh_public_keys = make([]string, 1)
	data[0].Ssh_public_keys[0] = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCz/g3zKdQRiEc3Dnb1DHnhDBFsLU+EZ4pAtD4ATBVPJxqTqB4+AQj3sYtyF/ze9sTR6aoECPe1d6PBeXzMW4uctvNzS4vtFwgPViwFMg6jieje4C5yi9XPnUqBRODaG9OGCwl000PY28WX5RFurB/DT+fvM/HcmkDjopzFh/XFBa3n0VQpl6dd5g0iXFV2mXHIqFQkJjyFPTe545dPbAyCWSV+bKRVN9homBz8iKC7WPBpe1FOVIl/1zNAKfiVHcxUwSiBD8aR0CvkoMaI7G5Z6ONFoHMOhEJHYNTH1Gr6fwcOJuRwIzkZ3rQWLo0E1TdeP+kIfhyc3lci4Izl43B2r4t9LsG86WLLGLHrNnY+85KYb3scjITxvoAC7jcJxtt429Gpa+XyiTTG6/DfBj2EDz9dQSORdjdazpa9FenHqLkk1q0n+9QLtRY638xbfVSqKYHpWsO+izqUAWo+3LZZX1Nh55JGoZOZgzmo62Pc4p48ZG18kZPx8zT362ejYYU9mfivLL8lpyUBx/MBWQyABMzG8kziWKjyDsFO4WOv+bJKnTUIPyawUIjyqX+tZ7X19MY1h8HAW9fFu6+jh/5BrcRtkuCpvL48b6qobDm7HXirzGKOW7mXp1dUdTLFYlYedFYfWqxFIfXNwhS0nuxbvLM148AyR2bF4yY4ThK9IQ== lucas@worker1"

	return data, nil
}

func UpdateReserva(id string, reserva models.Reserva) (models.Reserva, error) {

	data := models.Reserva{}
	data.Id = id
	data.Status = reserva.Status

	// number of logs
	nLog := len(reserva.Logs)
	data.Logs = make([]models.Log, nLog)
	for i := 0; i < nLog; i++ {
		data.Logs[i].Category = reserva.Logs[i].Category
		data.Logs[i].Host = reserva.Logs[i].Host
		data.Logs[i].Time = reserva.Logs[i].Time
		fmt.Println(data.Logs[i].Host)
		fmt.Println(data.Logs[i].Category)
		fmt.Println(data.Logs[i].Time)
	}

	return data, nil
}
