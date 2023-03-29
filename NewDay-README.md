# NewDay build instructions

## Linux (platform engineering image)

Make sure go is installed

Add gopath if not present

```sh
export GOPATH=$(go env GOPATH)
```

Build

```sh
$ make tools
...
$ make build
...
$ $GOPATH/bin/terraform-provider-azurerm
...
```

**Note:** If you have issues running make build, change the line ending of the files in /scripts to be unix (LF).

Copy file to the expected local library, adjust for the version:

```sh
cp /root/go/bin/terraform-provider-azurerm ~/.terraform.d/plugins/registry.terraform.io/hashicorp/azurerm/<version number>/linux_amd64
```

Or

```sh
cp /root/go/bin/terraform-provider-azurerm ~/.terraform.d/plugins/registry.terraform.io/hashicorp/azurerm/2.75.0-nd3/linux_amd64
```
