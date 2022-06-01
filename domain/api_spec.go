package domain

type OpenApi struct {
	Version    string              `yaml:"openapi,omitempty"`
	Info       *OpenApi_Info       `yaml:"info,omitempty"`
	Servers    *OpenApi_Servers    `yaml:"servers,omitempty"`
	Components *OpenApi_Components `yaml:"components,omitempty"`
	Paths      *OpenApi_Paths      `yaml:"paths,omitempty"`
}

type OpenApi_Info struct {
	Title       string `yaml:"title,omitempty"`
	Description string `yaml:"description,omitempty"`
	Version     string `yaml:"version,omitempty"`
	GoPackage   string `yaml:"x-go-package,omitempty"`
}

type OpenApi_Servers []*Server

type Server struct {
	Url         *ServerUrl `yaml:"url,omitempty"`
	Description string     `yaml:"description,omitempty"`
}

// Format example map["/users"]["get"]Path{}
type OpenApi_Paths map[string]map[string]*Path

type OpenApi_Components struct {
	Schemas map[string]*StructToBe `yaml:"schemas,omitempty"`
}

type StructToBe struct {
	Type       string             `yaml:"type,omitempty"`
	Required   []string           `yaml:"required,omitempty"`
	Properties map[string]*Schema `yaml:"properties,omitempty"`
}
