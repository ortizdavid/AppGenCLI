package pythonsamples


type PageError struct {
	
}


func (perr *PageError) Error404() string  {
return `<p>Page Not Found</p>
`
}



func (perr *PageError) Error404Menu() string  {
return `<p>Menu Not Found</p>
`
}


