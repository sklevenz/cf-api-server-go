package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sklevenz/cf-api-server/internal/generated"
	"github.com/sklevenz/cf-api-server/pkg/httpx"
)

const jsonInput = `
{
  "links": {
    "apps": {
      "href": "https://api.example.com/v3/apps",
      "method": "GET"
    },
    "buildpacks": {
      "href": "https://api.example.com/v3/buildpacks",
      "method": "GET"
    },
    "builds": {
      "href": "https://api.example.com/v3/builds",
      "method": "POST"
    },
    "deployments": {
      "href": "https://api.example.com/v3/deployments",
      "method": "GET"
    },
    "domains": {
      "href": "https://api.example.com/v3/domains",
      "method": "GET"
    },
    "droplets": {
      "href": "https://api.example.com/v3/droplets",
      "method": "GET"
    },
    "environment_variable_groups": {
      "href": "https://api.example.com/v3/environment_variable_groups",
      "method": "GET"
    },
    "feature_flags": {
      "href": "https://api.example.com/v3/feature_flags",
      "method": "GET"
    },
    "info": {
      "href": "https://api.example.com/v3/info"
    },
    "isolation_segments": {
      "href": "https://api.example.com/v3/isolation_segments",
      "method": "GET"
    },
    "jobs": {
      "href": "https://api.example.com/v3/jobs",
      "method": "GET"
    },
    "organization_quotas": {
      "href": "https://api.example.com/v3/organization_quotas",
      "method": "GET"
    },
    "organizations": {
      "href": "https://api.example.com/v3/organizations",
      "method": "GET"
    },
    "packages": {
      "href": "https://api.example.com/v3/packages",
      "method": "POST"
    },
    "processes": {
      "href": "https://api.example.com/v3/processes",
      "method": "GET"
    },
    "resource_matches": {
      "href": "https://api.example.com/v3/resource_matches",
      "method": "POST"
    },
    "roles": {
      "href": "https://api.example.com/v3/roles",
      "method": "GET"
    },
    "routes": {
      "href": "https://api.example.com/v3/routes",
      "method": "GET"
    },
    "security_groups": {
      "href": "https://api.example.com/v3/security_groups",
      "method": "GET"
    },
    "self": {
      "href": "https://api.example.com/v3/self",
      "method": "GET"
    },
    "service_brokers": {
      "href": "https://api.example.com/v3/service_brokers",
      "method": "GET"
    },
    "service_credential_bindings": {
      "href": "https://api.example.com/v3/service_credential_bindings",
      "method": "GET"
    },
    "service_instances": {
      "href": "https://api.example.com/v3/service_instances",
      "method": "GET"
    },
    "service_offerings": {
      "href": "https://api.example.com/v3/service_offerings",
      "method": "GET"
    },
    "service_plan_visibilities": {
      "href": "https://api.example.com/v3/service_plan_visibilities",
      "method": "GET"
    },
    "service_plans": {
      "href": "https://api.example.com/v3/service_plans",
      "method": "GET"
    },
    "service_route_bindings": {
      "href": "https://api.example.com/v3/service_route_bindings",
      "method": "GET"
    },
    "service_usage_events": {
      "href": "https://api.example.com/v3/service_usage_events",
      "method": "GET"
    },
    "sidecars": {
      "href": "https://api.example.com/v3/sidecars",
      "method": "GET"
    },
    "space_quotas": {
      "href": "https://api.example.com/v3/space_quotas",
      "method": "GET"
    },
    "spaces": {
      "href": "https://api.example.com/v3/spaces",
      "method": "GET"
    },
    "stacks": {
      "href": "https://api.example.com/v3/stacks",
      "method": "GET"
    },
    "tasks": {
      "href": "https://api.example.com/v3/tasks",
      "method": "POST"
    },
    "users": {
      "href": "https://api.example.com/v3/users",
      "method": "GET"
    }
  }
}
`

// GetRoot handles GET requests to the root endpoint ("/").
// If the request path is empty (r.URL.Path == ""), it redirects permanently to "/".
// Otherwise, it returns an empty JSON response with HTTP 200 OK.
func (Server) GetRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "" {
		// Empty path – redirect to root
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	var root generated.N200Root
	if err := json.Unmarshal([]byte(jsonInput), &root); err != nil {
		panic(err)
	}

	w.Header().Set(httpx.HeaderContentType, httpx.ContentTypeJSON)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(root)
}
