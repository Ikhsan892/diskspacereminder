package diskspace

type TelegramConfig struct {
	BaseUrl string `mapstructure:"baseUrl"`
	Token   string `mapstructure:"token"`
	GroupId string `mapstructure:"groupId"`
}

type DiskConfig struct {
	DiskPath      string `mapstructure:"diskpath"`
	MaxPercentage int    `mapstructure:"max_percentage"`
}

type Config struct {
	Telegram TelegramConfig `mapstructure:"telegram"`
	Disk     DiskConfig     `mapstructure:"disk"`
}
