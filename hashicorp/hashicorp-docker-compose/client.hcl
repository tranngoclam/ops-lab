client {
  enabled = true

  options {
    "docker.privileged.enabled" = "true"
    "driver.raw_exec.enable"    = "1"
  }
}

ports {
  http = 5656
}

plugin "docker" {
  config {
    volumes {
      enabled = true
    }
  }
}

plugin "raw_exec" {
  config {
    enabled = true
  }
}
