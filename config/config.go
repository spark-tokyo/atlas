package config

type AppEnv string

const (
	appPrd   = "prd"
	appTest  = "test"
	appDev   = "dev"
	appLocal = "local"
)

type Config struct {
	Stage      AppEnv
	AppName    string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func NewLoadConfig() (*Config, error) {
	// TODO: 環境を分けるようにする
	// stage := os.Getenv("STAGE")
	// if stage == appLocal || stage == appTest {
	// 	err := godotenv.Load("../local.env")
	// 	if err != nil {
	// 		// 失敗した場合はログを吐いてプログラムを終了
	// 		log.Fatal("Error loading local.env file")
	// 	}
	// } else {
	// 	err := godotenv.Load("../.env")
	// 	if err != nil {
	// 		// 失敗した場合はログを吐いてプログラムを終了
	// 		log.Fatal("Error loading .env file")
	// 	}
	// }

	cfg := &Config{
		Stage:      "local",
		AppName:    "atlas",
		DBUser:     "root",
		DBName:     "atlas",
		DBPassword: "67b9VPFhTcX7",
		DBHost:     "127.0.0.1",
		DBPort:     "3306",
	}
	return cfg, nil
}

func (cfg *Config) IsLocal() bool {
	return cfg.Stage == "local" || cfg.Stage == "test"
}

func (cfg *Config) IsDev() bool {
	return cfg.Stage == "dev"
}
