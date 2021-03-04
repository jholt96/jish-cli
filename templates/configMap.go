package templates

type simple_configEnv struct {
	ConfigMapKeyRef struct {
		Name string `yaml:"name"`
		Key  string `yaml:"key"`
	} `yaml:"configMapKeyRef"`
}

type simple_configMount struct {
	Name      string `yaml:"name"`
	ConfigMap struct {
		Name string `yaml:"name"`
	} `yaml:"configMap"`
}
