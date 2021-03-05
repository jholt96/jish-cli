package templates

type ConfigEnvValueFrom struct {
	ConfigMapKeyRef struct {
		Name string `yaml:"name"`
		Key  string `yaml:"key"`
	} `yaml:"configMapKeyRef"`
}

type ConfigMount struct {
	Name      string `yaml:"name"`
	ConfigMap struct {
		Name string `yaml:"name"`
	} `yaml:"configMap"`
}

type ConfigMap struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Data struct {
	} `yaml:"data"`
}
