package main

import (
	"bytes"
	"flag"
	"os"
	"strconv"

	"github.com/ncarlier/webhookd/pkg/auth"
)

// Config contain global configuration
type Config struct {
	ListenAddr          *string
	NbWorkers           *int
	Debug               *bool
	Timeout             *int
	ScriptDir           *string
	Authentication      *string
	AuthenticationParam *string
}

var config = &Config{
	ListenAddr:     flag.String("listen", getEnv("LISTEN_ADDR", ":8080"), "HTTP service address (e.g.address, ':8080')"),
	NbWorkers:      flag.Int("nb-workers", getIntEnv("NB_WORKERS", 2), "The number of workers to start"),
	Debug:          flag.Bool("debug", getBoolEnv("DEBUG", false), "Output debug logs"),
	Timeout:        flag.Int("timeout", getIntEnv("HOOK_TIMEOUT", 10), "Hook maximum delay before timeout (in second)"),
	ScriptDir:      flag.String("scripts", getEnv("SCRIPTS_DIR", "scripts"), "Scripts directory"),
	Authentication: flag.String("auth", getEnv("AUTH", "none"), ""),
	AuthenticationParam: flag.String("auth-param", getEnv("AUTH_PARAM", ""), func() string {
		authdocwriter := bytes.NewBufferString("Authentication method. Available methods: ")

		for key, method := range auth.AvailableMethods {
			authdocwriter.WriteRune('\n')
			authdocwriter.WriteString(key)
			authdocwriter.WriteRune(':')
			authdocwriter.WriteString(method.Usage())
		}
		return authdocwriter.String()
	}()),
}

func init() {
	flag.StringVar(config.ListenAddr, "l", *config.ListenAddr, "HTTP service (e.g address: ':8080')")
	flag.BoolVar(config.Debug, "d", *config.Debug, "Output debug logs")

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv("APP_" + key); ok {
		return value
	}
	return fallback
}

func getIntEnv(key string, fallback int) int {
	strValue := getEnv(key, strconv.Itoa(fallback))
	if value, err := strconv.Atoi(strValue); err == nil {
		return value
	}
	return fallback
}

func getBoolEnv(key string, fallback bool) bool {
	strValue := getEnv(key, strconv.FormatBool(fallback))
	if value, err := strconv.ParseBool(strValue); err == nil {
		return value
	}
	return fallback
}
