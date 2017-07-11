package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		var cookie []*http.Cookie = make([]*http.Cookie, 3)
		for i := 0; i < 3; i++ {
			cookie[i] = new(http.Cookie)
			cookie[i].Expires = time.Now().Add(24 * time.Hour)
		}
		cookie[0].Name = "username"
		cookie[0].Value = "jon"
		cookie[1].Name = "sex"
		cookie[1].Value = "male"
		cookie[2].Name = "old"
		cookie[2].Value = "22"
		for i := 0; i < 3; i++ {
			c.SetCookie(cookie[i])
		}
		return c.String(http.StatusOK, "Hello World")
	}
}
