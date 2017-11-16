package errors

import (
	"errors"
	"strings"
	"net/http"
)

const (
	StatusBadParameter = "Bad parameter"
	StatusUnknow = "Unknow"
	StatusIncorrectJson = "Incorrect json"
	StatusNotUnique = "Not Unique"
	StatusForbidden = "Forbidden"
)

type Err struct {
	Code string
	Attr string
	origin error
	http *Http
}
//-------------------------------------------------------------------------------------------------


func BadParametr(attr string, e ...error) *Err {
	return  &Err{
		Code: StatusBadParameter,
		Attr: attr,
		origin: getError(attr + ": bad parameter", e...),
		http: HTTP.getBadParameter(attr),
	}
}

func IncorectJson( e ...error) *Err {
	return  &Err{
		Code: StatusIncorrectJson,
		origin: getError("incorect json", e...),
		http: HTTP.getIncorrectJSON(),
	}
}

func Forbidden( e ...error) *Err {
	return  &Err{
		Code: StatusForbidden,
		origin: getError("forbidden", e...),
		http: HTTP.getForbidden(),
	}
}

func Unknown( e ...error) *Err {
	return  &Err{
		Code: StatusUnknow,
		origin: getError("unknown error", e...),
		http: HTTP.getUnknown(),
	}
}

func (e *Err) Err() error {
	return e.origin
}

func (e *Err) Http(w http.ResponseWriter){
	e.http.send(w)
}
type err struct{
	s string
}

func New(text string) *err{
	return &err{text}
}

func (e *err) Error() string {
	return e.s
}

func (e *err) Unathorized(err ...error) *Err {
	return &Err{
		Code: http.StatusText(http.StatusUnauthorized),
		origin: getError(joinNameAndMessage(e.s, "access denied"), err...),
		http: HTTP.getUnauthorized(),
	}
}

func (e *err) NotFound(err ...error) *Err {
	return &Err{
		Code: http.StatusText(http.StatusNotFound),
		origin: getError(joinNameAndMessage(e.s, "not found"), err...),
		http: HTTP.getNotFound(e.s),
	}
}

func (e *err) NotUnique(attr string, err ...error) *Err {
	return &Err{
		Code: StatusNotUnique,
		origin: getError(joinNameAndMessage(e.s, strings.ToLower(attr) + " not unique"), err...),
		http: HTTP.getNotUnique(strings.ToLower(attr)),
	}
}

func (e *err) BadParameter(attr string, err ...error) *Err {
	return  &Err{
		Code: StatusBadParameter,
		Attr: attr,
		origin: getError(joinNameAndMessage(e.s, "bad parameter" + strings.ToLower(attr)), err...),
		http: HTTP.getBadParameter(attr),
	}
}

func (e *err) IncorrectJSON(err ...error) *Err {
	return  &Err{
		Code: StatusIncorrectJson,
		origin: getError(joinNameAndMessage(e.s, "incorrect json") , err...),
		http: HTTP.getIncorrectJSON(),
	}
}

func (e *err) Forbidden(err ...error) *Err {
	return  &Err{
		Code: StatusForbidden,
		origin: getError(joinNameAndMessage(e.s, "forbidden") , err...),
		http: HTTP.getForbidden(),
	}
}

func (e *err) Unknown(err ...error) *Err {
	return  &Err{
		Code: StatusUnknow,
		origin: getError(joinNameAndMessage(e.s, "unknown error") , err...),
		http: HTTP.getUnknown(),
	}
}

//------------------------------------------------------------------------------------------------
func getError(msg string, err ...error) error{
	if len(err) == 0 {
		return errors.New(msg)
	} else {
		return err[0]
	}
}

func joinNameAndMessage(name, message string) string{
	return toUpperFirstChar(name) + ": " + message
}

func toUpperFirstChar(str string) string {
	return strings.ToUpper(str[0:1] + str[1:])
}