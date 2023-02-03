package main

import (
	"fmt"
	"net/url"

	"github.com/mmatczuk/anyflag"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func newCommand() *cobra.Command {
	var (
		site  *url.URL
		realm []*url.URL
	)

	cmd := &cobra.Command{
		Use: "redact",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Flags().Visit(func(f *pflag.Flag) {
				fmt.Fprintln(cmd.OutOrStdout(), f.Name, f.Value)
			})
		},
	}

	fs := cmd.Flags()

	redact := func(u *url.URL) string {
		return u.Redacted()
	}

	value := anyflag.NewValueWithRedact[*url.URL](nil, &site, url.ParseRequestURI, redact)
	slice := anyflag.NewSliceValueWithRedact[*url.URL](nil, &realm, url.ParseRequestURI, redact)

	fs.VarP(value, "site", "", "site")
	fs.VarP(slice, "realm", "", "realm")

	return cmd
}

func main() {
	newCommand().Execute()
}
