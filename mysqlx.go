package mysqlx

import "fmt"

type Mysqlx struct {
}

//Config
type SessionSettings struct {
	Host       string
	Port       int
	DbUser     string
	DbPassword string
	Schema     string //default connect Schema
}

//interface Stringer
func (this *SessionSettings) String() string {
	return fmt.Sprintf("%s:%s@%s:%i",
		this.DbUser,
		this.DbPassword,
		this.Host,
		this.Port,
	)
}

func GetSession(config SessionSettings) (*XSession, error) {
	return &XSession{}, nil
}

func GetSessionByURL(url string) (*XSession, error) {
	msqlxConfig, err := parseURL(url)
	if err != nil {
		return nil, err
	}
	return GetSession(msqlxConfig)
}

func GetNodeSession(config SessionSettings) (*NodeSession, error) {
	return &NodeSession{}, nil
}

func GetNodeSessionByURL(url string) (*NodeSession, error) {
	msqlxConfig, err := parseURL(url)
	if err != nil {
		return nil, err
	}
	return GetNodeSession(msqlxConfig)
}

func parseURL(url string) (SessionSettings, error) {
	return SessionSettings{}, nil
}
