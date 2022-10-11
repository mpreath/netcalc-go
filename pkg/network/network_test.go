package network

import (
	"encoding/json"
	"testing"

	"github.com/mpreath/netcalc/pkg/utils"
)

func TestMarshalJSON(t *testing.T) {
	testCases := []struct {
		ddAddress string
		ddMask    string
	}{
		{"192.168.1.1", "255.255.255.0"},
	}

	for _, testCase := range testCases {
		testAddress, err := utils.ParseAddress(testCase.ddAddress)
		if err != nil {
			t.Fatalf(err.Error())
		}
		testMask, err := utils.ParseAddress(testCase.ddMask)
		if err != nil {
			t.Fatalf(err.Error())
		}
		testNetwork, _ := New(testAddress, testMask)

		s, err := json.Marshal(testNetwork)
		if err != nil {
			t.Errorf(err.Error())
		}

		if len(s) <= 0 {
			t.Errorf("didn't receive any output from marshal")
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	testCases := []struct {
		ddAddress  string
		jsonString string
	}{
		{"192.168.1.0", "{ \"address\": \"192.168.1.0\", \"mask\": \"255.255.255.0\" }"},
	}

	for _, testCase := range testCases {
		var testNetwork Network
		err := json.Unmarshal([]byte(testCase.jsonString), &testNetwork)
		if err != nil {
			t.Errorf(err.Error())
		}

		expectedResult, err := utils.ParseAddress(testCase.ddAddress)
		if err != nil {
			t.Errorf(err.Error())
		}

		if testNetwork.Address != expectedResult {
			t.Errorf("unmarshalled address (%s) doesn't match spec address (%s)", utils.ExportAddress(testNetwork.Address), utils.ExportAddress(expectedResult))

		}
	}
}

func TestNew(t *testing.T) {
	testCases := []struct {
		ddAddress        string
		ddMask           string
		ddNetworkAddress string
		errorString      string
	}{
		{"192.168.1.1", "255.255.255.0", "192.168.1.0", ""},
		{"192.168.1.1", "255.0.255.0", "192.168.1.0", "network.New: invalid subnet mask"},
	}

	for _, testCase := range testCases {
		testAddress, err := utils.ParseAddress(testCase.ddAddress)
		if err != nil {
			t.Fatalf(err.Error())
		}
		testMask, err := utils.ParseAddress(testCase.ddMask)
		if err != nil {
			t.Fatalf(err.Error())
		}
		testNetwork, err := New(testAddress, testMask)
		if err != nil {
			if err.Error() != testCase.errorString {
				t.Errorf(err.Error())
			}
			continue
		}

		ddTestNetwork := utils.ExportAddress(testNetwork.Address)
		ddTestMask := utils.ExportAddress(testNetwork.Mask)

		if ddTestNetwork != testCase.ddNetworkAddress {
			t.Errorf("generated address (%s) doesn't match spec address (%s)", ddTestNetwork, testCase.ddNetworkAddress)
		}

		if ddTestMask != testCase.ddMask {
			t.Errorf("generated mask (%s) doesn't match spec mask (%s)", ddTestMask, testCase.ddMask)
		}
	}
}

func TestHosts(t *testing.T) {
	testCases := []struct {
		ddAddress string
		ddMask    string
		hostCount int
	}{
		{"192.168.1.0", "255.255.255.0", 254},
		{"192.168.1.0", "255.255.255.128", 126},
	}

	for _, testCase := range testCases {
		testAddress, err := utils.ParseAddress(testCase.ddAddress)
		if err != nil {
			t.Fatalf(err.Error())
		}
		testMask, err := utils.ParseAddress(testCase.ddMask)
		if err != nil {
			t.Fatalf(err.Error())
		}
		testNetwork, err := New(testAddress, testMask)
		if err != nil {
			t.Fatalf(err.Error())
		}

		testHosts := testNetwork.Hosts()

		if len(testHosts) != testCase.hostCount {
			t.Errorf("generated host count (%d) doesn't match spec count (%d)", len(testHosts), testCase.hostCount)
		}
	}
}

func TestHostCount(t *testing.T) {
	testCases := []struct {
		ddAddress string
		ddMask    string
		hostCount int
	}{
		{"192.168.1.0", "255.255.255.0", 254},
		{"192.168.1.0", "255.255.255.128", 126},
	}

	for _, testCase := range testCases {
		testAddress, err := utils.ParseAddress(testCase.ddAddress)
		if err != nil {
			t.Fatalf(err.Error())
		}
		testMask, err := utils.ParseAddress(testCase.ddMask)
		if err != nil {
			t.Fatalf(err.Error())
		}
		testNetwork, err := New(testAddress, testMask)
		if err != nil {
			t.Fatalf(err.Error())
		}

		testHostCount := testNetwork.HostCount()

		if testHostCount != testCase.hostCount {
			t.Errorf("generated host count (%d) doesn't match spec count (%d)", testHostCount, testCase.hostCount)
		}
	}
}
