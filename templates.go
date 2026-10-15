package mkprj

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"text/template"
	"time"
)

func getCurrentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

var templateRoot = filepath.Join(getCurrentDir(), "../templates")

func formatDate(t time.Time) string {
	return t.Format("02.01.2006")
}

func formatMoney(amount float64, currency string) string {
	return fmt.Sprintf("%s %s", localizeNumber(amount), currency)
}

func localizeNumber(number float64) string {
	re := regexp.MustCompile(`(\d)\.(\d)`)
	return re.ReplaceAllString(fmt.Sprintf("%.2f", number), "$1,$2")
}

func parseProse(s string) string {
	re := regexp.MustCompile(`([&])`)
	s = re.ReplaceAllString(s, "\\$1")
	re = regexp.MustCompile(`(?mUs:\*(\w.*\w)\*)`)
	return re.ReplaceAllString(s, "\\textit{$1}")
}

var funcMap = template.FuncMap{
	"formatdate":     formatDate,
	"formatmoney":    formatMoney,
	"localizenumber": localizeNumber,
	"parseprose":     parseProse,
}

func parseTemplates() *template.Template {
	templ := template.New("")
	err := filepath.Walk(templateRoot, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".tmpl") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}
		return err
	})
	if err != nil {
		panic(err)
	}

	return templ
}

var templates = parseTemplates()

func init() {
	funcMap := template.FuncMap{
		"formatdate": func(t time.Time) string {
			return t.Format("02.01.2006")
		},
		"formatmoney": func(amount float64, currency string) string {
			return fmt.Sprintf("%.2f %s", amount, currency)
		},
	}
	templates = templates.Funcs(funcMap).Delims("<%=", "%>")
	//     log.Printf("%+v", templates.DefinedTemplates())
}
