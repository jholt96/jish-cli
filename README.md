root command = jish

commands = generate

sub commands= deployment, configmap, secret, service


deployment flags:
config = env or volume
secret = env or volume
service = clusterIP, nodeport, or loadbalancer
mode = simple, standard, complex

complex does every possible deployment method, standard does normal fields, and simple does the bare minimum. 
resources defines what kind of other resources that will be created with the deployment. 

the deployment flags determine what other resources are to be created with the deployment yaml that are already set up to be used within the deployment yaml or in its own yaml. 

Example: jish generate deployment nginx --mode simple --configmap env --secret mount --service clusterip

This command will create a simple deployment yaml called nginx that has a configmap referenced as a environment variable and a secret mounted as a volume. It will also create a service using the app=nginx label. It will lastely create the yaml definitions for the configmap, secret, service, and deployment. 