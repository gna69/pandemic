package pandemic

type Config struct {
	Port      int    `env:"PORT,required"`
	PathToMap string `env:"PATH_TO_MAP,required"`
	Debug     bool   `env:"DEBUG" envDefault:"false"`
}
