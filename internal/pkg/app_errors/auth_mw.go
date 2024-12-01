package app_errors

var (
	AuthMwMissingToken = newErr("Missing token", 401)
)
