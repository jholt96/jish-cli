package templates

type secretEnvValueFrom struct {
	SecretKeyRef struct {
		Name string `yaml:"name"`
		Key  string `yaml:"key"`
	} `yaml:"secretKeyRef"`
}

type secretMountVolume struct {
	Name   string `yaml:"name"`
	Secret struct {
		SecretName string `yaml:"secretName"`
	} `yaml:"secret"`
}

type secret struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Type       string `yaml:"type"`
	StringData struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"stringData"`
}
