package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/B1scuit/create-sdk-from-openapi/domain"
	"github.com/Masterminds/sprig/v3"
)

func generateClient(openApi *domain.OpenApi) error {

	var functionList []string

	for _, methods := range *openApi.Paths {
		for _, pathSpec := range methods {
			functionList = append(functionList, pathSpec.OperationId)
		}

	}
	os.Mkdir(pkgName, os.ModePerm)
	f, err := os.Create(fmt.Sprintf("%s/client.auto.go", pkgName))
	if err != nil {
		return err
	}

	defer f.Close()

	tmpl := template.Must(template.New("client.tmpl").Funcs(template.FuncMap(sprig.FuncMap())).Funcs(funcMap).ParseFiles("templates/client.tmpl"))
	return tmpl.Execute(f, clientTemplate{
		Name:    "Client",
		Package: pkgName,
		Actions: functionList,
	})
}

type clientTemplate struct {
	Name    string
	Package string
	Server  string
	Actions []string
}
