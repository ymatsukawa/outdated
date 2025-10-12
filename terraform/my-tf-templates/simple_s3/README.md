# simple s3 - example

Bucket will be created `{project-name}-{env}-example`

* variable declaration(interface) is in `modules/myservice/s3/variable.tf`
* variable setting(implementation) is in `env/myservice/s3/{env}/locals.tf`

`init`, `plan` and `apply` should be run under `env/{service_name}/{env}`
