package templates

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

const secretAPIVersion, configMapAPIVersion, namespace string = "v1", "v1", "<namespace>"
const configMapKind, secretKind, secretType string = "ConfigMap", "Secret", "Opaque"
const username, password, exampleData string = "user", "Pass", "world"

func CreateConfigMap(name string) {

	fmt.Printf("Creating ConfigMap...\n\n")

	configMap := &ConfigMap{}

	configMap.ApiVersion = configMapAPIVersion
	configMap.Kind = configMapKind
	configMap.Metadata.Name = name
	configMap.Metadata.Namespace = namespace

	newConfigMapYaml, err := yaml.Marshal(configMap)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print(string(newConfigMapYaml))
	fmt.Print("\n\n")
}

func CreateSecret(name string) {

	fmt.Printf("Creating Secret...\n\n")

	secret := &Secret{}

	secret.ApiVersion = secretAPIVersion
	secret.Kind = secretKind
	secret.Metadata.Name = name
	secret.Metadata.Namespace = namespace
	secret.StringData.Username = username
	secret.StringData.Password = password
	secret.Type = secretType

	newSecretYaml, err := yaml.Marshal(secret)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print(string(newSecretYaml))
	fmt.Print("\n\n")
}
