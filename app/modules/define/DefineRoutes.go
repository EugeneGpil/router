package define

import (
	"net/http"

	"github.com/EugeneGpil/router/app/modules/define/methods"
)

func DefineRoutes(mux *http.ServeMux) {
	groupedByUrlRoutes := methods.GetGroupedByUrlRoutes()

	methods.DefineGroupedByUrlRoutes(groupedByUrlRoutes, mux)
}
