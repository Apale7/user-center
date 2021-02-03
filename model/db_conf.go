package model

type Conf struct {
	Mysql Mysql `json:"mysql"`
}
type Mysql struct {
	Dbname   string `json:"dbname"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
}