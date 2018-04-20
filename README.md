# Handlebars Terraform Provider

This Terraform provides allows to use [handlerbars](http://handlebarsjs.com/) templates, that unlocks the ability to place logic in the templates. 
This provider uses the go library [raymond](https://github.com/aymerick/raymond) that implements the handlerbars template DSL.

### Maintainers

This provider plugin is maintained by [Sedicii](https://sedicii.com/).

### Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

### Installation

```bash
curl https://raw.githubusercontent.com/Sedicii/terraform-provider-handlebars/master/scripts/install-handlebars-tf-pluging.sh | bash
```

### Usage

```
provider "handlebars" {
  version = "~> 0.2.0"
}

data "handlebars_template" "test" {
  template = "${file("${path.module}/templates/test.conf.hbs")}"
  json_context = "${jsonencode("${var.context}")}"
}
```

**For a more detailed example look at the example directory !!**

### Disclaimer on Handlebars behaviour

The handlerbars engine have been modified to play nicely with Terraform.

The differences with a standard engine are :

* The double moustache interpolation `{{var}}` does not scapes variables to html so are the same as the triple moustache `{{{var}}}`
* The if block does some type coercions related on how terraform `jsonencode()` function works (execute the example to see the behaviour)
    * if evaluates `"0"` as false (`"0"` is terraform false)
    * if evaluates `"1"` as true  (`"1"` is terraform true)
    * if evaluates positive numbers `"2"` as true
    * if evaluates negative numbers `"-2"` as false
    * if evaluates empty strings `""` as false
    * if evaluates not empty strings `"text"` as true


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
