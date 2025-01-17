package helper

var version = 1.1
var Application = "Golang"

// fungsi ini tidak bisa diakses dari luar package
func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}