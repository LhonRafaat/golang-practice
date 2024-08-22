package server

import "lhon/postgres-rest/internal/router"




func StartServer() {

	r:= router.SetupRouter()
	r.Run()
}