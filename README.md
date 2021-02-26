root command = jish

commands = generate

sub commands= deployment, pod, statefulset, configmap, secret, service, ingress, serviceaccount, role, rolebinding


deployment flags:
resources = all,config,secret,service, none(default)
type = simple, standard, complex


type defines what kind of parameters are in the yaml. complex does every possible deployment method, standard does the usual, and simple does the bare minimun. 
resources defines what kind of other resources that will be created with the deployment. 