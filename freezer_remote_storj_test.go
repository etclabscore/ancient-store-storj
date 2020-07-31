package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"testing"

	"github.com/ethereum/go-ethereum/cmd/utils"
)

var (
	ipcPath = os.TempDir() + "/ancient.ipc"
)

func runMain(bucket string) {
	fmt.Println(ipcPath)
	os.Args = append([]string{"./ancient-store-storj", "--bucket", fmt.Sprintf("unit-test-%v", rand.Int()), "--loglevel", "3", "--ipcpath", ipcPath})
	main()

}

func TestIntegration(t *testing.T) {
	defer os.RemoveAll(ipcPath)
	defer stopServer()
	go func() {
		runMain("etclabs-storj-freezer-test")
	}()
	testCmd := exec.Command("go", "test", "github.com/ethereum/go-ethereum/core", "-count", "1", "-v", "-run", "_RemoteFreezer")
	testCmd.Env = os.Environ()
	testCmd.Env = append(os.Environ(), fmt.Sprintf("GETH_ANCIENT_RPC=%s", ipcPath))
	testCmd.Stderr = os.Stderr
	testCmd.Stdout = os.Stdout
	err := testCmd.Run()
	if err != nil {
		utils.Fatalf("Unexpected error", err)
	}
}
