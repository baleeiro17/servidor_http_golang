package repositories

import (
	"fmt"
	"servidor_http_golang/internal/models"
)

func GetReserva() (models.Reserva, error) {

	data := models.Reserva{}
	data.Id = "1"
	data.Status = "SCHEDULED"
	data.Template.Metadata.Baremetal = make([]models.Baremetal, 2)
	data.Template.Metadata.Baremetal[0].Hostname = "host1"
	data.Template.Metadata.Baremetal[0].Image = "ubuntu_18_04"
	data.Template.Metadata.Baremetal[1].Hostname = "host2"
	data.Template.Metadata.Baremetal[1].Image = "ubuntu_generic"
	data.Sshkeys.Ssh_public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCz/g3zKdQRiEc3Dnb1DHnhDBFsLU+EZ4pAtD4ATBVPJxqTqB4+AQj3sYtyF/ze9sTR6aoECPe1d6PBeXzMW4uctvNzS4vtFwgPViwFMg6jieje4C5yi9XPnUqBRODaG9OGCwl000PY28WX5RFurB/DT+fvM/HcmkDjopzFh/XFBa3n0VQpl6dd5g0iXFV2mXHIqFQkJjyFPTe545dPbAyCWSV+bKRVN9homBz8iKC7WPBpe1FOVIl/1zNAKfiVHcxUwSiBD8aR0CvkoMaI7G5Z6ONFoHMOhEJHYNTH1Gr6fwcOJuRwIzkZ3rQWLo0E1TdeP+kIfhyc3lci4Izl43B2r4t9LsG86WLLGLHrNnY+85KYb3scjITxvoAC7jcJxtt429Gpa+XyiTTG6/DfBj2EDz9dQSORdjdazpa9FenHqLkk1q0n+9QLtRY638xbfVSqKYHpWsO+izqUAWo+3LZZX1Nh55JGoZOZgzmo62Pc4p48ZG18kZPx8zT362ejYYU9mfivLL8lpyUBx/MBWQyABMzG8kziWKjyDsFO4WOv+bJKnTUIPyawUIjyqX+tZ7X19MY1h8HAW9fFu6+jh/5BrcRtkuCpvL48b6qobDm7HXirzGKOW7mXp1dUdTLFYlYedFYfWqxFIfXNwhS0nuxbvLM148AyR2bF4yY4ThK9IQ== lucas@worker1\n"

	return data, nil
}

func UpdateReserva(id, status string) (models.Reserva, error) {

	data := models.Reserva{}
	data.Id = id
	data.Status = status
	data.Template.Metadata.Baremetal = make([]models.Baremetal, 1)
	data.Template.Metadata.Baremetal[0].Hostname = "host1"
	data.Template.Metadata.Baremetal[0].Image = "ubuntu_18_04"
	fmt.Println(status)

	return data, nil
}
