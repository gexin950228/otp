package utils

import (
	"net"
	"otp/models"
)

func IsIP(machine models.Machine) bool {
	var VerifyPass = true
	ip := net.ParseIP(machine.IP)
	if ip == nil {
		VerifyPass = false
	}
	if machine.Password == "" {
		VerifyPass = false
	}
	return VerifyPass
}
