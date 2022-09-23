package main

func Example() {
	cmd := newCommand()
	cmd.SetArgs([]string{"--site", "https://www.google.com?q=foo", "--realm", "http://foo.com", "--realm", "http://bar.com"})
	cmd.Execute()

	// output:
	// https://www.google.com?q=foo
	// [http://foo.com http://bar.com]
	//
}
