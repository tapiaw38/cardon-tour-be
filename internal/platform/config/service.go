package config

var configService *ConfigurationService

const (
	ReleaseMode GinModeServer = "release"
	DebugMode   GinModeServer = "debug"
)

type (
	GinModeServer string

	ConfigurationService struct {
		ServerConfig ServerConfig
		DBConfig     DBConfig
		S3Config     S3Config
	}

	ServerConfig struct {
		GinMode   GinModeServer
		Port      string
		Host      string
		JWTSecret string
	}

	DBConfig struct {
		DatabaseURL string
	}

	S3Config struct {
		AWSRegion          string
		AWSAccessKeyID     string
		AWSSecretAccessKey string
		AWSBucket          string
	}
)

func InitConfigService(config *ConfigurationService) {
	if configService == nil {
		configService = config
	}
}

func GetConfigService() ConfigurationService {
	return *configService
}
