root command = jish

commands = generate

sub commands= deployment, pod, statefulset, configmap, secret, service, ingress, serviceaccount, role, rolebinding


deployment flags:
config = env or volume
secret = env or volume
service = clusterIP, nodeport, or 
type = simple, standard, complex

complex does every possible deployment method, standard does the usual, and simple does the bare minimun. 
resources defines what kind of other resources that will be created with the deployment. 