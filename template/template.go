// 模板
package main

import (
	"html/template"
	"os"
)
type Person struct {
	Name string
	Age int
	Addres string
}

func main() {

	p := Person{"new1024kb", 25, "Shanghai"}
	tp := template.New("tp")
	tp.Parse("{{.Name}} {{.Age}} {{.Addres}}")
	tp.Execute(os.Stdout, p)

	// output: new1024kb 25 Shanghai
}