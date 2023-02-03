package main

func Example() {
	cmd := newCommand()
	cmd.SetArgs([]string{"--site", "https://user:password@www.google.com", "--realm", "http://foo.com", "--realm", "http://u:p@bar.com"})
	cmd.Execute()

	// Output:
	// realm [http://foo.com,http://u:xxxxx@bar.com]
	// site https://user:xxxxx@www.google.com
	//
}
