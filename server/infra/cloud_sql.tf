resource "google_sql_database_instance" "mattbutterfield" {
  name             = "mattbutterfield"
  region           = var.default_region
  database_version = "POSTGRES_13"

  settings {
    tier      = "db-f1-micro"
    disk_size = 10
  }
}

resource "google_sql_database" "heatmap-creator" {
  name     = "heatmap-creator"
  instance = google_sql_database_instance.mattbutterfield.name
}

resource "google_sql_user" "heatmap-creator" {
  name     = "heatmap-creator"
  instance = google_sql_database_instance.mattbutterfield.name
  password = var.db_password
}

resource "google_project_iam_member" "heatmap-creator_cloud_run_cloud_sql" {
  project = var.project
  role    = "roles/cloudsql.client"
  member  = "serviceAccount:${google_service_account.heatmap-creator_cloud_run.email}"
}
