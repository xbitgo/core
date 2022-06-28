package env

import (
	"flag"
	"os"
	"strconv"
)

// deploy env.
const (
	EnvironmentLOCAL = "local"
	EnvironmentDEV   = "dev"
	EnvironmentTEST  = "test"
	EnvironmentPRE   = "pre"
	EnvironmentPROD  = "prod"
)

var (
	// Environment 环境
	Environment string
	// ServiceName 服务名
	ServiceName string
	// Debug 调试模式
	Debug bool
)

func init() {
	addFlag(flag.CommandLine)
}

func addFlag(fs *flag.FlagSet) {
	fs.StringVar(&Environment, "env", defaultString("Environment", EnvironmentDEV), "环境变量")
	fs.StringVar(&ServiceName, "service.name", defaultString("ServiceName", "default_service_name"), "服务名")
	fs.BoolVar(&Debug, "debug", defaultBool("Debug", false), "set debug")
}

func defaultString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func defaultBool(key string, defaultVal bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		ret, _ := strconv.ParseBool(val)
		return ret
	}
	return defaultVal
}
