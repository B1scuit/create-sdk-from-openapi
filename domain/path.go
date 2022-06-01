package domain

type Path struct {

	// Some internal methods for cocurrency saftey
	requestPath   string
	requestMethod string

	OperationId string              `yaml:"operationId,omitempty"` // What we will call the resuling function
	Summary     string              `yaml:"summary,omitempty"`
	Description string              `yaml:"description,omitempty"`
	Deprecated  bool                `yaml:"deprecated,omitempty"`
	Responses   map[string]Response `yaml:"responses,omitempty"`
	Parameters  Parameters          `yaml:"parameters,omitempty"`
}

type Response struct{}

type Parameters []Parameter

type Parameter struct {
	Name        string `yaml:"name,omitempty"`
	In          string `yaml:"in,omitempty"`
	Description string `yaml:"description,omitempty"`
	Required    bool   `yaml:"required,omitempty"`
	Schema      Schema `yaml:"schema,omitempty"`
}
