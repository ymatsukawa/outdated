module "s3" {
  source = "../../../modules/myservice/s3"

  env          = local.env
  project_name = local.project_name
  bucket_name  = "${local.project_name}-${local.env}-example"
}
