package templates

type simpleSpec struct {
	Containers []struct {
		Name  string `yaml:"name"`
		Image string `yaml:"image"`
		Ports []struct {
			ContainerPort int `yaml:"containerPort"`
		} `yaml:"ports"`
	} `yaml:"containers"`
}

type simpleSpecEnv struct {
	Containers []struct {
		Name  string `yaml:"name"`
		Image string `yaml:"image"`
		Ports []struct {
			ContainerPort int `yaml:"containerPort"`
		} `yaml:"ports"`
		Env []struct {
			Name      string      `yaml:"name"`
			ValueFrom interface{} `yaml:"valueFrom"`
		} `yaml:"env"`
	} `yaml:"containers"`
}

type simpleSpecMount struct {
	Containers []struct {
		Name  string `yaml:"name"`
		Image string `yaml:"image"`
		Ports []struct {
			ContainerPort int `yaml:"containerPort"`
		} `yaml:"ports"`
		VolumeMounts []struct {
			Name      string `yaml:"name"`
			MountPath string `yaml:"mountPath"`
		} `yaml:"volumeMounts"`
	} `yaml:"containers"`
	Volumes []interface{} `yaml:"volumes"`
}

type simpleSpecEnvmount struct {
	Containers []struct {
		Name  string `yaml:"name"`
		Image string `yaml:"image"`
		Ports []struct {
			ContainerPort int `yaml:"containerPort"`
		} `yaml:"ports"`
		VolumeMounts []struct {
			Name      string `yaml:"name"`
			MountPath string `yaml:"mountPath"`
		} `yaml:"volumeMounts"`
		Env []struct {
			Name      string      `yaml:"name"`
			ValueFrom interface{} `yaml:"valueFrom"`
		} `yaml:"env"`
	} `yaml:"containers"`
	Volumes []interface{} `yaml:"volumes"`
}

type simpleDeployment struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
		Labels    struct {
			App string `yaml:"app"`
		} `yaml:"labels"`
	} `yaml:"metadata"`
	Spec struct {
		Replicas int `yaml:"replicas"`
		Selector struct {
			MatchLabels struct {
				App string `yaml:"app"`
			} `yaml:"matchLabels"`
		} `yaml:"selector"`
		Template struct {
			Metadata struct {
				Labels struct {
					App string `yaml:"app"`
				} `yaml:"labels"`
			} `yaml:"metadata"`
			Spec interface{} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
}
