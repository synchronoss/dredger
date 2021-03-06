# dredger

Dredger is a utility to help convert helm charts to Terraform modules using kubernetes provider.

Dredger is made of dark magic and cannot fully convert a helm chart. It is designed to perform the bulk of the work but will still require some knowledge of Terraform HCL.

## quick start
```
# Run dredger against the bitnami helm chart and write module to /tmp/dredger-redis
dredger --outputdir /tmp/dredger-redis helm --repo https://charts.bitnami.com/bitnami redis-cluster --set networkPolicy.enabled=true

# Now create the terraform to consume this module

mkdir /tmp/dredger-main
cd /tmp/dredger-main

cat > main.tf <<EOF
terraform {
  backend "kubernetes" {
    namespace        = "kube-system"
    secret_suffix    = "dredger-quick-start-state"
    load_config_file = true
  }
}

provider "kubernetes" {} # Configure the kubernetes provider

module "redis" {
  source = "/tmp/dredger-redis"
  name = "myredis"
  namespace = "myredis"
  redis-cluster-replicas = 3
  redis-cluster-secrets = {
    redis-password = "JustAnExamplePassword"
  }
}
EOF
# Now we just init and apply!
terraform init
terraform apply
```

## building
Make sure that go have the go compiler installed on at least version 1.16.

```sh
make
```

## usage
To convert the chart foo in repository bar, run the folowing command:
```sh
dredger helm bar/foo

# Alternatively to split output into a directory
dredger --outputdir /tmp/foo_terraform/ helm bar/foo

# Any arguments specified after the helm option are passed into helm
dredger helm bar/foo -f values.yaml
```

## injecting variables
In some cases you may want to inject a variable into the output module. If the variable is scalar then the easiest way to do this is by injecting via the helm values, causing the interpolated variable to appear in the resulting module.
```
dredger helm --set 'some.var=${var.some_var}'
```

## configuring
To alter the way that dredger generates terraform modules you first need to create a config file. You can start with dredger's in-built default config by running this command:
```
dredger --dumpconfig > /tmp/custom_dredger_config.yaml
```
Then edit the output file makeing any alterations you require. Details on how the configuration works can be found at the top of the dumped config.

When you want to execute dredger with this config use the flag `--config` to specify the path to your custom file.
```
dredger --config /tmp/custom_dredger_config.yaml helm bar/foo
```
