package dataSource

import "otp/models"

func GetMachine(machine models.Machine) []models.Machine {
	conditionMap := map[string]string{}
	var machines []models.Machine
	if machine.IP != "" {
		conditionMap["IP"] = machine.IP
	}
	if machine.Maintainer != "" {
		conditionMap["maintainer"] = machine.Maintainer
	}
	if machine.Department != "" {
		conditionMap["department"] = machine.Department
	}

	return machines
}
