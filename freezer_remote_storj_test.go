package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"testing"
)

var (
	apiKey    = ""
	secret    = ""
	satellite = "us-central-1.tardigrade.io:7777"
	ipcPath   = os.TempDir() + "ancient.ipc"
)

func runMain(bucket string) {
	os.Setenv("STORJ_API_KEY", apiKey)
	os.Setenv("STORJ_SECRET", secret)
	os.Setenv("STORJ_SATELLITE", satellite)

	os.Args = append([]string{"./ancient-store-storj", "--bucket", fmt.Sprintf("unit-test-%v", rand.Int()), "--loglevel", "3", "--ipcpath", ipcPath})
	main()

}

func getTestCases() []string {
	testCmd := exec.Command("go", "test", "github.com/ethereum/go-ethereum/core", "-list", "_RemoteFreezer")
	testCmd.Stderr = os.Stderr
	output, err := testCmd.Output()
	if err != nil {
		panic(err)
	}
	cases := strings.Split(string(output), "\n")
	parsedCases := []string{}
	for _, c := range cases {
		if strings.HasPrefix(c, "Test") {
			parsedCases = append(parsedCases, c)
		}
	}
	return parsedCases
}

func runTestCase(testCase string) {
	if "TestTransactionIndices_RemoteFreezer" == testCase {
		return
	}
	go func() {
		runMain(testCase)
	}()
	testFilter := fmt.Sprintf("-run=%s", testCase)
	fmt.Println("TESTCASE===================:", testCase)
	testCmd := exec.Command("go", "test", "github.com/ethereum/go-ethereum/core", "-count=1", "-v", testFilter)
	testCmd.Env = os.Environ()
	testCmd.Env = append(os.Environ(), fmt.Sprintf("GETH_ANCIENT_RPC=%s", ipcPath))
	testCmd.Stderr = os.Stderr
	//	testCmd.Stdout = os.Stdout
	err := testCmd.Run()
	if err != nil {
		panic(err)
	}
}

func TestIntegration(t *testing.T) {

	testCases := getTestCases()
	for _, testCase := range testCases {
		runTestCase(testCase)
	}
}
