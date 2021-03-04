package templates

const apiVersion, kind, image string = "apps/v1", "Deployment", "<image:tag>"
const replicas, containerPort int = 3, 80

type SimpleSpec struct {
	Containers []struct {
		Name  string `yaml:"name"`
		Image string `yaml:"image"`
		Ports []struct {
			ContainerPort int `yaml:"containerPort"`
		} `yaml:"ports"`
	} `yaml:"containers"`
}

type SimpleSpecEnv struct {
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

type SimpleSpecMount struct {
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

type SimpleSpecEnvmount struct {
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

type SimpleDeployment struct {
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

func createSimpleSpecBase(name string) SimpleSpec {
	simpleSpec := SimpleSpec{}

	simpleSpec.Containers = make([]struct {
		Name  string "yaml:\"name\""
		Image string "yaml:\"image\""
		Ports []struct {
			ContainerPort int "yaml:\"containerPort\""
		} "yaml:\"ports\""
	}, 1)
	simpleSpec.Containers[0].Ports = make([]struct {
		ContainerPort int "yaml:\"containerPort\""
	}, 1)

	simpleSpec.Containers[0].Name = name + "-container"
	simpleSpec.Containers[0].Image = image

	simpleSpec.Containers[0].Ports[0].ContainerPort = containerPort

	return simpleSpec
}

func CreateSimpleDeploymentYaml(name string, configMap string, secret string, service string) *SimpleDeployment {

	/*
		configMount := false
		configEnv := false
		switch configMap {
		case "":
			configMount = false
			configEnv = false
		case "env":
			configEnv = true
		case "mount":
			configMount = true
		default:
			configMount = false
			configEnv = false
		}

		if !configEnv && !configMount {
			//if all are false then use simple
		}*/

	simpleDeployment := &SimpleDeployment{}

	simpleDeployment.ApiVersion = apiVersion
	simpleDeployment.Kind = kind
	simpleDeployment.Spec.Replicas = replicas
	simpleDeployment.Metadata.Labels.App = name
	simpleDeployment.Metadata.Name = name
	simpleDeployment.Metadata.Namespace = "<namespace>"
	simpleDeployment.Spec.Selector.MatchLabels.App = name
	simpleDeployment.Spec.Template.Metadata.Labels.App = name
	simpleDeployment.Spec.Template.Spec = createSimpleSpecBase(name)

	return simpleDeployment
}
