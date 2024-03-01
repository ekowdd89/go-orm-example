package repository

type RepoConfigs struct {
	DbConfig DBConfig
}

type DBConfig struct {
	Host,
	User,
	Pass,
	Port,
	DbName string
}
