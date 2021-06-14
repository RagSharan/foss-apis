package main

import (
	"fmt"
	"net/http"

	"github.com/ragsharan/foss-apis/router"
)

var (
	httpRouter  router.IRouter  = router.NewMuxRouter()
	httpMapping router.IMapping = router.InstMapping()
)

func main() {
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "server up and running...")
	})
	httpMapping.UriMappings()
	//	httpRouter.SERVE(os.Getenv(":PORT"))
	httpRouter.SERVE(":8000")
}
