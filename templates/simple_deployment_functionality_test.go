package templates

import "testing"

func TestSimpleDeployment(t *testing.T) {

	CreateSimpleDeploymentYaml("jish", "env", "", "")
}
