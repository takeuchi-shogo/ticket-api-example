package controllers

type Context interface {
	BindJSON(obj interface{}) error
	JSON(code int, obj interface{})
	MustGet(key string) interface{}
	Param(key string) string
}
