package rotas

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                    "/v1/usuarios",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/v1/usuarios",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUsers,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/v1/usuarios/{usuarioId}",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/v1/usuarios/{usuarioId}",
		Method:                 http.MethodPut,
		Function:               controllers.ChangeUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/v1/usuarios/{usuarioId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: false,
	},
}
