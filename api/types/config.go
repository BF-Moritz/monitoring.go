package types

type ConfigType struct {
	Port   uint32           `yaml:"port"`
	Influx InfluxConfigType `yaml:"influx"`
}

type InfluxConfigType struct {
	URL    string `yaml:"url"`
	APIKey string `yaml:"apikey"`
	Org    string `yaml:"org"`
}
