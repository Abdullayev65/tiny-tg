package config

import (
	"os"
	"strconv"
	"time"

	"github.com/go-shafaq/timep"
	"github.com/joho/godotenv"
)

var (
	PORT = "80"

	POSTGRES_URI = "postgres://postgres:password@localhost:5432/tg?sslmode=disable"

	JWT_SIGNING_KEY     = "[.jnjnhj,>l97vg34xtg]hbn>{k}^&"
	JWT_EXPIRY_DURATION = 10 * 24 * time.Hour
	UPLOADS_DIR         = "./uploads"
)

func LoadVarsFromEnv() {
	setIfExistsStr(&PORT, "PORT")

	setIfExistsStr(&POSTGRES_URI, "POSTGRES_URI")

	setIfExistsStr(&JWT_SIGNING_KEY, "JWT_SIGNING_KEY")
	setIfExistsDur(&JWT_EXPIRY_DURATION, "JWT_EXPIRY_DURATION")

	setIfExistsStr(&UPLOADS_DIR, "JWT_EXPIRY_DURATION")

}

func setIfExists[V any](ptr *V, key string, parser func(string) (V, bool)) bool {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return false
	}

	val, ok := parser(envVal)
	if !ok {
		return false
	}

	*ptr = val

	return ok
}

func setIfExistsStr(ptr *string, key string) bool {
	return setIfExists(ptr, key,
		func(s string) (string, bool) { return s, true })
}

func setIfExistsBool(ptr *bool, key string) bool {
	return setIfExists(ptr, key,
		func(s string) (bool, bool) {
			b, err := strconv.ParseBool(s)
			return b, err == nil
		})
}

func setIfExistsDur(ptr *time.Duration, key string) bool {
	return setIfExists(ptr, key,
		func(s string) (time.Duration, bool) {
			dur, err := timep.ParseDuration(s)
			return dur, err == nil
		})
}

func init() {
	godotenv.Load(".env")

	LoadVarsFromEnv()
}
