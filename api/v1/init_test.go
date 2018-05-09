package v1_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/adams-sarah/test2doc/test"
	"github.com/labstack/echo"
	"github.com/lapis-zero09/tada-server/api/router"
)

var server *test.Server

func TestMain(m *testing.M) {
	var err error
	r := router.NewRouter()
	// リクエストのURLに含まれるパラメータ関連の情報を教えてあげる
	// Request -> /api/v1/users/1
	// Test focus -> /api/v1/users/:user_id
	// return of MakeURLExtractor -> map[user_id:1]
	// into doc.apib
	// endpoint = /api/v1/users/{user_id}
	// example user_id -> 1
	test.RegisterURLVarExtractor(makeURLVarExtractor(r))
	server, err = test.NewServer(r)
	if err != nil {
		log.Fatal(err.Error())
	}

	code := m.Run()
	// make doc.apib
	server.Finish()
	os.Exit(code)
}

func makeURLVarExtractor(e *echo.Echo) func(req *http.Request) map[string]string {
	return func(req *http.Request) map[string]string {
		ctx := e.AcquireContext()
		defer e.ReleaseContext(ctx)
		pnames := ctx.ParamNames()
		if len(pnames) == 0 {
			return nil
		}

		paramsMap := make(map[string]string, len(pnames))
		for _, name := range pnames {
			paramsMap[name] = ctx.Param(name)
		}
		return paramsMap
	}
}
