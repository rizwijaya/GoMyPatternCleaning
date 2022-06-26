package config

type Conf struct {
	App struct {
		Name string `env:"APP_NAME"`
		Port string `env:"APP_PORT"`
	}
	Db struct {
		Host string `env:"DB_HOST"`
		Name string `env:"DB_NAME"`
		User string `env:"DB_USER"`
		Pass string `env:"DB_PASSWORD"`
		Port string `env:"DB_PORT"`
	}
}