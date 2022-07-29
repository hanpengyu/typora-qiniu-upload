package config

func GetQnCfg() QiNiuConfig {
	v := GetConfig().Sub("qiniu")
	if v == nil {
		panic("no qiniu config")
	}

	var cfg QiNiuConfig
	err := v.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func GetImageCompressCfg() ImageCompressConfig {
	v := GetConfig().Sub("image")
	if v == nil {
		panic("no image config")
	}

	var cfg ImageCompressConfig
	err := v.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func GetLogCfg() LogConfig {
	v := GetConfig().Sub("logger")
	if v == nil {
		panic("no log config")
	}

	var cfg LogConfig
	err := v.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
