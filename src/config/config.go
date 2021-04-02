package config
import (
	"strings"
	"time"
)

// 新建一个配置对象
func New(filename string, autoLoad time.Duration) (*Config) {
	var config  = Config {
		fileName: filename,
		config: make(map[string] sectionType),
		autoLoad: autoLoad,
		stringMap: make(map[string] autoLoadString),
		int64Map: make( map[/*section:key*/ string] autoLoadInt64),
		int32Map: make( map[/*section:key*/ string] autoLoadInt32),
		uint64Map: make( map[/*section:key*/ string] autoLoadUint64),
		uint32Map: make( map[/*section:key*/ string] autoLoadUint32),
		float32Map: make( map[/*section:key*/ string] autoLoadFloat32),
		float64Map: make( map[/*section:key*/ string] autoLoadFloat64),
		boolMap: make( map[/*section:key*/ string] autoLoadBool),
		sectionMap: make(map[/*section*/ string] autoLoadSection),
	}
	
	loadFile(&config)
	if autoLoad > 0 {
		go config.autoLoadConfig()
	}
	return &config
}

func loadFile(config *Config) {

}

/**
	自动载入机制
 */
func (config *Config) autoLoadConfig() {

}

// 获取字符串型的配置
func (config *Config) GetStrConfig(section string, key string, def string, value *string) bool {
	
	var configVal = def
	var isExists = false
	if len(config.config[section]) > 0 {
		if len(config.config[section][key]) > 0 {
			configVal = config.config[section][key]
			isExists = true
		}
	}
	
	if !isExists {
		// 配置不存在，尝试从环境变量中获取
		envConfig := os.Getenv(strings.ToUpper(section + "_" + key))
		if envConfig != "" {
			configVal = envConfig
			isExists = true
		}
	}
	
	if config.autoLoad > 0 && isExists {
		config.stringMap[section+":"+key] = autoLoadString{
			section: section,
			key: key,
			def: def,
			value: value,
		}
	}
	*value = configVal
	return isExists
}