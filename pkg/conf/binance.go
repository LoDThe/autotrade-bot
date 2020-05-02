package conf

type Binance struct {
	APIKey string `env:"BINANCE_API_KEY,required"`
	Secret string `env:"BINANCE_SECRET,required"`
	Debug  bool   `env:"BINANCE_DEBUG" envDefault:"true"`
}
