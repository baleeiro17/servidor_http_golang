package repositories

import "servidor_http_golang/internal/models"

func GetReserva() (models.Reserva, error) {

	data := models.Reserva{}
	data.Id = "1"
	data.Status = "SCHEDULED"
	data.Template.Metadata.Baremetal = make([]models.Baremetal, 1)
	data.Template.Metadata.Baremetal[0].Hostname = "host1"
	data.Template.Metadata.Baremetal[0].Image = "ubuntu_18_04"

	return data, nil
}

func UpdateReserva(id, status string) (models.Reserva, error) {

	data := models.Reserva{}
	data.Id = id
	data.Status = status
	data.Template.Metadata.Baremetal = make([]models.Baremetal, 1)
	data.Template.Metadata.Baremetal[0].Hostname = "host1"
	data.Template.Metadata.Baremetal[0].Image = "ubuntu_18_04"
	data.Sshkeys.Ssh_public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC9+2WATlQ21GeoLRQXag4dxjgnbwS3MBQj8P7T4+wo6J6pmfIbYebtFwAFScLx9MK6kQAe7IbYPZEpwAA8ITdRidySEVjA+Si/1Erujbfw5A0D18yocq1Ui3zvuBeTrHpXW2kX/IUYY8drdoEwmOrHC8dwOxXwsIOkf9obCzcOGUGQP/Sc61jdjJRybBs5F69aZsBmfWud5teFPgxfsUTPJ9RgaU9q5L8YmPeWwKvoyeBqoQLeu3x0SqQFZw9HW9lcuFOZ7/6HfeQCMUUPfkSQUepi6yGex8zQ5TTtq4sV+YZ3z29M8rS5Np/TlbE3doKLf7kN12+kNbjb+C8ni0mJuOXomDC/Lv9EPtOynApj/8W3+rOqw0J+H35svRfe/6O6VxhCDtQ7/B7105CCl8NA1K/xcl5e9e32/7oiPKCsTR7JZwTRwrC29ZItqKPxWVmoyShLJoeYuiliBdBuXQHIY03xDaOu5sNfq7TxS4i2z0WtyjPWVavDonepLrmhSLdoE0Jfm7Y1D+3wGZIkNPyEe64LLj6rl/ZKyzqFH21tbRjFed/DpxUal/rG71VBizAd9CPTboDBj6i9hX20VLiFTKfObx/HKYxK8fGXmDH5q0X2/FwLkQV0ZCYemPSdmzcTgRSS1nfeYU6oEA8oTlwYA5snC9twUZYSiFcpIRPKYw== lucas-baleeiro@lucasbaleeiro-Aspire-A515-52G\n"

	return data, nil
}
