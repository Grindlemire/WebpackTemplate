package start

import (
	"io"
	SYS "syscall"

	"github.com/grindlemire/WebpackTemplate/rest"
	DEATH "github.com/vrecan/death"
)

// Run starts the webserver
func Run() {
	var goRoutines []io.Closer
	death := DEATH.NewDeath(SYS.SIGINT, SYS.SIGTERM)

	restService := rest.NewRestService()
	goRoutines = append(goRoutines, restService)
	restService.Start()

	death.WaitForDeath(goRoutines...)

}
