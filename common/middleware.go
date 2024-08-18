package common

var AllowedRoles = map[string][]string{
	`^/api/auth/login$`:           {"POST"},
	`^/api/auth/forgot-password$`: {"POST"},
	`^/api/auth/change-password$`: {"PUT"},
	`^/api/auth/verify$`:          {"POST"},
	`^/api$`:                      {"GET"},
}

const (
	RoleAdmin uint = 1
	RoleUser  uint = 2
)
