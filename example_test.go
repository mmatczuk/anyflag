package anyflag_test

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/mmatczuk/anyflag"
	"github.com/spf13/cobra"
)

func parseBasicAuth(val string) (*url.Userinfo, error) {
	u, p, ok := strings.Cut(val, ":")
	if !ok {
		return nil, fmt.Errorf("invalid format")
	}
	return url.UserPassword(u, p), nil
}

func ExampleBasicAuth() {
	var ba *url.Userinfo
	cmd := &cobra.Command{
		Use: "example",
		Run: func(cmd *cobra.Command, args []string) {
			p, _ := ba.Password()
			fmt.Fprintln(cmd.OutOrStdout(), ba.Username(), p)
		},
	}

	cmd.Flags().VarP(anyflag.NewValue[*url.Userinfo](nil, &ba, parseBasicAuth), "basic-auth", "", "basic auth")

	cmd.SetArgs([]string{"--basic-auth", "user:pwd"})
	if err := cmd.Execute(); err != nil {
		panic(err)
	}

	// output:
	//user pwd
	//
}

func ExampleBasicAuthSlice() {
	var bas []*url.Userinfo
	cmd := &cobra.Command{
		Use: "example",
		Run: func(cmd *cobra.Command, args []string) {
			for _, ba := range bas {
				p, _ := ba.Password()
				fmt.Fprintln(cmd.OutOrStdout(), ba.Username(), p)
			}
		},
	}

	cmd.Flags().VarP(anyflag.NewSliceValue[*url.Userinfo](nil, &bas, parseBasicAuth), "basic-auth", "", "basic auth")

	cmd.SetArgs([]string{"--basic-auth", "user0:pwd0", "--basic-auth", "user1:pwd1"})
	if err := cmd.Execute(); err != nil {
		panic(err)
	}

	// output:
	//user0 pwd0
	//user1 pwd1
	//
}
