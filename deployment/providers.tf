provider "google" {
  credentials = "${file("~/.gcp/juju-gce-1-sa.json")}"
  project     = "bitsx-vc-dev-poc"
  zone        = "asia-northeast1-c"
  region      = "asia-northeast1"
}
