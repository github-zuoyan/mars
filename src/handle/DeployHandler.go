package handle

import (
	"fmt"
	"net/http"
)

func DeployHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello , Deploy:")
}
