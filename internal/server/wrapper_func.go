package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/NotFound1911/mserver"
	"io"
	"net/http"
)

func WrapBody[Req any](bizFn func(ctx *mserver.Context, req Req) (Result, error)) mserver.HandleFunc {
	return func(ctx *mserver.Context) error {
		var req Req
		r := ctx.GetRequest()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			ctx.SetStatus(http.StatusInternalServerError).Json("请求错误")
			return nil
		}
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		ctx.SetRequest(r)
		if err := json.Unmarshal(body, &req); err != nil {
			ctx.SetStatus(http.StatusInternalServerError).Json("请求内容错误")
			return nil
		}
		res, err := bizFn(ctx, req)
		if err != nil {
			fmt.Printf("执行失败:%v", err)
		}
		ctx.SetStatus(http.StatusOK).Json(res)
		return nil
	}
}
