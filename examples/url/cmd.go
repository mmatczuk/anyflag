package main

import (
	"fmt"
	"github.com/mmatczuk/anyflag"
	"github.com/spf13/cobra"
	"net/url"
)

func newCommand() *cobra.Command {
	var (
		site  *url.URL
		realm []*url.URL
	)

	cmd := &cobra.Command{
		Use: "url",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), site)
			fmt.Fprintln(cmd.OutOrStdout(), realm)
		},
	}

	fs := cmd.Flags()
	value := anyflag.NewValue[*url.URL](nil, &site, url.ParseRequestURI)
	slice := anyflag.NewSliceValue[*url.URL](nil, &realm, url.ParseRequestURI)

	fs.VarP(value, "site", "", "site")
	fs.VarP(slice, "realm", "", "realm")

	return cmd
}

func main() {
	newCommand().Execute()
}
