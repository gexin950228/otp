package dataSource

import (
	"net"
	"otp/models"
)

type SearchPassword models.Machine

func SearchMachine(department, ip string) (machines []models.Machine) {
	var machineInfo models.Machine
	ips := net.ParseIP(ip)

	if ips != nil {
		if department != "" {
			Db := InitDb()
			Db.Where("ip=?", ip).Where("isDeleted=?", false).Where("department=?", department).Where("isDeleted=?", false).Find(&machineInfo, machines)
		} else {
			Db.Where("ip=?", ip).Where("isDeleted=?", false).Where("isDeleted=?", false).Find(&machineInfo)
		}
	}
	return machines
}
