package pandemic

type Config struct {
	PathToMap string `env:"PATH_TO_MAP,required"`
	Debug     bool   `env:"DEBUG" envDefault:"false"`
}
