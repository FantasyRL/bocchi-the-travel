package config

type server struct {
	Secret  []byte
	Version string
	Name    string
}

type service struct {
	Name     string
	AddrList []string
	LB       bool `mapstructure:"load-balance"`
}

type etcd struct {
	Addr string
}

type mySQL struct {
	Addr     string
	User     string
	Password string
	Database string
}
type redis struct {
	Addr string
}

type oss struct {
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	MainDirectory   string `mapstructure:"main-directory"`
}

type rabbitMQ struct {
	Addr     string
	User     string
	Password string
}

type email struct {
	Host     string
	Port     string
	Sender   string
	Password string
	From     string
	Subject  string
}

type config struct {
	Server   server
	MySQL    mySQL
	Etcd     etcd
	RabbitMQ rabbitMQ
	Redis    redis
	OSS      oss
	Email    email
}
