package cmd

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"text/template"

	"github.com/B1scuit/create-sdk-from-openapi/domain"
	"github.com/Masterminds/sprig/v3"
)

var funcMap template.FuncMap

func generateComponents(components map[string]*domain.StructToBe) error {
	var wg sync.WaitGroup
	for name, structTobe := range components {
		wg.Add(1)
		go generateComponent(&wg, name, structTobe)
	}

	wg.Wait()

	return nil
}

func generateComponent(wg *sync.WaitGroup, name string, in *domain.StructToBe) error {
	defer wg.Done()

	switch in.Type {
	case "object":
		if err := generateStruct(name, in); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}

	return nil
}

func generateStruct(name string, in *domain.StructToBe) error {
	var base structTemplate

	base.Name = name
	base.Package = pkgName

	for fieldName, v := range in.Properties {

		baseType := v.ReflectType()

		base.Fields = append(base.Fields, &Field{
			Name: fieldName,
			Type: baseType,
		})
	}

	os.Mkdir(pkgName, os.ModePerm)
	f, err := os.Create(fmt.Sprintf("%s/%s.auto.go", pkgName, strings.ToLower(name)))
	if err != nil {
		return err
	}

	defer f.Close()

	var tmpl = template.Must(template.New("struct.tmpl").Funcs(template.FuncMap(sprig.FuncMap())).Funcs(funcMap).ParseFiles("templates/struct.tmpl"))
	if err := tmpl.Execute(f, base); err != nil {
		return err
	}

	return nil
}

type structTemplate struct {
	Name    string
	Package string
	Imports []string
	Fields  []*Field
}

type Field struct {
	Name string
	Type string
}
