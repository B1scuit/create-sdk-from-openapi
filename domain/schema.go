package domain

type Schema struct {
	Type   string `yaml:"type,omitempty"`
	Format string `yaml:"format,omitempty"`
}

func (s *Schema) ReflectType() string {

	switch s.Type {
	case "datetime":
		return "time.Time"
	case "integer":
		return s.IntegerType()
	}

	return s.Type
}

func (s *Schema) IntegerType() string {
	return s.Format
}
