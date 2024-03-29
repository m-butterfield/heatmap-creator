resource "google_storage_bucket" "content" {
  name     = "heatmap-creator-content"
  location = "US"

  cors {
    origin          = ["*"]
    method          = ["GET", "HEAD", "PUT", "POST", "DELETE"]
    response_header = ["*"]
    max_age_seconds = 3600
  }
}
