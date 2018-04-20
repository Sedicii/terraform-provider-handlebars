

provider "handlebars" {
  version = "~> 0.2.1"
}


variable "context" {
  default = {
    name = "test_name"
    ips = ["10.0.0.1","10.1.0.1","10.2.0.1"]
    boolean_true = true
    boolean_false = false
    number = 1
    float = 1.1
    map = {
      k1 = 1
      k2 = 2
      k3 = "test"
    }
  }
  type = "map"
}

data "handlebars_template" "test" {
  template = "${file("${path.module}/templates/test.conf.hbs")}"
  json_context = "${jsonencode("${var.context}")}"
}

output "context" {
  value = "${jsonencode("${var.context}")}"
}

output "test_rendered" {
  value = "${data.handlebars_template.test.rendered}"
}