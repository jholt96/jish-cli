package cmd

import "testing"

func TestSimpleDeployment(t *testing.T) {
	createDeploymentYaml("env", "", "", "simple", "jish")
	createDeploymentYaml("mount", "", "", "simple", "jish")

	/*
		createDeploymentYaml("", "env", "", "simple", "jish")
		createDeploymentYaml("", "mount", "", "simple", "jish")
			createDeploymentYaml("env", "env", "", "simple", "jish")
			createDeploymentYaml("mount", "mount", "", "simple", "jish")
			createDeploymentYaml("env", "mount", "", "simple", "jish")
			createDeploymentYaml("mount", "env", "", "simple", "jish")
	*/

}
