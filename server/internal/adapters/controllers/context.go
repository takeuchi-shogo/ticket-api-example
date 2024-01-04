package controllers

type Context interface {
	BindJSON(obj interface{}) error
	Header(key, value string)
	JSON(code int, obj interface{})
	MustGet(key string) interface{}
	Param(key string) string
	PostForm(key string) string
	Query(key string) string
	ShouldBind(obj interface{}) error
}
