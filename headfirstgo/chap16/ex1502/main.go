package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Part struct {
	Name  string
	Count int
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	text := "Here's my template!\n"
	teml, err := template.New("test").Parse(text)
	check(err)
	err = teml.Execute(os.Stdout, nil)
	check(err)
	templateText := "Template start\nAction: {{.}}\nTemplate end\n"
	teml, err = template.New("test").Parse(templateText)
	check(err)
	err = teml.Execute(os.Stdout, "ABC")
	check(err)
	err = teml.Execute(os.Stdout, 42)
	check(err)
	err = teml.Execute(os.Stdout, true)
	check(err)
	fmt.Println("----------------------")
	executeTemplate("Dot is : {{.}}!\n", "ABC")
	executeTemplate("Dot is : {{.}}!\n", 123.5)
	fmt.Println("----------------------")
	executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", true)
	executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", false)
	fmt.Println("----------------------")

	templateText = "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"
	executeTemplate(templateText, []string{"do", "re", "mi"})
	fmt.Println("----------------------")

	templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	executeTemplate(templateText, []float64{1.25, 0.99, 27})
	executeTemplate(templateText, []float64{})
	executeTemplate(templateText, nil)
	fmt.Println("----------------------")

	templateText = "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(templateText, Part{Name: "Fuses", Count: 5})
	executeTemplate(templateText, Part{Name: "Cables", Count: 2})
	fmt.Println("----------------------")

	templateText = "Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	subscriber := Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	executeTemplate(templateText, subscriber)
	subscriber = Subscriber{Name: "Joy Carr", Rate: 5.99, Active: false}
	executeTemplate(templateText, subscriber)
}
