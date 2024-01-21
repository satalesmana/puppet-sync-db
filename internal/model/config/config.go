package config

type Config struct {
	App           App        `yaml:"app"`
	MongoDBRemote MongoDB    `yaml:"mongo_db_remote"`
	MongoDBLocal  MongoDB    `yaml:"mongo_db_local"`
	Collection    Collection `yaml:"collection"`
	Synclimit     int64      `yaml:"synclimit"`
	Email         MailConfig `yaml:"mail_config"`
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
	Source string `yaml:"source"`
}

type LocalCollection struct {
	ActivityCollection string `yaml:"activity"`
	DatabaseCollection string `yaml:"database"`
}

type MailConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	SenderName   string `yaml:"sender_name"`
	AuthEmail    string `yaml:"auth_email"`
	AuthPassword string `yaml:"auth_password"`
	Template     string `yaml:"template"`
	Subject      string `yaml:"subject"`
}
