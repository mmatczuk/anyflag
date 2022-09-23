package main

func Example() {
	cmd := newCommand()
	cmd.SetArgs([]string{"--map", `{"foo": "bar"}`})
	cmd.Execute()

	// output:
	// map[foo:bar]
	//
}
