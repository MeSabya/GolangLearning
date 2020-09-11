package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Emails []string
}

const tmpl = `The name is {{.Name}},
{{range .Emails}}
	His email id is {{.}}
{{end}}
`

func main() {
	person := Person{
		Name:   "Sabya",
		Emails: []string{"sabya@gmail.com", "sabyasachi@gmail.com"},
	}

	//This will create a new template with a given name
	t := template.New("Person template")

	//Parse will parse a string into template
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse error")
		return
	}

	//“Execute” applies a parsed template to the specified data object,
	//and writes the output to “os.Stdout”.
	//t is the parsed template ,execute applies parsed template to person
	//Here basically we need to remember two methods , one is
	//Parse another is Execute ....
	err = t.Execute(os.Stdout, person)
	if err != nil {
		log.Fatal("Execution error", err)
		return
	}

}
