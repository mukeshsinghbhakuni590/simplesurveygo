package servicehandlers

import (
	"fmt"
	"net/http"
	"simplesurveygo/dao"
	//"simplesurveygo/decorators"
)

type SessionHandler struct {
}

func (p SessionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p SessionHandler) Get(r *http.Request) SrvcRes {
	tokens, _ := r.Header["Token"]

	fmt.Println("TOKENS : ", tokens)
	fmt.Println("HEADERS : ", r.Header)

	token := tokens[0]
	if dao.validate_token(token) {
    user := dao.GetSessionDetails(token)
	return Response200OK(user)
	} else {
		return  ResponseNotImplemented()
	}
}

func (p SessionHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p SessionHandler) Post(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}
