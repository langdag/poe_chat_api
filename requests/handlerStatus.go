package requests

import (
	"fmt"
	"net/http"
)

func HandlerStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK! I'm alive!")
}
