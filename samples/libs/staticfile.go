package libs

type StaticFile struct {}


func (st *StaticFile) CssContent() string {
return `/* Write your CSS code.*/
body {
	font-size: 14px;
	background-color: rgb(250, 250, 250);
}`
}


func (st *StaticFile) JsContent() string {
return `// Write your JavaScript code.
function Test() {
	alert("Its Working");
}`
}