package mkprj

import (
	"bytes"
	"os"
	"regexp"
	"text/template"
)

func loadFile(tmplFile string) (map[string]func(interface{}) string, error) {
	f, err := os.ReadFile(tmplFile)
	if err != nil {
		return nil, err
	}
	tmpl := string(f)
	re := regexp.MustCompile("(?U)SQLStatement: (\\w+)\n((?sm).+)(--|\\z)")
	fm := make(map[string]func(interface{}) string)
	for _, match := range re.FindAllStringSubmatch(tmpl, -1) {
		name := match[1]
		stmnt := match[2]
		t := template.Must(template.New(name).Parse(stmnt))
		fm[name] = func(input interface{}) string {
			var doc bytes.Buffer
			t.Execute(&doc, input)
			return doc.String()
		}
	}
	return fm, nil
}

func mkprj() error {
	return nil
}
