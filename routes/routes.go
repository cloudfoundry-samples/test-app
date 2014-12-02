package routes

import "github.com/tedsuo/rata"

const (
	Env   = "ENV"
	Hello = "HELLO"
	Exit  = "EXIT"
)

var Routes = rata.Routes{
	{Path: "/", Method: "GET", Name: Hello},
	{Path: "/env", Method: "GET", Name: Env},
	{Path: "/exit", Method: "GET", Name: Exit},
}
