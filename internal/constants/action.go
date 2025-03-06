package constants

type Action struct {
	CREATE string
	UPDATE string
	DELETE string
}

var ACTION = Action{
	CREATE: "CREATE",
	UPDATE: "UPDATE",
	DELETE: "DELETE",
}
