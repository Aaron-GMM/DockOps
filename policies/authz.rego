package dockops.authz

default allow = false

allow{
    input.role == "admin"
}

allow {
    input.role == "developer"
    input.method == "POST"
    input.path == "/api/v1/containers"
}

allow {
    input.role == "viewer"
    input.method == "GET"
    input.path == "/api/v1/containers"
}



