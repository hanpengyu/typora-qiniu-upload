package config

type QiNiuConfig struct {
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Bucket    string `mapstructure:"bucket"`
	BucketDir string `mapstructure:"bucket_dir"`
	CdnUrl    string `mapstructure:"cdn_url"`
}

type ImageCompressConfig struct {
	CompressSwitch int `mapstructure:"compress_switch"`
	MaxKb          int `mapstructure:"max_kb"`
	Width          int `mapstructure:"width"`
	Quality        int `mapstructure:"quality"`
}

type LogConfig struct {
	LogFile string `mapstructure:"log_file"`
}
