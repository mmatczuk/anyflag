package main

func Example() {
	cmd := newCommand()
	cmd.SetArgs([]string{"--auth", "user:pass"})
	cmd.Execute()
	cmd.SetArgs([]string{"--auth", "user:"})
	cmd.Execute()
	cmd.SetArgs([]string{"--auth", ":pass"})
	cmd.Execute()

	// output:
	// user:pass
	//
}
