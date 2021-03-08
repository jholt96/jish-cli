package templates

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func createStandardSpecBase(name string) standardSpec {
	standardSpec := standardSpec{}

	standardSpec.Containers = make([]struct {
		Name            string "yaml:\"name\""
		Image           string "yaml:\"image\""
		ImagePullPolicy string "yaml:\"imagePullPolicy\""
		Ports           []struct {
			ContainerPort int "yaml:\"containerPort\""
		} "yaml:\"ports\""
		Resources struct {
			Requests struct {
				CPU    string "yaml:\"cpu\""
				Memory string "yaml:\"memory\""
			} "yaml:\"requests\""
			Limits struct {
				CPU    string "yaml:\"cpu\""
				Memory string "yaml:\"memory\""
			} "yaml:\"limits\""
		} "yaml:\"resources\""
		Env []struct {
			Name  string "yaml:\"name\""
			Value string "yaml:\"value\""
		} "yaml:\"env\""
		VolumeMounts []struct {
			Name      string "yaml:\"name\""
			MountPath string "yaml:\"mountPath\""
		} "yaml:\"volumeMounts\""
	}, 1)

	standardSpec.Containers[0].Ports = make([]struct {
		ContainerPort int "yaml:\"containerPort\""
	}, 1)
	standardSpec.Containers[0].Env = make([]struct {
		Name  string "yaml:\"name\""
		Value string "yaml:\"value\""
	}, 1)
	standardSpec.Containers[0].VolumeMounts = make([]struct {
		Name      string "yaml:\"name\""
		MountPath string "yaml:\"mountPath\""
	}, 1)
	standardSpec.Volumes = make([]emtpyDir, 1)
	standardSpec.Containers[0].VolumeMounts = make([]struct {
		Name      string "yaml:\"name\""
		MountPath string "yaml:\"mountPath\""
	}, 1)

	standardSpec.Containers[0].Name = name + "-container"
	standardSpec.Containers[0].Image = image
	standardSpec.Containers[0].Ports[0].ContainerPort = containerPort

	standardSpec.Containers[0].ImagePullPolicy = imagePullPolicy
	standardSpec.Containers[0].Resources.Requests.CPU = resourceRequestCPU
	standardSpec.Containers[0].Resources.Requests.Memory = resourceRequestMemory
	standardSpec.Containers[0].Resources.Limits.CPU = resourceLimitCPU
	standardSpec.Containers[0].Resources.Limits.Memory = resourceLimitMemory

	standardSpec.Containers[0].Env[0].Name = usernameKey
	standardSpec.Containers[0].Env[0].Value = username

	standardSpec.Containers[0].VolumeMounts[0].MountPath = mountPath
	standardSpec.Containers[0].VolumeMounts[0].Name = name

	standardSpec.Volumes[0].Name = name

	return standardSpec
}

func createStandardSpecEnv(name string, envName []string, numOfEnv int, valueFrom ...interface{}) *standardSpecEnv {
	standardSpecEnv := &standardSpecEnv{}

	standardSpecEnv.Containers = make([]struct {
		Name            string "yaml:\"name\""
		Image           string "yaml:\"image\""
		ImagePullPolicy string "yaml:\"imagePullPolicy\""
		Ports           []struct {
			ContainerPort int "yaml:\"containerPort\""
		} "yaml:\"ports\""
		Resources struct {
			Requests struct {
				CPU    string "yaml:\"cpu\""
				Memory string "yaml:\"memory\""
			} "yaml:\"requests\""
			Limits struct {
				CPU    string "yaml:\"cpu\""
				Memory string "yaml:\"memory\""
			} "yaml:\"limits\""
		} "yaml:\"resources\""
		Env []struct {
			Name      string      "yaml:\"name\""
			ValueFrom interface{} "yaml:\"valueFrom\""
		} "yaml:\"env\""
		VolumeMounts []struct {
			Name      string "yaml:\"name\""
			MountPath string "yaml:\"mountPath\""
		} "yaml:\"volumeMounts\""
	}, 1)

	standardSpecEnv.Containers[0].Ports = make([]struct {
		ContainerPort int "yaml:\"containerPort\""
	}, 1)

	standardSpecEnv.Containers[0].Image = image
	standardSpecEnv.Containers[0].Name = name + "-container"
	standardSpecEnv.Containers[0].Ports[0].ContainerPort = containerPort

	standardSpecEnv.Containers[0].Env = make([]struct {
		Name      string      "yaml:\"name\""
		ValueFrom interface{} "yaml:\"valueFrom\""
	}, numOfEnv)

	standardSpecEnv.Containers[0].VolumeMounts = make([]struct {
		Name      string "yaml:\"name\""
		MountPath string "yaml:\"mountPath\""
	}, 1)

	standardSpecEnv.Containers[0].ImagePullPolicy = imagePullPolicy
	standardSpecEnv.Containers[0].Resources.Requests.CPU = resourceRequestCPU
	standardSpecEnv.Containers[0].Resources.Requests.Memory = resourceRequestMemory
	standardSpecEnv.Containers[0].Resources.Limits.CPU = resourceLimitCPU
	standardSpecEnv.Containers[0].Resources.Limits.Memory = resourceLimitMemory

	standardSpecEnv.Containers[0].VolumeMounts[0].MountPath = mountPath
	standardSpecEnv.Containers[0].VolumeMounts[0].Name = name

	standardSpecEnv.Volumes = make([]emtpyDir, 1)
	standardSpecEnv.Volumes[0].Name = name

	for i, env := range valueFrom {
		standardSpecEnv.Containers[0].Env[i].Name = envName[i]

		standardSpecEnv.Containers[0].Env[i].ValueFrom = env
	}

	return standardSpecEnv

}

func createStandardSpecMount(name string, envName []string, numOfVolumes int, volumeMount ...interface{}) *standardSpecMount {

	standardSpecMount := &standardSpecMount{}

	standardSpecMount.Containers = make([]struct {
		Name            string "yaml:\"name\""
		Image           string "yaml:\"image\""
		ImagePullPolicy string "yaml:\"imagePullPolicy\""
		Ports           []struct {
			ContainerPort int "yaml:\"containerPort\""
		} "yaml:\"ports\""
		Resources struct {
			Requests struct {
				CPU    string "yaml:\"cpu\""
				Memory string "yaml:\"memory\""
			} "yaml:\"requests\""
			Limits struct {
				CPU    string "yaml:\"cpu\""
				Memory string "yaml:\"memory\""
			} "yaml:\"limits\""
		} "yaml:\"resources\""
		Env []struct {
			Name  string "yaml:\"name\""
			Value string "yaml:\"value\""
		} "yaml:\"env\""
		VolumeMounts []struct {
			Name      string "yaml:\"name\""
			MountPath string "yaml:\"mountPath\""
		} "yaml:\"volumeMounts\""
	}, 1)

	standardSpecMount.Containers[0].Ports = make([]struct {
		ContainerPort int "yaml:\"containerPort\""
	}, 1)

	standardSpecMount.Containers[0].Env = make([]struct {
		Name  string "yaml:\"name\""
		Value string "yaml:\"value\""
	}, 1)
	standardSpecMount.Containers[0].Name = name + "-container"
	standardSpecMount.Containers[0].Image = image
	standardSpecMount.Containers[0].Ports[0].ContainerPort = containerPort

	standardSpecMount.Containers[0].ImagePullPolicy = imagePullPolicy
	standardSpecMount.Containers[0].Resources.Requests.CPU = resourceRequestCPU
	standardSpecMount.Containers[0].Resources.Requests.Memory = resourceRequestMemory
	standardSpecMount.Containers[0].Resources.Limits.CPU = resourceLimitCPU
	standardSpecMount.Containers[0].Resources.Limits.Memory = resourceLimitMemory

	standardSpecMount.Containers[0].Env[0].Name = usernameKey
	standardSpecMount.Containers[0].Env[0].Value = username

	standardSpecMount.Containers[0].VolumeMounts = make([]struct {
		Name      string "yaml:\"name\""
		MountPath string "yaml:\"mountPath\""
	}, numOfVolumes)

	for i := range standardSpecMount.Containers[0].VolumeMounts {

		standardSpecMount.Containers[0].VolumeMounts[i].MountPath = mountPath
		standardSpecMount.Containers[0].VolumeMounts[i].Name = envName[i]
	}

	standardSpecMount.Volumes = make([]interface{}, numOfVolumes)

	for i, volume := range volumeMount {

		standardSpecMount.Volumes[i] = volume
	}

	return standardSpecMount
}

func createStandardEnvandMount(name string, envMountEnvName []string, volumeMount interface{}, envValueFrom interface{}) *standardSpecEnvmount {
	standardSpecEnvmount := &standardSpecEnvmount{}

	standardSpecEnvmount.Containers = make([]struct {
		Name            string "yaml:\"name\""
		Image           string "yaml:\"image\""
		ImagePullPolicy string "yaml:\"imagePullPolicy\""
		Ports           []struct {
			ContainerPort int "yaml:\"containerPort\""
		} "yaml:\"ports\""
		Resources struct {
			Requests struct {
				CPU    string "yaml:\"cpu\""
				Memory string "yaml:\"memory\""
			} "yaml:\"requests\""
			Limits struct {
				CPU    string "yaml:\"cpu\""
				Memory string "yaml:\"memory\""
			} "yaml:\"limits\""
		} "yaml:\"resources\""
		VolumeMounts []struct {
			Name      string "yaml:\"name\""
			MountPath string "yaml:\"mountPath\""
		} "yaml:\"volumeMounts\""
		Env []struct {
			Name      string      "yaml:\"name\""
			ValueFrom interface{} "yaml:\"valueFrom\""
		} "yaml:\"env\""
	}, 1)

	standardSpecEnvmount.Containers[0].Ports = make([]struct {
		ContainerPort int "yaml:\"containerPort\""
	}, 1)

	standardSpecEnvmount.Containers[0].Name = name + "-container"
	standardSpecEnvmount.Containers[0].Image = image
	standardSpecEnvmount.Containers[0].Ports[0].ContainerPort = containerPort

	standardSpecEnvmount.Containers[0].ImagePullPolicy = imagePullPolicy
	standardSpecEnvmount.Containers[0].Resources.Requests.CPU = resourceRequestCPU
	standardSpecEnvmount.Containers[0].Resources.Requests.Memory = resourceRequestMemory
	standardSpecEnvmount.Containers[0].Resources.Limits.CPU = resourceLimitCPU
	standardSpecEnvmount.Containers[0].Resources.Limits.Memory = resourceLimitMemory

	standardSpecEnvmount.Containers[0].Env = make([]struct {
		Name      string      "yaml:\"name\""
		ValueFrom interface{} "yaml:\"valueFrom\""
	}, 1)

	standardSpecEnvmount.Containers[0].Env[0].Name = envMountEnvName[0]
	standardSpecEnvmount.Containers[0].Env[0].ValueFrom = envValueFrom

	standardSpecEnvmount.Containers[0].VolumeMounts = make([]struct {
		Name      string "yaml:\"name\""
		MountPath string "yaml:\"mountPath\""
	}, 1)

	standardSpecEnvmount.Containers[0].VolumeMounts[0].MountPath = mountPath
	standardSpecEnvmount.Containers[0].VolumeMounts[0].Name = envMountEnvName[1]

	standardSpecEnvmount.Volumes = make([]interface{}, 1)
	standardSpecEnvmount.Volumes[0] = volumeMount

	return standardSpecEnvmount
}

func CreateStandardDeploymentYaml(name, configMap, secret, service string, ch chan<- string) {

	fmt.Printf("Creating standard Deployment...\n\n")

	configMount, configEnv := false, false
	secretMount, secretEnv := false, false

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

	switch secret {
	case "":
		secretMount = false
		secretEnv = false
	case "env":
		secretEnv = true
	case "mount":
		secretMount = true
	default:
		secretMount = false
		secretEnv = false
	}

	standardDeployment := &standardDeployment{}

	standardDeployment.ApiVersion = deploymentAPIVersion
	standardDeployment.Kind = deploymentKind
	standardDeployment.Metadata.Labels.App = name
	standardDeployment.Metadata.Name = name
	standardDeployment.Metadata.Namespace = namespace
	standardDeployment.Metadata.Annotations.Key = deploymentKind
	standardDeployment.Spec.Replicas = replicas
	standardDeployment.Spec.Selector.MatchLabels.App = name
	standardDeployment.Spec.Template.Metadata.Labels.App = name
	standardDeployment.Spec.Strategy.RollingUpdate.MaxSurge = "25%"
	standardDeployment.Spec.Strategy.RollingUpdate.MaxUnavailable = "25%"
	standardDeployment.Spec.Strategy.Type = strategyType

	if !configEnv && !configMount && !secretEnv && !secretMount { //if no flags secret or configmap
		standardDeployment.Spec.Template.Spec = createStandardSpecBase(name)

	} else if configEnv && !configMount && !secretEnv && !secretMount { //if only configmap with env is set
		envName := []string{name + "-configmap"}
		valueFrom := &configEnvValueFrom{}
		valueFrom.ConfigMapKeyRef.Key = envKey
		valueFrom.ConfigMapKeyRef.Name = envName[0]

		standardDeployment.Spec.Template.Spec = *(createStandardSpecEnv(name, envName, 1, valueFrom))

	} else if !configEnv && configMount && !secretEnv && !secretMount { //if only configmap with volume mount is set
		envName := []string{name + "-configmap"}
		configVolume := configMountVolume{}

		configVolume.Name = envName[0]
		configVolume.ConfigMap.Name = envName[0]

		standardDeployment.Spec.Template.Spec = *(createStandardSpecMount(name, envName, 1, configVolume))

	} else if !configEnv && !configMount && secretEnv && !secretMount { //if only secret with env is set
		envName := []string{name + "-secret"}
		valueFrom := &secretEnvValueFrom{}
		valueFrom.SecretKeyRef.Name = envName[0]
		valueFrom.SecretKeyRef.Key = envName[0] + "-key"

		standardDeployment.Spec.Template.Spec = *createStandardSpecEnv(name, envName, 1, valueFrom)

	} else if !configEnv && !configMount && !secretEnv && secretMount { //if only secret with volume mount is set
		envName := []string{name + "-secret"}
		SecretVolume := secretMountVolume{}

		SecretVolume.Name = envName[0]
		SecretVolume.Secret.SecretName = envName[0]

		standardDeployment.Spec.Template.Spec = *createStandardSpecMount(name, envName, 1, SecretVolume)

	} else if configEnv && !configMount && secretEnv && !secretMount { // if configmap with env and secret with env

		envName := []string{name + "-secret", name + "-configmap"}

		valueFromSecret := &secretEnvValueFrom{}
		valueFromSecret.SecretKeyRef.Name = envName[0]
		valueFromSecret.SecretKeyRef.Key = envName[0] + "-key"

		valueFromConfig := &configEnvValueFrom{}
		valueFromConfig.ConfigMapKeyRef.Key = envKey
		valueFromConfig.ConfigMapKeyRef.Name = envName[1]

		standardDeployment.Spec.Template.Spec = *(createStandardSpecEnv(name, envName, 2, valueFromSecret, valueFromConfig))

	} else if configEnv && !configMount && !secretEnv && secretMount { //if configmap with env and secret with mount
		envName := []string{name + "-configmap", name + "-secret"}

		valueFromConfig := &configEnvValueFrom{}
		valueFromConfig.ConfigMapKeyRef.Key = envKey
		valueFromConfig.ConfigMapKeyRef.Name = envName[1]

		SecretVolume := secretMountVolume{}
		SecretVolume.Name = envName[0]
		SecretVolume.Secret.SecretName = envName[0]

		standardDeployment.Spec.Template.Spec = *(createStandardEnvandMount(name, envName, SecretVolume, valueFromConfig))

	} else if !configEnv && configMount && secretEnv && !secretMount { //if configmap with mount and secret with env
		envName := []string{name + "-secret", name + "-configmap"}

		valueFromSecret := &secretEnvValueFrom{}
		valueFromSecret.SecretKeyRef.Name = envName[0]
		valueFromSecret.SecretKeyRef.Key = envName[0] + "-key"

		configVolume := configMountVolume{}
		configVolume.Name = envName[1]
		configVolume.ConfigMap.Name = envName[1]

		standardDeployment.Spec.Template.Spec = *(createStandardEnvandMount(name, envName, configVolume, valueFromSecret))

	} else if !configEnv && configMount && !secretEnv && secretMount { //if configmap with mount and secret with mount
		envName := []string{name + "-secret", name + "-configmap"}

		SecretVolume := secretMountVolume{}
		SecretVolume.Name = envName[0]
		SecretVolume.Secret.SecretName = envName[0]

		configVolume := configMountVolume{}
		configVolume.Name = envName[1]
		configVolume.ConfigMap.Name = envName[1]

		standardDeployment.Spec.Template.Spec = createStandardSpecMount(name, envName, 2, SecretVolume, configVolume)

	} else { //default to just base
		standardDeployment.Spec.Template.Spec = createStandardSpecBase(name)
	}

	newDeploymentYaml, err := yaml.Marshal(standardDeployment)

	if err != nil {
		fmt.Printf(err.Error())
	}

	writeFileErr := ioutil.WriteFile((name + "-deployment.yaml"), newDeploymentYaml, 0755)

	if writeFileErr != nil {
		ch <- writeFileErr.Error()
	} else {
		ch <- (name + "-deployment.yaml has been created in the current directory\n")
	}
}
