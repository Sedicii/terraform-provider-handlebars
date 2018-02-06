# Handlebars Terraform Provider

This Terraform provides allows to use [handlerbars](http://handlebarsjs.com/) templates, that unlocks the ability to place logic in the templates. 
This provider uses the go library [raymond](https://github.com/aymerick/raymond) that implements the handlerbars template DSL.

### Maintainers

This provider plugin is maintained by [Sedicii](https://sedicii.com/).

### Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

### Installation

### Usage

```
provider "handlebars_template" {
  version = "~> 0.1.0"
}






```



### Building The Provider

Clone repository to: `$GOPATH/src/github.com/sedicii/terraform-provider-handlebars`

```sh
$ mkdir -p $GOPATH/src/github.com/sedicii; cd $GOPATH/src/github.com/sedicii
$ git clone git@github.com:sedicii/terraform-provider-handlebars
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/sedicii/terraform-provider-handlebars
$ make build
```

### Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-handlebars
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```
