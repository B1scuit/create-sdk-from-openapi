package domain

import (
	"net/url"

	"gopkg.in/yaml.v3"
)

type ServerUrl url.URL

func (su *ServerUrl) UnmarshalYAML(value *yaml.Node) error {

	parsedUrl, err := url.Parse(value.Value)
	if err != nil {
		return err
	}

	serverUrl := ServerUrl(*parsedUrl)

	su = &serverUrl

	return nil
}
