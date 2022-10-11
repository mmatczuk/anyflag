package main

import (
	"fmt"
	"github.com/mmatczuk/anyflag"
	"github.com/spf13/cobra"
)

type Pill string

const (
	Unspecified   Pill = ""
	Placebo       Pill = "Placebo"
	Aspirin       Pill = "Aspirin"
	Ibuprofen     Pill = "Ibuprofen"
	Paracetamol   Pill = "Paracetamol"
	Acetaminophen      = Paracetamol
)

func (p Pill) String() string {
	return string(p)
}

func newCommand() *cobra.Command {
	var (
		pill Pill
		jar  []Pill
	)

	cmd := &cobra.Command{
		Use: "enum",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), pill)
			fmt.Fprintln(cmd.OutOrStdout(), jar)
		},
	}

	fs := cmd.Flags()
	parser := anyflag.EnumParser[Pill](Placebo, Aspirin, Ibuprofen, Paracetamol)
	value := anyflag.NewValue[Pill](Unspecified, &pill, parser)
	slice := anyflag.NewSliceValue[Pill](nil, &jar, parser)

	fs.VarP(value, "pill", "", "pill")
	fs.VarP(slice, "jar", "", "jar")

	return cmd
}

func main() {
	newCommand().Execute()
}
