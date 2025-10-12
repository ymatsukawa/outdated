# when you changed properties and terraform plan,
# plan shows diff

resource "aws_route53_zone" "primary" {
  name = "i-changed.com"
}

resource "aws_route53_record" "www" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "www.i-changed.com"
  type    = "A"
  ttl     = "300"
  records = ["192.168.0.1"]
}
