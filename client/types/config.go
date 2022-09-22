package types

type ConfigType struct {
	ServerURL           string `yaml:"serverurl"`
	Name                string `yaml:"name"`
	TimeBetweenReadings int64  `yaml:"timeBetweenReadings"`
}
