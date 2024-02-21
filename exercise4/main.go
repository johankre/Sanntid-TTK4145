package main

import "os/exec"

func main() {
    cmd:= exec.Command("cmd","/C","start","c:\\windows\\system32\\wsl.exe")
    err := cmd.run()
    if err != nil {
        panic(err)
    }
}
