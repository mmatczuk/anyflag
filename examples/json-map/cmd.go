package main

import (
	"encoding/json"
	"fmt"

	"github.com/mmatczuk/anyflag"
	"github.com/spf13/cobra"
)

func parseJSONMap(val string) (map[string]interface{}, error) {
	var m map[string]interface{}
	return m, json.Unmarshal([]byte(val), &m)
}

func newCommand() *cobra.Command {
	var m map[string]interface{}

	cmd := &cobra.Command{
		Use: "json-map",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), m)
		},
	}

	fs := cmd.Flags()
	value := anyflag.NewValue[map[string]interface{}](nil, &m, parseJSONMap)

	fs.VarP(value, "map", "", "map")

	return cmd
}

func main() {
	newCommand().Execute()
}
