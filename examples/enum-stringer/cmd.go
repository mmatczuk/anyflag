package main

import (
	"fmt"
	"github.com/mmatczuk/anyflag"
	"github.com/spf13/cobra"
)

//go:generate stringer -type=Pill
type Pill int

const (
	Unspecified Pill = iota
	Placebo
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

func newCommand() *cobra.Command {
	var (
		pill Pill
		jar  []Pill
	)

	cmd := &cobra.Command{
		Use: "enum-stringer",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), pill)
			fmt.Fprintln(cmd.OutOrStdout(), jar)
		},
	}

	fs := cmd.Flags()
	parser := anyflag.StringerParser[Pill](_Pill_index[:], _Pill_name)
	value := anyflag.NewValue[Pill](Unspecified, &pill, parser)
	slice := anyflag.NewSliceValue[Pill](nil, &jar, parser)

	fs.VarP(value, "pill", "", "pill")
	fs.VarP(slice, "jar", "", "jar")

	return cmd
}

func main() {
	newCommand().Execute()
}
