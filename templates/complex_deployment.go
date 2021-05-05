package templates

type complexSpec struct {
	Containers []struct {
		Name            string `yaml:"name"`
		Image           string `yaml:"image"`
		ImagePullPolicy string `yaml:"imagePullPolicy"`
		Ports           []struct {
			ContainerPort int `yaml:"containerPort"`
		} `yaml:"ports"`
		Resources struct {
			Requests struct {
				CPU    string `yaml:"cpu"`
				Memory string `yaml:"memory"`
			} `yaml:"requests"`
			Limits struct {
				CPU    string `yaml:"cpu"`
				Memory string `yaml:"memory"`
			} `yaml:"limits"`
		} `yaml:"resources"`
		Env []struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
		} `yaml:"env"`
		VolumeMounts []struct {
			Name      string `yaml:"name"`
			MountPath string `yaml:"mountPath"`
		} `yaml:"volumeMounts"`
	} `yaml:"containers"`
	Volumes []emtpyDir `yaml:"volumes"`
}

type complexSpecEnv struct {
	Containers []struct {
		Name            string `yaml:"name"`
		Image           string `yaml:"image"`
		ImagePullPolicy string `yaml:"imagePullPolicy"`
		Ports           []struct {
			ContainerPort int `yaml:"containerPort"`
		} `yaml:"ports"`
		Resources struct {
			Requests struct {
				CPU    string `yaml:"cpu"`
				Memory string `yaml:"memory"`
			} `yaml:"requests"`
			Limits struct {
				CPU    string `yaml:"cpu"`
				Memory string `yaml:"memory"`
			} `yaml:"limits"`
		} `yaml:"resources"`
		Env []struct {
			Name      string      `yaml:"name"`
			ValueFrom interface{} `yaml:"valueFrom"`
		} `yaml:"env"`
		VolumeMounts []struct {
			Name      string `yaml:"name"`
			MountPath string `yaml:"mountPath"`
		} `yaml:"volumeMounts"`
	} `yaml:"containers"`
	Volumes []emtpyDir `yaml:"volumes"`
}

type complexSpecMount struct {
	Containers []struct {
		Name            string `yaml:"name"`
		Image           string `yaml:"image"`
		ImagePullPolicy string `yaml:"imagePullPolicy"`
		Ports           []struct {
			ContainerPort int `yaml:"containerPort"`
		} `yaml:"ports"`
		Resources struct {
			Requests struct {
				CPU    string `yaml:"cpu"`
				Memory string `yaml:"memory"`
			} `yaml:"requests"`
			Limits struct {
				CPU    string `yaml:"cpu"`
				Memory string `yaml:"memory"`
			} `yaml:"limits"`
		} `yaml:"resources"`
		Env []struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
		} `yaml:"env"`
		VolumeMounts []struct {
			Name      string `yaml:"name"`
			MountPath string `yaml:"mountPath"`
		} `yaml:"volumeMounts"`
	} `yaml:"containers"`
	Volumes []interface{} `yaml:"volumes"`
}

type complexSpecEnvmount struct {
	Containers []struct {
		Name            string `yaml:"name"`
		Image           string `yaml:"image"`
		ImagePullPolicy string `yaml:"imagePullPolicy"`
		Ports           []struct {
			ContainerPort int `yaml:"containerPort"`
		} `yaml:"ports"`
		Resources struct {
			Requests struct {
				CPU    string `yaml:"cpu"`
				Memory string `yaml:"memory"`
			} `yaml:"requests"`
			Limits struct {
				CPU    string `yaml:"cpu"`
				Memory string `yaml:"memory"`
			} `yaml:"limits"`
		} `yaml:"resources"`
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

type complexDeployment struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
		Labels    struct {
			App string `yaml:"app"`
		} `yaml:"labels"`
		Annotations struct {
			Key string `yaml:"key"`
		} `yaml:"annotations"`
	} `yaml:"metadata"`
	Spec struct {
		Replicas int `yaml:"replicas"`
		Strategy struct {
			Type          string `yaml:"type"`
			RollingUpdate struct {
				MaxSurge       string `yaml:"maxSurge"`
				MaxUnavailable string `yaml:"maxUnavailable"`
			} `yaml:"rollingUpdate"`
		} `yaml:"strategy"`
		Selector struct {
			MatchLabels struct {
				App string `yaml:"app"`
			} `yaml:"matchLabels"`
		} `yaml:"selector"`
		Template struct {
			ImagePullSecrets []struct {
				Name string `yaml:"name"`
			} `yaml:"imagePullSecrets"`
			Metadata struct {
				Labels struct {
					App string `yaml:"app"`
				} `yaml:"labels"`
			} `yaml:"metadata"`
			Spec interface{} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
}
