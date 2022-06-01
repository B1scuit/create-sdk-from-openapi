package cmd

import (
	"bytes"
	"errors"
	"os"

	"github.com/B1scuit/create-sdk-from-openapi/domain"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	errFileMissing = errors.New("input file is missing")

	pkgName string
)

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errFileMissing
		}

		openApi, err := openYAMLSpec(args[0])
		if err != nil {
			return err
		}

		pkgName = openApi.Info.GoPackage

		if openApi.Components != nil {
			generateComponents(openApi.Components.Schemas)
		}

		if err := generateActions(openApi.Paths); err != nil {
			return err
		}

		if err := generateClient(openApi); err != nil {
			return err
		}

		return nil
	},
}

func openYAMLSpec(fileName string) (*domain.OpenApi, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	b := bytes.Buffer{}
	b.ReadFrom(f)

	var openapiSpec domain.OpenApi
	if err := yaml.Unmarshal(b.Bytes(), &openapiSpec); err != nil {
		return nil, err
	}

	return &openapiSpec, nil
}

func Run() error {
	return rootCmd.Execute()
}
