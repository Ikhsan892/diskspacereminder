package diskspace

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"strconv"
)

func GetCPU() {

	percentage, _ := cpu.Percent(0, true)

	for idx, cpupercent := range percentage {
		fmt.Println("Current CPU utilization: [" + strconv.Itoa(idx) + "] " + strconv.FormatFloat(cpupercent, 'f', 2, 64))
	}

}
