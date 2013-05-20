package main

import "os"
import "os/exec"

//Unbuffer stdin
var _ = exec.Command("/bin/stty", "-F", "/dev/tty", "-icanon", "min", "1" ).Run()

//Single character buffer to read to
var getChBuff = make([]byte, 1)

func getCh() (byte, error){
    ct,err := os.Stdin.Read(getChBuff)
    if ct != 0  {
        return getChBuff[0], nil
    }
    return 0, err
}


