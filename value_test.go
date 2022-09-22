package anyflag

import (
	"fmt"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/spf13/pflag"
)

func parseIP(val string) (net.IP, error) {
	ip := net.ParseIP(strings.TrimSpace(val))
	if ip == nil {
		return nil, fmt.Errorf("failed to parse IP: %q", val)
	}
	return ip, nil
}

func setUpIP(ip *net.IP) *pflag.FlagSet {
	f := pflag.NewFlagSet("test", pflag.ContinueOnError)
	f.VarP(NewValue[net.IP](net.ParseIP("0.0.0.0"), ip, parseIP), "address", "", "IP address")
	return f
}

func TestIP(t *testing.T) {
	testCases := []struct {
		input    string
		success  bool
		expected string
	}{
		{"0.0.0.0", true, "0.0.0.0"},
		{" 0.0.0.0", true, "0.0.0.0"},
		{"1.2.3.4", true, "1.2.3.4"},
		{"127.0.0.1", true, "127.0.0.1"},
		{"255.255.255.255", true, "255.255.255.255"},
		{"", false, ""},
		{"0", false, ""},
		{"localhost", false, ""},
		{"0.0.0", false, ""},
		{"0.0.0.", false, ""},
		{"0.0.0.0.", false, ""},
		{"0.0.0.256", false, ""},
		{"0 . 0 . 0 . 0", false, ""},
	}

	devnull, _ := os.Open(os.DevNull)
	os.Stderr = devnull
	for i := range testCases {
		var addr net.IP
		f := setUpIP(&addr)

		tc := &testCases[i]

		arg := fmt.Sprintf("--address=%s", tc.input)
		err := f.Parse([]string{arg})
		if err != nil && tc.success == true {
			t.Errorf("expected success, got %q", err)
			continue
		} else if err == nil && tc.success == false {
			t.Errorf("expected failure")
			continue
		} else if tc.success {
			if addr.String() != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, addr.String())
			}
		}
	}
}
