package pythonsamples

type StaticFile struct {}


func (st *StaticFile) CssContent() string {
return `/* Write your CSS code.*/
body {
	font-size: 14px;
	background-color: rgb(229, 230, 231);
}`
}


func (st *StaticFile) JsContent() string {
return `// Write your JavaScript code.
function Test() {
	alert("Its Working");
}`
}