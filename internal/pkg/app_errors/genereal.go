package app_errors

var (
	AccessDenied        = newErr("access denied", 400)
	BadRequest          = newErr("bad request", 400)
	EntityAlreadyExists = newErr("entity already exists", 422)
)
