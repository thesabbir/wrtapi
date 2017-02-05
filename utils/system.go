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

func (u Ubus) LanStatus() ([]byte, error) {
    return u("network.interface.lan status")
}

func (u Ubus) WirelessStatus() ([]byte, error) {
    return u("network.wireless status")
}

func (u Ubus) ServiceList() ([]byte, error) {
    return u("service list")
}

func (u Ubus) InterfacesList() ([]byte, error) {
    return u("network.device status")
}

func (u Ubus) NetworkConfig() ([]byte, error) {
    return u("uci get {\"config\":\"network\"}")
}