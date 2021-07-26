package repositories

import (
	"fmt"
	"servidor_http_golang/internal/models"
)

func GetReserva() (models.Reserva, error) {

	data := models.Reserva{}

	data.Result = make([]models.Result, 2)

	data.Result[0].Id = "1"
	data.Result[0].Status = "SCHEDULED"

	// baremetal tasks
	data.Result[0].Host_images = make([]models.Baremetal, 1)
	data.Result[0].Host_images[0].Host = "60f5b2e38298c5795c147e06"
	data.Result[0].Host_images[0].Image = "60d22b62e8c38ebe76abfb38"

	// ssh tasks
	data.Result[0].Ssh_public_keys = make([]string, 1)
	data.Result[0].Ssh_public_keys[0] = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCz/g3zKdQRiEc3Dnb1DHnhDBFsLU+EZ4pAtD4ATBVPJxqTqB4+AQj3sYtyF/ze9sTR6aoECPe1d6PBeXzMW4uctvNzS4vtFwgPViwFMg6jieje4C5yi9XPnUqBRODaG9OGCwl000PY28WX5RFurB/DT+fvM/HcmkDjopzFh/XFBa3n0VQpl6dd5g0iXFV2mXHIqFQkJjyFPTe545dPbAyCWSV+bKRVN9homBz8iKC7WPBpe1FOVIl/1zNAKfiVHcxUwSiBD8aR0CvkoMaI7G5Z6ONFoHMOhEJHYNTH1Gr6fwcOJuRwIzkZ3rQWLo0E1TdeP+kIfhyc3lci4Izl43B2r4t9LsG86WLLGLHrNnY+85KYb3scjITxvoAC7jcJxtt429Gpa+XyiTTG6/DfBj2EDz9dQSORdjdazpa9FenHqLkk1q0n+9QLtRY638xbfVSqKYHpWsO+izqUAWo+3LZZX1Nh55JGoZOZgzmo62Pc4p48ZG18kZPx8zT362ejYYU9mfivLL8lpyUBx/MBWQyABMzG8kziWKjyDsFO4WOv+bJKnTUIPyawUIjyqX+tZ7X19MY1h8HAW9fFu6+jh/5BrcRtkuCpvL48b6qobDm7HXirzGKOW7mXp1dUdTLFYlYedFYfWqxFIfXNwhS0nuxbvLM148AyR2bF4yY4ThK9IQ== lucas@worker1"

	data.Result[1].Id = "2"
	data.Result[1].Status = "SCHEDULED"

	// baremetal tasks
	data.Result[1].Host_images = make([]models.Baremetal, 1)
	data.Result[1].Host_images[0].Host = "60f5b2e38298c5795c147e05"
	data.Result[1].Host_images[0].Image = "60d22b62e8c38ebe76abfb37"

	// ssh tasks
	data.Result[1].Ssh_public_keys = make([]string, 1)
	data.Result[1].Ssh_public_keys[0] = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCz/g3zKdQRiEc3Dnb1DHnhDBFsLU+EZ4pAtD4ATBVPJxqTqB4+AQj3sYtyF/ze9sTR6aoECPe1d6PBeXzMW4uctvNzS4vtFwgPViwFMg6jieje4C5yi9XPnUqBRODaG9OGCwl000PY28WX5RFurB/DT+fvM/HcmkDjopzFh/XFBa3n0VQpl6dd5g0iXFV2mXHIqFQkJjyFPTe545dPbAyCWSV+bKRVN9homBz8iKC7WPBpe1FOVIl/1zNAKfiVHcxUwSiBD8aR0CvkoMaI7G5Z6ONFoHMOhEJHYNTH1Gr6fwcOJuRwIzkZ3rQWLo0E1TdeP+kIfhyc3lci4Izl43B2r4t9LsG86WLLGLHrNnY+85KYb3scjITxvoAC7jcJxtt429Gpa+XyiTTG6/DfBj2EDz9dQSORdjdazpa9FenHqLkk1q0n+9QLtRY638xbfVSqKYHpWsO+izqUAWo+3LZZX1Nh55JGoZOZgzmo62Pc4p48ZG18kZPx8zT362ejYYU9mfivLL8lpyUBx/MBWQyABMzG8kziWKjyDsFO4WOv+bJKnTUIPyawUIjyqX+tZ7X19MY1h8HAW9fFu6+jh/5BrcRtkuCpvL48b6qobDm7HXirzGKOW7mXp1dUdTLFYlYedFYfWqxFIfXNwhS0nuxbvLM148AyR2bF4yY4ThK9IQ== lucas@worker1"

	return data, nil
}

func GetReserva2() (models.Reserva, error) {

	data := models.Reserva{}

	data.Result = make([]models.Result, 1)

	data.Result[0].Id = "1"
	data.Result[0].Status = "SCHEDULED"

	// baremetal tasks
	data.Result[0].Host_images = make([]models.Baremetal, 2)
	data.Result[0].Host_images[0].Host = "60f5b2e38298c5795c147e06"
	data.Result[0].Host_images[0].Image = "60d22b62e8c38ebe76abfb38"
	data.Result[0].Host_images[1].Host = "60f5b2e38298c5795c147e05"
	data.Result[0].Host_images[1].Image = "60d22b62e8c38ebe76abfb37"

	// ssh tasks
	data.Result[0].Ssh_public_keys = make([]string, 1)
	data.Result[0].Ssh_public_keys[0] = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCz/g3zKdQRiEc3Dnb1DHnhDBFsLU+EZ4pAtD4ATBVPJxqTqB4+AQj3sYtyF/ze9sTR6aoECPe1d6PBeXzMW4uctvNzS4vtFwgPViwFMg6jieje4C5yi9XPnUqBRODaG9OGCwl000PY28WX5RFurB/DT+fvM/HcmkDjopzFh/XFBa3n0VQpl6dd5g0iXFV2mXHIqFQkJjyFPTe545dPbAyCWSV+bKRVN9homBz8iKC7WPBpe1FOVIl/1zNAKfiVHcxUwSiBD8aR0CvkoMaI7G5Z6ONFoHMOhEJHYNTH1Gr6fwcOJuRwIzkZ3rQWLo0E1TdeP+kIfhyc3lci4Izl43B2r4t9LsG86WLLGLHrNnY+85KYb3scjITxvoAC7jcJxtt429Gpa+XyiTTG6/DfBj2EDz9dQSORdjdazpa9FenHqLkk1q0n+9QLtRY638xbfVSqKYHpWsO+izqUAWo+3LZZX1Nh55JGoZOZgzmo62Pc4p48ZG18kZPx8zT362ejYYU9mfivLL8lpyUBx/MBWQyABMzG8kziWKjyDsFO4WOv+bJKnTUIPyawUIjyqX+tZ7X19MY1h8HAW9fFu6+jh/5BrcRtkuCpvL48b6qobDm7HXirzGKOW7mXp1dUdTLFYlYedFYfWqxFIfXNwhS0nuxbvLM148AyR2bF4yY4ThK9IQ== lucas@worker1"

	return data, nil
}

func UpdateReserva(id string, reserva models.Result) (models.Result, error) {

	data := models.Result{}

	data.Id = id
	data.Status = reserva.Status
	fmt.Println("status: ", reserva.Status)

	// number of logs
	nLog := len(reserva.Logs)
	data.Logs = make([]models.Log, nLog)
	for i := 0; i < nLog; i++ {
		data.Logs[i].Category = reserva.Logs[i].Category
		data.Logs[i].Resource = reserva.Logs[i].Resource
		data.Logs[i].Time = reserva.Logs[i].Time
		fmt.Println("resource:", data.Logs[i].Resource)
		fmt.Println("category:", data.Logs[i].Category)
		fmt.Println("time:", data.Logs[i].Time)
	}

	return data, nil
}
