package utils

import (
	"os/exec"
	"strings"
)

type Ubus func(arg string) ([]byte, error)

var UbusCall = Ubus(func(arg string) ([]byte, error) {
	args := strings.Split(arg, " ")
	return exec.Command("ubus", args...).Output()
})

func (u Ubus) Info() ([]byte, error) {
	return u("-S call system info")
}

func (u Ubus) WanStatus() ([]byte, error) {
	return u("-S call network.interface.wan status")
}
