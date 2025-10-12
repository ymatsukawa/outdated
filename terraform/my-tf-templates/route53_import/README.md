# route53 import

## UseCase

Route53 host zone and record is already created.

And you want to import them to `terraform.tfstate`

## Operation

Prepare blank resource.
ex.) [main_before.tf](./env/prod/route53/main_before.tf)

and run follow commands

```
$ cd ./env/prod/route53
$ terraform init
$ terraform import aws_route53_zone.primary {host_zone_id}
$ terraform import aws_route53_zone.primary {host_zone_id}_www.example.com_A
```

`{host_zone_id}` is shown on AWS managed console.
