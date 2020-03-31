package config
import (
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