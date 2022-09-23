package main

func Example() {
	cmd := newCommand()
	cmd.SetArgs([]string{"--pill", "Placebo", "--jar", "Aspirin", "--jar", "Ibuprofen"})
	cmd.Execute()

	// output:
	// Placebo
	// [Aspirin Ibuprofen]
	//
}
