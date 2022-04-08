package configs

import (
	"os"
	"strconv"
)

const (
	productionEnv = "production"
	environmentKey = "ENVIRONMENT"
	portKey = "PORT"
	defaultPort = 8080
)

var (
	environment string
	port        int
)

// IsProduction checks the environment name and indicates if it is production or
// not.
func IsProduction() bool {
	return environment == productionEnv
}

// GetPort returns the port value where the server will be listening for new
// connections.
func GetPort() int {
	return port
}

func loadEnvs() {
	environment = os.Getenv(environmentKey)

	if p, err := strconv.Atoi(os.Getenv(portKey)); err != nil {
		port = defaultPort
	} else {
		port = p
	}
}

func init() {
	loadEnvs()
}