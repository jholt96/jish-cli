/*
All structs must be exported
*/

package templates

type configEnvValueFrom struct {
	ConfigMapKeyRef struct {
		Name string `yaml:"name"`
		Key  string `yaml:"key"`
	} `yaml:"configMapKeyRef"`
}

type configMountVolume struct {
	Name      string `yaml:"name"`
	ConfigMap struct {
		Name string `yaml:"name"`
	} `yaml:"configMap"`
}

type configMap struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Data struct {
	} `yaml:"data"`
}
