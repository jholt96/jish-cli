package templates

type service struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Spec struct {
		Selector struct {
			App string `yaml:"app"`
		} `yaml:"selector"`
		Ports []struct {
			Port       int `yaml:"port"`
			TargetPort int `yaml:"targetPort"`
		} `yaml:"ports"`
	} `yaml:"spec"`
}

type serviceWithNodePort struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		Type     string `yaml:"type"`
		Selector struct {
			App string `yaml:"app"`
		} `yaml:"selector"`
		Ports []struct {
			Port       int `yaml:"port"`
			TargetPort int `yaml:"targetPort"`
			NodePort   int `yaml:"nodePort"`
		} `yaml:"ports"`
	} `yaml:"spec"`
}
