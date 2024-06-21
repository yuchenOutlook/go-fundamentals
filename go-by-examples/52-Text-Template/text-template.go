package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("value is {{.}}\n")
	t1.Execute(os.Stdout, "first value")
	t1.Execute(os.Stdout, "second value")
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("another value is {{.}}\n"))

	t1.Execute(os.Stdout, "Some text.")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"one", "two", "three"})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	t2 := Create("t2", "Name: {{.Name}}\n")

		
    t2.Execute(os.Stdout, struct {
        Name string
    }{"Jane Doe"})

	t2.Execute(os.Stdout, map[string]string{"Name": "Scarlett Johansson"})

	/*
		What the Code Means Within Curly Braces
	{{if . -}} yes {{else -}} no {{end}}:
	{{if .}}: Checks if the dot (.) (which represents the data passed to the template) is non-zero or non-empty.
	{{else}}: Defines the alternative text to render if the if condition is false.
	{{end}}: Ends the if action.
	The - characters are used for whitespace trimming, ensuring there are no extra spaces around the rendered output.
	*/
	t3 := Create("t3",
	"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	/*
	t4 is a template that uses the range action to iterate over a collection.
	Range: {{range .}}{{.}} {{end}}\n is the template string.
	{{range .}}: The range action iterates over the data passed to the template, which is expected to be a collection (like a slice or array).
	{{.}}: Within the range block, . represents the current item in the iteration.
	{{end}}: Ends the range action.
	*/
	t4 := Create("t4",
	"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
	[]string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	/*
	Understanding {{range .}}...{{end}}
Dot (.) in Templates:

The . (dot) is a special variable in Go templates that represents the current data context.
Initially, . refers to the data passed to the Execute method of the template.
{{range .}}:

The range action iterates over the data that . refers to.
When you use {{range .}}, it means you are iterating over the collection (slice, array, map, or channel) that was passed to the template.
Inner Dot (.) Inside {{range}}:

Inside the {{range}} block, . is redefined to refer to the current item in the iteration.
If you iterate over a slice, . will represent each element of the slice during each iteration.
{{end}}:

The {{end}} action marks the end of the {{range}} block.
	*/
}