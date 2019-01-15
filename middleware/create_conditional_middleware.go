package middleware

import "github.com/kataras/iris"

func CreateConditionalMiddleware(
	middleware func(ctx iris.Context),
	condition func(ctx iris.Context) bool,
) func(ctx iris.Context) {

	return func(ctx iris.Context) {
		pass := condition(ctx)

		if pass {
			middleware(ctx)
		} else {
			ctx.Next()
		}
	}
}
