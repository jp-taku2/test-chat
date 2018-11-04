package setting

var Server = server{}

type server struct {
	Port string
}

func Load()  {
	Server.Port = ":8080"
}