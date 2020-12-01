package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type Cpu struct {
	Cpus  []float64 `json:"cpus"`
	Cores int       `json:"cores"`
}

func InitCPU() (c Cpu, err error) {
	if cores, err := cpu.Counts(false); err != nil {
		return c, err
	} else {
		c.Cores = cores
	}
	if cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return c, err
	} else {
		c.Cpus = cpus
	}
	return c, nil
}
