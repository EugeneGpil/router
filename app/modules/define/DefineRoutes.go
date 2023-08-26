package define

import (
	"net/http"

	"github.com/EugeneGpil/router/app/modules/define/methods"
	internalDefine "github.com/EugeneGpil/router/app/modules/define/methods/define"
)

func DefineRoutes(mux *http.ServeMux) {
	groupedByUrlRoutes := methods.GetGroupedByUrlRoutes()

	internalDefine.DefineGroupedByUrlRoutes(groupedByUrlRoutes, mux)
}
