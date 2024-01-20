package configupdateremoteflag

type Config struct {
	App        App        `yaml:"app"`
	MongoDB    MongoDB    `yaml:"mongo_db_remote"`
	Collection Collection `yaml:"collection"`
}

type App struct {
	ENV      string `yaml:"env"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	TimeZone string `yaml:"timezone"`
}

type MongoDB struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Collection struct {
	PrimaryCollection   string `yaml:"primary"`
	SecondaryCollection string `yaml:"secondary"`
}
