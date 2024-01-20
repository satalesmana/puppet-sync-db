package config

type Config struct {
	App           App        `yaml:"app"`
	MongoDBRemote MongoDB    `yaml:"mongo_db_remote"`
	MongoDBLocal  MongoDB    `yaml:"mongo_db_local"`
	Collection    Collection `yaml:"collection"`
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
	Remote RemoteCollection `yaml:"remote"`
	Local  LocalCollection  `yaml:"local"`
}

type RemoteCollection struct {
	PrimaryCollection   string `yaml:"primary"`
	SecondaryCollection string `yaml:"secondary"`
}

type LocalCollection struct {
	ActivityCollection string `yaml:"activity"`
	DatabaseCollection string `yaml:"database"`
}
