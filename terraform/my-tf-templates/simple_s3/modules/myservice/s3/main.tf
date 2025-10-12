resource "aws_s3_bucket" "myservice_static" {
  bucket = var.bucket_name
}

resource "aws_s3_bucket_public_access_block" "private" {
  bucket                  = aws_s3_bucket.myservice_static.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_ownership_controls" "bucket_owner_preferred" {
  bucket = aws_s3_bucket.myservice_static.id

  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_versioning" "tracking" {
  bucket = aws_s3_bucket.myservice_static.id
  versioning_configuration {
    mfa_delete = "Disabled"
    status     = "Enabled"
  }
}
