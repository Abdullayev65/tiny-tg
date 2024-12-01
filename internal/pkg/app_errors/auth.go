package app_errors

var (
	AuthEmptyUsernameOrEmail = newErr("username or email cannot be empty", 400)
	AuthEmptyPassword        = newErr("password cannot be empty", 400)
	AuthInvalidRole          = newErr("invalid role", 400)
	AuthInvalidEmail         = newErr("invalid email format", 400)
	AuthDuplicateEmail       = newErr("Email already exists", 400)
	AuthDuplicateUsername    = newErr("username already exists", 400)
	AuthLoginDataRequired    = newErr("Username and password are required", 400)
	AuthUserNotFound         = newErr("User not found", 404)
	AuthInvalidPassword      = newErr("Invalid username or password", 401)
)
