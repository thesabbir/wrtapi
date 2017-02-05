package utils

import (
    "os/exec"
    "strings"
)

type Ubus func(arg string) ([]byte, error)

var UbusCall = Ubus(func(arg string) ([]byte, error) {
    args := strings.Split(arg, " ")
    args = append([]string{"-S", "call"}, args...)
    return exec.Command("ubus", args...).Output()
})

func (u Ubus) Info() ([]byte, error) {
    return u("system info")
}

func (u Ubus) BoardInfo() ([]byte, error) {
    return u("system board")
}

func (u Ubus) WanStatus() ([]byte, error) {
    return u("network.interface.wan status")
}

func (u Ubus) WirelessStatus() ([]byte, error) {
    return u("network.wireless status")
}