package repositories

import (
	"fmt"
	"servidor_http_golang/internal/models"
)

func GetReserva() ([]models.Reserva, error) {

	data := make([]models.Reserva, 2)

	data[0].Id = "1"
	data[0].Status = "SCHEDULED"

	// baremetal task
	data[0].Template.Metadata.Baremetal = make([]models.Baremetal, 1)
	data[0].Template.Metadata.Baremetal[0].Hostname = "host1"
	data[0].Template.Metadata.Baremetal[0].Image = "ubuntu_18_04"

	// ssh task
	data[0].Sshkeys = make([]models.Sshkey, 1)
	data[0].Sshkeys[0].Ssh_public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCz/g3zKdQRiEc3Dnb1DHnhDBFsLU+EZ4pAtD4ATBVPJxqTqB4+AQj3sYtyF/ze9sTR6aoECPe1d6PBeXzMW4uctvNzS4vtFwgPViwFMg6jieje4C5yi9XPnUqBRODaG9OGCwl000PY28WX5RFurB/DT+fvM/HcmkDjopzFh/XFBa3n0VQpl6dd5g0iXFV2mXHIqFQkJjyFPTe545dPbAyCWSV+bKRVN9homBz8iKC7WPBpe1FOVIl/1zNAKfiVHcxUwSiBD8aR0CvkoMaI7G5Z6ONFoHMOhEJHYNTH1Gr6fwcOJuRwIzkZ3rQWLo0E1TdeP+kIfhyc3lci4Izl43B2r4t9LsG86WLLGLHrNnY+85KYb3scjITxvoAC7jcJxtt429Gpa+XyiTTG6/DfBj2EDz9dQSORdjdazpa9FenHqLkk1q0n+9QLtRY638xbfVSqKYHpWsO+izqUAWo+3LZZX1Nh55JGoZOZgzmo62Pc4p48ZG18kZPx8zT362ejYYU9mfivLL8lpyUBx/MBWQyABMzG8kziWKjyDsFO4WOv+bJKnTUIPyawUIjyqX+tZ7X19MY1h8HAW9fFu6+jh/5BrcRtkuCpvL48b6qobDm7HXirzGKOW7mXp1dUdTLFYlYedFYfWqxFIfXNwhS0nuxbvLM148AyR2bF4yY4ThK9IQ== lucas@worker1\n"

	// new json with other host
	data[1].Id = "2"
	data[1].Status = "SCHEDULED"

	// baremetal task
	data[1].Template.Metadata.Baremetal = make([]models.Baremetal, 1)
	data[1].Template.Metadata.Baremetal[0].Hostname = "host2"
	data[1].Template.Metadata.Baremetal[0].Image = "ubuntu_generic"

	// ssh task
	data[1].Sshkeys = make([]models.Sshkey, 1)
	data[1].Sshkeys[0].Ssh_public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCz/g3zKdQRiEc3Dnb1DHnhDBFsLU+EZ4pAtD4ATBVPJxqTqB4+AQj3sYtyF/ze9sTR6aoECPe1d6PBeXzMW4uctvNzS4vtFwgPViwFMg6jieje4C5yi9XPnUqBRODaG9OGCwl000PY28WX5RFurB/DT+fvM/HcmkDjopzFh/XFBa3n0VQpl6dd5g0iXFV2mXHIqFQkJjyFPTe545dPbAyCWSV+bKRVN9homBz8iKC7WPBpe1FOVIl/1zNAKfiVHcxUwSiBD8aR0CvkoMaI7G5Z6ONFoHMOhEJHYNTH1Gr6fwcOJuRwIzkZ3rQWLo0E1TdeP+kIfhyc3lci4Izl43B2r4t9LsG86WLLGLHrNnY+85KYb3scjITxvoAC7jcJxtt429Gpa+XyiTTG6/DfBj2EDz9dQSORdjdazpa9FenHqLkk1q0n+9QLtRY638xbfVSqKYHpWsO+izqUAWo+3LZZX1Nh55JGoZOZgzmo62Pc4p48ZG18kZPx8zT362ejYYU9mfivLL8lpyUBx/MBWQyABMzG8kziWKjyDsFO4WOv+bJKnTUIPyawUIjyqX+tZ7X19MY1h8HAW9fFu6+jh/5BrcRtkuCpvL48b6qobDm7HXirzGKOW7mXp1dUdTLFYlYedFYfWqxFIfXNwhS0nuxbvLM148AyR2bF4yY4ThK9IQ== lucas@worker1\n"

	return data, nil
}

func GetReserva2() (models.Reserva, error) {

	data := models.Reserva{}

	data.Id = "1"
	data.Status = "SCHEDULED"

	// baremetal tasks
	data.Template.Metadata.Baremetal = make([]models.Baremetal, 2)
	data.Template.Metadata.Baremetal[0].Hostname = "host1"
	data.Template.Metadata.Baremetal[0].Image = "ubuntu_18_04"
	data.Template.Metadata.Baremetal[1].Hostname = "host2"
	data.Template.Metadata.Baremetal[1].Image = "ubuntu_generic"

	// ssh tasks
	data.Sshkeys = make([]models.Sshkey, 1)
	data.Sshkeys[0].Ssh_public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCz/g3zKdQRiEc3Dnb1DHnhDBFsLU+EZ4pAtD4ATBVPJxqTqB4+AQj3sYtyF/ze9sTR6aoECPe1d6PBeXzMW4uctvNzS4vtFwgPViwFMg6jieje4C5yi9XPnUqBRODaG9OGCwl000PY28WX5RFurB/DT+fvM/HcmkDjopzFh/XFBa3n0VQpl6dd5g0iXFV2mXHIqFQkJjyFPTe545dPbAyCWSV+bKRVN9homBz8iKC7WPBpe1FOVIl/1zNAKfiVHcxUwSiBD8aR0CvkoMaI7G5Z6ONFoHMOhEJHYNTH1Gr6fwcOJuRwIzkZ3rQWLo0E1TdeP+kIfhyc3lci4Izl43B2r4t9LsG86WLLGLHrNnY+85KYb3scjITxvoAC7jcJxtt429Gpa+XyiTTG6/DfBj2EDz9dQSORdjdazpa9FenHqLkk1q0n+9QLtRY638xbfVSqKYHpWsO+izqUAWo+3LZZX1Nh55JGoZOZgzmo62Pc4p48ZG18kZPx8zT362ejYYU9mfivLL8lpyUBx/MBWQyABMzG8kziWKjyDsFO4WOv+bJKnTUIPyawUIjyqX+tZ7X19MY1h8HAW9fFu6+jh/5BrcRtkuCpvL48b6qobDm7HXirzGKOW7mXp1dUdTLFYlYedFYfWqxFIfXNwhS0nuxbvLM148AyR2bF4yY4ThK9IQ== lucas@worker1\n"

	return data, nil
}

func GetReserva3() (models.Reserva, error) {

	data := models.Reserva{}

	data.Id = "1"
	data.Status = "SCHEDULED"

	// baremetal tasks
	data.Template.Metadata.Baremetal = make([]models.Baremetal, 1)
	data.Template.Metadata.Baremetal[0].Hostname = "host1"
	data.Template.Metadata.Baremetal[0].Image = "ubuntu_18_04"

	// ssh task
	data.Sshkeys = make([]models.Sshkey, 2)
	data.Sshkeys[0].Ssh_public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCz/g3zKdQRiEc3Dnb1DHnhDBFsLU+EZ4pAtD4ATBVPJxqTqB4+AQj3sYtyF/ze9sTR6aoECPe1d6PBeXzMW4uctvNzS4vtFwgPViwFMg6jieje4C5yi9XPnUqBRODaG9OGCwl000PY28WX5RFurB/DT+fvM/HcmkDjopzFh/XFBa3n0VQpl6dd5g0iXFV2mXHIqFQkJjyFPTe545dPbAyCWSV+bKRVN9homBz8iKC7WPBpe1FOVIl/1zNAKfiVHcxUwSiBD8aR0CvkoMaI7G5Z6ONFoHMOhEJHYNTH1Gr6fwcOJuRwIzkZ3rQWLo0E1TdeP+kIfhyc3lci4Izl43B2r4t9LsG86WLLGLHrNnY+85KYb3scjITxvoAC7jcJxtt429Gpa+XyiTTG6/DfBj2EDz9dQSORdjdazpa9FenHqLkk1q0n+9QLtRY638xbfVSqKYHpWsO+izqUAWo+3LZZX1Nh55JGoZOZgzmo62Pc4p48ZG18kZPx8zT362ejYYU9mfivLL8lpyUBx/MBWQyABMzG8kziWKjyDsFO4WOv+bJKnTUIPyawUIjyqX+tZ7X19MY1h8HAW9fFu6+jh/5BrcRtkuCpvL48b6qobDm7HXirzGKOW7mXp1dUdTLFYlYedFYfWqxFIfXNwhS0nuxbvLM148AyR2bF4yY4ThK9IQ== lucas@worker1\n"
	data.Sshkeys[1].Ssh_public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDZlWO0LD2uCqj50Oc8yOCeAWNE0lfqm+vurahoaG01UZgvtdyzng9eFEtHtCpat90zG4bUMRDN/LUhZRJwEAkHpkD0sV4VhlP35VXucSz6GkSHqiqVZLdStL5CMev1mFOTHcjkVv9bN78Ax5iiw9A9hIBhly+bwkN8W3yibKKqL6WYi2Gb+SmO+F34Jyxg0eyQhVGH3MTFmYU4F6bT1lr0sj+OG8NaQ/HpjEoX2X0zDB0y4VMsujcrhkZX0p54iCeBkfTWOW5MBdYt9lxtMQW4zMqmn7+RNRGPNMsdLOnoZvlHA78jYBtyDRY0sLBGF7M5iTWr441YeOn8NJ2awQyhukcxFzwkbg9missx1fb03ljQ5f16k50OtAdW/0+s7p3dEkG1gJRngJeG5ujyjbaymYvFYudO8yTyW/520MMbdlNxLtrsram9sO8PeN1hpJWQY9DUf8hPpw1P9Li9WEJMCCFvWHMIXweXYtXjOYXvSfOpgkqdVAevubjYpSF3ODgHigjMF4/4BQKd01ekKMsUHzCrJZ8dn7AVOhMZQoBcR22EcW0VFykOufSKVF2gFn3Cu1TmL+DGvA/F65WkeiFI5Emu1xWHBnTzK4vp9Pf0kukVgn0VmX0lQ8+GQx+1KxM5xgAYtMnm3LmiNcvkLqZZgzCa+gw52e4UVhi5+sHwTw== lucas@lucas-Standard-PC-Q35-ICH9-2009"

	return data, nil
}

func UpdateReserva(id, status string) (models.Reserva, error) {

	data := models.Reserva{}
	data.Id = id
	data.Status = status
	data.Template.Metadata.Baremetal = make([]models.Baremetal, 1)
	data.Template.Metadata.Baremetal[0].Hostname = "host1"
	data.Template.Metadata.Baremetal[0].Image = "ubuntu_18_04"

	fmt.Println(id)
	fmt.Println(status)
	fmt.Println(status)

	return data, nil
}
