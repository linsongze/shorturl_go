package routers

import (
	"fmt"
	"github.com/linsongze/shorturl_go/service"
	"github.com/linsongze/shorturl_go/utils"
	"github.com/valyala/fasthttp"
)

var storeService service.StoreService = service.NewRamStore()


func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "curl http://host:port/to/short?url=http%3a%2f%2flocalhost%2fxxx%3fa%3da%26b%3db")
}
func Short(ctx *fasthttp.RequestCtx){
	if ctx.QueryArgs().Peek("url") == nil {
		fmt.Fprintf(ctx, "")
		return
	}
	url := string(ctx.QueryArgs().Peek("url"))
	if !utils.CheckUrl(url) {
		fmt.Fprintf(ctx, "")
		return
	}
	id := storeService.IncAndGet()
	shortCode := utils.Ten_To_62(id)
	storeService.Save(shortCode,url)
	fmt.Fprintf(ctx, shortCode)
}
func redirect(ctx *fasthttp.RequestCtx){
	 var sc string =  ctx.UserValue("sc").(string)
	 url := storeService.Get(sc)
	if url == "" {
		ctx.SetStatusCode(404)
		return
	}
	ctx.Redirect(url,302)
}