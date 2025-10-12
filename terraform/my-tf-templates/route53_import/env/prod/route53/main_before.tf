# prepare blank resource and run script
# $ terraform init
# $ terraform import aws_route53_zone.primary {host_zone_id}
# $ terraform import aws_route53_zone.primary {host_zone_id}_www.example.com_A

# ref. https://registry.terraform.io/providers/rgeraskin/aws2/latest/docs/resources/route53_zone#import
resource "aws_route53_zone" "primary" {
}

# ref. https://registry.terraform.io/providers/rgeraskin/aws2/latest/docs/resources/route53_record#import
resource "aws_route53_record" "www" {
}
