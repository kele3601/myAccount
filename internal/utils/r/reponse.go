package r

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	FLag bool   `json:"flag"`
	Data any    `json:"data"`
	Code int    `json:"code"`
	Mes  string `json:"mes"`
}

func JSON(ctx *gin.Context, r *Response) {
	ctx.JSON(r.Code, r)
}

func Error(mes string) *Response {
	return &Response{FLag: false, Code: 999, Mes: mes}
}

func OK(mes string) *Response {
	return &Response{FLag: true, Code: 200, Mes: mes}
}

func (r *Response) SetMes(mes string) *Response {
	r.Mes = mes
	return r
}

func (r *Response) SetData(data any) *Response {
	r.Data = data
	return r
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}
