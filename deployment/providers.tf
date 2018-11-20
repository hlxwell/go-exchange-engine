provider "google" {
  credentials = "${var.gcp_access_key}"
  project     = "planar-sunrise-208601"
  region      = "us-central1"
  zone        = "asia-northeast1"
}
