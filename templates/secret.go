package templates

type simple_secretEnv struct {
	SecretKeyRef struct {
		Name string `yaml:"name"`
		Key  string `yaml:"key"`
	} `yaml:"secretKeyRef"`
}

type simple_secretMount struct {
	Name   string `yaml:"name"`
	Secret struct {
		SecretName string `yaml:"secretName"`
	} `yaml:"secret"`
}
