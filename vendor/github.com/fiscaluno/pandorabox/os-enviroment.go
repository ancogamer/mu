package pandorabox

import (
	"os"
	"strings"
)

// GetOSEnvironment get a env var or set to default
func GetOSEnvironment(env string, defaultEnv string) string {
	if e := os.Getenv(env); e != "" {
		return strings.ToLower(os.Getenv(env))
	}
	return defaultEnv
}
