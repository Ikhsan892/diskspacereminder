package diskspace

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

func GetMemory() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	if v.UsedPercent > 70 {
		fmt.Println("Please clean up the data")
	}

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}
