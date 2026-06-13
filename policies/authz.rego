package dockops.authz

import rego.v1

default allow := false

allow if {
    input.role == "admin"
}

allow if {
    input.role == "developer"
    input.method == ["POST", "GET"][_]
    startswith(input.path, "/api/v1/containers")
}

allow if {
    input.role == "viewer"
    input.method == "GET"
    startswith(input.path, "/api/v1/containers")
}
