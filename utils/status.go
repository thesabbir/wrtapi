package utils

import (
    "os/exec"
    "strings"
)

func ubusCall(arg string) ([]byte, error) {
    args := strings.Split(arg, " ")
    return exec.Command("ubus", args...).Output()
}

func StatusWan() ([]byte, error) {
    return ubusCall("-S call network.interface.wan status")
}

func SystemInfo() ([]byte, error) {
    return ubusCall("-S call system info")
}