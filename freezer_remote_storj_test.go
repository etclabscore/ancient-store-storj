package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func setup() {

}

func TestIntegration(t *testing.T) {

	ipcPath := os.TempDir() + "/ancient.ipc"

	testCmd := exec.Command("go", "test", "github.com/ethereum/go-ethereum/core", "-count=1", "-v", "-run=_RemoteFreezer")
	testCmd.Env = append(os.Environ(), fmt.Sprintf("GETH_ANCIENT_RPC=%s", ipcPath))
	testCmd.Stderr = os.Stderr
	err := testCmd.Run()
	if err != nil {
		panic(err)
	}
}
