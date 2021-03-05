package templates

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const deploymentAPIVersion, deploymentKind, image, envKey string = "apps/v1", "Deployment", "<image:tag>", "<key>"
const replicas, containerPort int = 3, 80
const mountPath string = "/"

func createSimpleSpecBase(name string) simpleSpec {
	simpleSpec := simpleSpec{}

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

func createSimpleSpecEnv(name string, envName []string, numOfEnv int, valueFrom ...interface{}) *simpleSpecEnv {
	simpleSpecEnv := &simpleSpecEnv{}

	simpleSpecEnv.Containers = make([]struct {
		Name  string "yaml:\"name\""
		Image string "yaml:\"image\""
		Ports []struct {
			ContainerPort int "yaml:\"containerPort\""
		} "yaml:\"ports\""
		Env []struct {
			Name      string      "yaml:\"name\""
			ValueFrom interface{} "yaml:\"valueFrom\""
		} "yaml:\"env\""
	}, 1)

	simpleSpecEnv.Containers[0].Ports = make([]struct {
		ContainerPort int "yaml:\"containerPort\""
	}, 1)

	simpleSpecEnv.Containers[0].Image = image
	simpleSpecEnv.Containers[0].Name = name + "-container"
	simpleSpecEnv.Containers[0].Ports[0].ContainerPort = containerPort

	simpleSpecEnv.Containers[0].Env = make([]struct {
		Name      string      "yaml:\"name\""
		ValueFrom interface{} "yaml:\"valueFrom\""
	}, numOfEnv)

	for i, env := range valueFrom {
		simpleSpecEnv.Containers[0].Env[i].Name = envName[i]

		simpleSpecEnv.Containers[0].Env[i].ValueFrom = env
	}

	return simpleSpecEnv

}

func createSimpleSpecMount(name string, envName []string, numOfVolumes int, volumeMount ...interface{}) *simpleSpecMount {

	simpleSpecMount := &simpleSpecMount{}

	simpleSpecMount.Containers = make([]struct {
		Name  string "yaml:\"name\""
		Image string "yaml:\"image\""
		Ports []struct {
			ContainerPort int "yaml:\"containerPort\""
		} "yaml:\"ports\""
		VolumeMounts []struct {
			Name      string "yaml:\"name\""
			MountPath string "yaml:\"mountPath\""
		} "yaml:\"volumeMounts\""
	}, 1)
	simpleSpecMount.Containers[0].Ports = make([]struct {
		ContainerPort int "yaml:\"containerPort\""
	}, 1)

	simpleSpecMount.Containers[0].Image = image
	simpleSpecMount.Containers[0].Name = name + "-container"
	simpleSpecMount.Containers[0].Ports[0].ContainerPort = containerPort
	simpleSpecMount.Containers[0].VolumeMounts = make([]struct {
		Name      string "yaml:\"name\""
		MountPath string "yaml:\"mountPath\""
	}, numOfVolumes)

	for i := range simpleSpecMount.Containers[0].VolumeMounts {

		simpleSpecMount.Containers[0].VolumeMounts[i].MountPath = mountPath
		simpleSpecMount.Containers[0].VolumeMounts[i].Name = envName[i]
	}

	simpleSpecMount.Volumes = make([]interface{}, numOfVolumes)

	for i, volume := range volumeMount {

		simpleSpecMount.Volumes[i] = volume
	}

	return simpleSpecMount
}

func createSimpleEnvandMount(name string, envMountEnvName []string, volumeMount interface{}, envValueFrom interface{}) *simpleSpecEnvmount {
	simpleSpecEnvmount := &simpleSpecEnvmount{}

	simpleSpecEnvmount.Containers = make([]struct {
		Name  string "yaml:\"name\""
		Image string "yaml:\"image\""
		Ports []struct {
			ContainerPort int "yaml:\"containerPort\""
		} "yaml:\"ports\""
		VolumeMounts []struct {
			Name      string "yaml:\"name\""
			MountPath string "yaml:\"mountPath\""
		} "yaml:\"volumeMounts\""
		Env []struct {
			Name      string      "yaml:\"name\""
			ValueFrom interface{} "yaml:\"valueFrom\""
		} "yaml:\"env\""
	}, 1)

	simpleSpecEnvmount.Containers[0].Ports = make([]struct {
		ContainerPort int "yaml:\"containerPort\""
	}, 1)

	simpleSpecEnvmount.Containers[0].Image = image
	simpleSpecEnvmount.Containers[0].Name = name + "-container"
	simpleSpecEnvmount.Containers[0].Ports[0].ContainerPort = containerPort

	simpleSpecEnvmount.Containers[0].Env = make([]struct {
		Name      string      "yaml:\"name\""
		ValueFrom interface{} "yaml:\"valueFrom\""
	}, 1)

	simpleSpecEnvmount.Containers[0].Env[0].Name = envMountEnvName[0]
	simpleSpecEnvmount.Containers[0].Env[0].ValueFrom = envValueFrom

	simpleSpecEnvmount.Containers[0].VolumeMounts = make([]struct {
		Name      string "yaml:\"name\""
		MountPath string "yaml:\"mountPath\""
	}, 1)

	simpleSpecEnvmount.Containers[0].VolumeMounts[0].MountPath = mountPath
	simpleSpecEnvmount.Containers[0].VolumeMounts[0].Name = envMountEnvName[1]

	simpleSpecEnvmount.Volumes = make([]interface{}, 1)
	simpleSpecEnvmount.Volumes[0] = volumeMount

	return simpleSpecEnvmount
}

func CreateSimpleDeploymentYaml(name string, configMap string, secret string, service string) {

	fmt.Printf("Creating Simple Deployment...\n\n\n")

	configMount := false
	configEnv := false
	secretMount := false
	secretEnv := false

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

	/*
		switch service {
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
	*/

	simpleDeployment := &simpleDeployment{}

	simpleDeployment.ApiVersion = deploymentAPIVersion
	simpleDeployment.Kind = deploymentKind
	simpleDeployment.Spec.Replicas = replicas
	simpleDeployment.Metadata.Labels.App = name
	simpleDeployment.Metadata.Name = name
	simpleDeployment.Metadata.Namespace = "<namespace>"
	simpleDeployment.Spec.Selector.MatchLabels.App = name
	simpleDeployment.Spec.Template.Metadata.Labels.App = name

	if !configEnv && !configMount && !secretEnv && !secretMount { //if no flags secret or configmap
		simpleDeployment.Spec.Template.Spec = createSimpleSpecBase(name)

	} else if configEnv && !configMount && !secretEnv && !secretMount { //if only configmap with env is set
		envName := []string{name + "-configmap"}
		valueFrom := &configEnvValueFrom{}
		valueFrom.ConfigMapKeyRef.Key = envKey
		valueFrom.ConfigMapKeyRef.Name = envName[0]

		simpleDeployment.Spec.Template.Spec = *(createSimpleSpecEnv(name, envName, 1, valueFrom))

	} else if !configEnv && configMount && !secretEnv && !secretMount { //if only configmap with volume mount is set
		envName := []string{name + "-configmap"}
		configVolume := configMountVolume{}

		configVolume.Name = envName[0]
		configVolume.ConfigMap.Name = envName[0]

		simpleDeployment.Spec.Template.Spec = *(createSimpleSpecMount(name, envName, 1, configVolume))

	} else if !configEnv && !configMount && secretEnv && !secretMount { //if only secret with env is set
		envName := []string{name + "-secret"}
		valueFrom := &secretEnvValueFrom{}
		valueFrom.SecretKeyRef.Name = envName[0]
		valueFrom.SecretKeyRef.Key = envName[0] + "-key"

		simpleDeployment.Spec.Template.Spec = *createSimpleSpecEnv(name, envName, 1, valueFrom)

	} else if !configEnv && !configMount && !secretEnv && secretMount { //if only secret with volume mount is set
		envName := []string{name + "-secret"}
		SecretVolume := secretMountVolume{}

		SecretVolume.Name = envName[0]
		SecretVolume.Secret.SecretName = envName[0]

		simpleDeployment.Spec.Template.Spec = *createSimpleSpecMount(name, envName, 1, SecretVolume)

	} else if configEnv && !configMount && secretEnv && !secretMount { // if configmap with env and secret with env

		envName := []string{name + "-secret", name + "-configmap"}

		valueFromSecret := &secretEnvValueFrom{}
		valueFromSecret.SecretKeyRef.Name = envName[0]
		valueFromSecret.SecretKeyRef.Key = envName[0] + "-key"

		valueFromConfig := &configEnvValueFrom{}
		valueFromConfig.ConfigMapKeyRef.Key = envKey
		valueFromConfig.ConfigMapKeyRef.Name = envName[1]

		simpleDeployment.Spec.Template.Spec = *(createSimpleSpecEnv(name, envName, 2, valueFromSecret, valueFromConfig))

	} else if configEnv && !configMount && !secretEnv && secretMount { //if configmap with env and secret with mount
		envName := []string{name + "-configmap", name + "-secret"}

		valueFromConfig := &configEnvValueFrom{}
		valueFromConfig.ConfigMapKeyRef.Key = envKey
		valueFromConfig.ConfigMapKeyRef.Name = envName[1]

		SecretVolume := secretMountVolume{}
		SecretVolume.Name = envName[0]
		SecretVolume.Secret.SecretName = envName[0]

		simpleDeployment.Spec.Template.Spec = *(createSimpleEnvandMount(name, envName, SecretVolume, valueFromConfig))

	} else if !configEnv && configMount && secretEnv && !secretMount { //if configmap with mount and secret with env
		envName := []string{name + "-secret", name + "-configmap"}

		valueFromSecret := &secretEnvValueFrom{}
		valueFromSecret.SecretKeyRef.Name = envName[0]
		valueFromSecret.SecretKeyRef.Key = envName[0] + "-key"

		configVolume := configMountVolume{}
		configVolume.Name = envName[1]
		configVolume.ConfigMap.Name = envName[1]

		simpleDeployment.Spec.Template.Spec = *(createSimpleEnvandMount(name, envName, configVolume, valueFromSecret))

	} else if !configEnv && configMount && !secretEnv && secretMount { //if configmap with mount and secret with mount
		envName := []string{name + "-secret", name + "-configmap"}

		SecretVolume := secretMountVolume{}
		SecretVolume.Name = envName[0]
		SecretVolume.Secret.SecretName = envName[0]

		configVolume := configMountVolume{}
		configVolume.Name = envName[1]
		configVolume.ConfigMap.Name = envName[1]

		simpleDeployment.Spec.Template.Spec = createSimpleSpecMount(name, envName, 2, SecretVolume, configVolume)

	} else { //default to just base
		simpleDeployment.Spec.Template.Spec = createSimpleSpecBase(name)
	}

	newDeploymentYaml, err := yaml.Marshal(simpleDeployment)

	if err != nil {
		fmt.Printf(err.Error())
	}

	writeFileErr := ioutil.WriteFile((name + "-deployment.yaml"), newDeploymentYaml, 0755)

	if writeFileErr != nil {
		fmt.Printf(writeFileErr.Error())
	} else {
		fmt.Printf("%s-deployment.yaml has been created in the current directory\n", name)
	}
}
