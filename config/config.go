package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

const CLIENT_HOST = "client.host"
const CLIENT_PORT = "client.port"

const ENV_PREFIX = "YGGY"

const HOST = "server.host"
const PORT = "server.port"
const PROXY = "server.proxy"
const ENV = "env"
const INFLUX_URL = "influx.url"
const INFLUX_ORG = "influx.org"
const INFLUX_BUCKET = "influx.bucket"
const INFLUX_TOKEN = "influx.token"

func InfluxServer() (url string, token string, org string, bucket string) {
	url = viper.GetString(INFLUX_URL)
	token = viper.GetString(INFLUX_TOKEN)
	org = viper.GetString(INFLUX_ORG)
	bucket = viper.GetString(INFLUX_BUCKET)
	return
}

func HostAndPort() string {
	return viper.GetString(HOST) + ":" + viper.GetString(PORT)
}

func TrustedProxies() []string {
	str := viper.GetString(PROXY)
	return strings.Split(str, ",")
}

func init() {
	viper.SetDefault(HOST, "0.0.0.0")
	viper.SetDefault(PORT, "9000")
	viper.SetDefault(PROXY, "0.0.0.0/0")
	viper.SetDefault(INFLUX_URL, "localhost:8086")
	viper.SetDefault(INFLUX_TOKEN, "")
	viper.SetDefault(INFLUX_BUCKET, "yggy")
	viper.SetDefault(INFLUX_ORG, "yggy")
	viper.SetDefault(ENV, "dev") // dev/prod
	load()
}

func load() {
	log.SetLevel(log.DebugLevel)
	viper.SetEnvPrefix(ENV_PREFIX)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if viper.GetString(ENV) != "dev" {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	} else {
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			log.Warn("config file does not exist. You can either add a config.yaml file, or set the " + ENV_PREFIX + "_ENV envVar to prod to not see this message anymore")
		} else {
			//seeeeeeeelf update config :wink: :wink: if you add new keys.
			viper.WriteConfig()
			log.Debug("Updating configuration")
		}
	}

}
