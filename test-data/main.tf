variable "instance_type" {
  description = "Type of instance to run"
}

variable "region" {
  description = "Region to deploy in"
}

resource "aws_instance" "example" {
  ami           = "ami-123456"
  instance_type = var.instance_type
}
