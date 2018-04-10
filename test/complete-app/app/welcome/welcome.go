package welcome

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Welcome to the Go toolset.")

}
