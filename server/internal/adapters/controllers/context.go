package controllers

import "net/http"

type Context interface {
	BindJSON(obj interface{}) error
	GetPostForm(key string) (string, bool)
	GetRawData() ([]byte, error)
	Header(key, value string)
	JSON(code int, obj interface{})
	MustGet(key string) interface{}
	Param(key string) string
	PostForm(key string) string
	Query(key string) string
	SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	SetSameSite(samesite http.SameSite)
	ShouldBind(obj interface{}) error
}
