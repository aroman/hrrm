package main 

import (
    "bytes"
    "io"
    "github.com/hoisie/web"
    "github.com/hoisie/mustache"
)

func index(ctx *web.Context) {
    var buf bytes.Buffer
    data := map[string]string{}
    data["name"] = "flerp"
    buf.WriteString(mustache.RenderFile("templates/index.html", data))
    io.Copy(ctx, &buf)
}

func main () {
    web.Get("/", index)
    web.Run(":8080")
}