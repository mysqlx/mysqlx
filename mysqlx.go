package mysqlx

type Mysqlx struct {
}

//Config
type SessionSettings struct {
	Host string
	Port int
	DbUser string
	DbPassword string
}


func  GetSession(config SessionSettings)(BaseSession,error ) {
	return BaseSession{},nil
}

func  GetSessionByURL(url string) (BaseSession,error ) {
	msqlxConfig,err:= parseURL(url)
	if err!=nil {
		return nil,err
	}
	return GetSession(msqlxConfig)
}


func  GetNodeSession(config SessionSettings) (NodeSession,error) {
	return NodeSession{},nil
}

func  GetNodeSessionByURL(url string) (NodeSession,error) {
	msqlxConfig,err:= parseURL(url)
	if err!=nil {
		return nil,err
	}
	return GetNodeSession(msqlxConfig)
}

func parseURL(url string) (SessionSettings,error) {
	return SessionSettings{},nil
}

