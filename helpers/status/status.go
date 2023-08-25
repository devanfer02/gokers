package status

import "net/http"

var BadRequest 		= http.StatusBadRequest
var Ok				= http.StatusOK
var Accepted		= http.StatusAccepted
var ServerError 	= http.StatusInternalServerError
var Unauthorized	= http.StatusUnauthorized
var Forbidden 		= http.StatusForbidden
var Created 		= http.StatusCreated
var Conflict		= http.StatusConflict
var NotFound 		= http.StatusNotFound
var TooMany			= http.StatusTooManyRequests