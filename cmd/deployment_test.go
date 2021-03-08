package cmd

import "testing"

func TestSimpleDeployment(t *testing.T) {
	mode := "standard"
	createDeploymentYaml("env", "", "", mode, "test1")
	createDeploymentYaml("mount", "", "", mode, "test2")
	createDeploymentYaml("", "env", "", mode, "test3")
	createDeploymentYaml("", "mount", "", mode, "test4")
	createDeploymentYaml("env", "env", "", mode, "test5")
	createDeploymentYaml("mount", "mount", "", mode, "test6")
	createDeploymentYaml("env", "mount", "", mode, "test7")
	createDeploymentYaml("mount", "env", "", mode, "test8")

}
