package htmlhelper

import (
	"fmt"

	"dora.com/fullstack/framework/allocator"
)

func HTMLGenerator(headers string, bodyCom string) []byte {
	return allocator.ConcatCopyPreAllocate([][]byte{
		[]byte("<html>"),
		[]byte(headers),
		[]byte("<body>"),
		[]byte(bodyCom),
		[]byte("<body>"),
		[]byte("</html>")})
}

//Div is div
func Div(id string, name string, style string, classes string, childs string) string {
	return fmt.Sprint(
		"<div id=\"", id, "\" name=\"", name, "\" class=\"", classes, "\" style=\"", style, "\" >",
		childs,
		"</div>")
}

//A is link
func A(id string, name string, href string, style string, classes string, text string) string {
	return fmt.Sprint(
		"<a id=\"", id, "\" name=\"", name, "\" href=\"", href, "\", class=\"", classes, "\" style=\"", style, "\" >",
		text,
		"</a>")
}

//Img is img
func Img(id string, name string, src string, style string, classes string) string {
	return fmt.Sprint(
		"<img id=\"", id, "\" name=\"", name, "\" src=\"", src, "\", class=\"",
		classes, "\" style=\"", style, "\" />")
}
