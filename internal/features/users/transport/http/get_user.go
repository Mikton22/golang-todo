package users_transport_http

import (
	"net/http"

	core_logger "github.com/Mikton22/golang-todo/internal/core/logger"
	core_http_response "github.com/Mikton22/golang-todo/internal/core/logger/transport/http/middleware/response"
	core_http_request "github.com/Mikton22/golang-todo/internal/core/logger/transport/http/request"
)

type GetUserResponse UserDTOResponse

// GetUser godoc
// @Summary Получение пользователя
// @Description Возвращает пользователя по идентификатору
// @Tags users
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} GetUserResponse "Пользователь успешно получен"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad Request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal Server Error"
// @Router /users/{id} [get]
func (h *UsersHttpHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userId, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userId path value")
	}
	user, err := h.usersService.GetUser(ctx, userId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	response := GetUserResponse(userDtoFromDomain(user))
	responseHandler.JSONResponse(response, http.StatusOK)
}
