package configs

var (
	ConfigEnvGlobalInstance ConfigEnvItf
)

func NewConfigEnvGlobalInstance(configPath string) {
	ConfigEnvGlobalInstance = newViperConfigLoaderEnv(configPath)
}
