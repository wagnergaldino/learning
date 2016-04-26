package hello

/*HelloWorld greets you or the world if you don't have a name.*/
func HelloWorld(name string) string {
	if name == "" {
		name = "Hello, World"
	}
	return "Hello, " + name + "!"
}
