package decorators

import (
	"simplesurveygo/dao"
	"reflect"
	"fmt"
)


func validate_token(token string) bool {
	session_info := dao.GetSessionDetails(token)
    if (reflect.DeepEqual(session_info,(dao.Session{}))) {
		return true
	} else {
		return false
	}
}

