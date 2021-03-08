package templates

const deploymentAPIVersion, deploymentKind, image, envKey string = "apps/v1", "Deployment", "<image:tag>", "<key>"
const replicas, containerPort int = 3, 80
const mountPath, imagePullPolicy string = "/", "Always"

const secretAPIVersion, configMapAPIVersion, serviceAPIVersion, namespace string = "v1", "v1", "v1", "<namespace>"
const configMapKind, secretKind, secretType, serviceKind string = "ConfigMap", "Secret", "Opaque", "Service"
const usernameKey, username, password, exampleData string = "username", "user", "Pass", "world"
const port, targetPort, nodePort int = 80, 80, 8080

const resourceRequestCPU, resourceRequestMemory, resourceLimitCPU, resourceLimitMemory string = "100m", "100Mi", "1000m", "1000Mi"
const strategyType, emptyDir string = "RollingUpdate", "{}"
