package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/mmatczuk/anyflag"
	"github.com/spf13/cobra"
)

func parseUserInfo(val string) (*url.Userinfo, error) {
	u, p, ok := strings.Cut(val, ":")
	if !ok {
		return nil, fmt.Errorf("invalid format")
	}
	if u == "" {
		return nil, fmt.Errorf("username is required")
	}
	if p == "" {
		return nil, fmt.Errorf("password is required")
	}
	return url.UserPassword(u, p), nil
}

func newCommand() *cobra.Command {
	var auth *url.Userinfo

	cmd := &cobra.Command{
		Use: "validation",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), auth)
		},
	}

	fs := cmd.Flags()
	value := anyflag.NewValue[*url.Userinfo](nil, &auth, parseUserInfo)

	fs.VarP(value, "auth", "", "auth")

	return cmd
}

func main() {
	newCommand().Execute()
}
