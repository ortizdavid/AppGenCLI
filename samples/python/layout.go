package pythonsamples

type Layout struct {}


func (l *Layout) AdminMenu() string  {
return `
<nav class="navbar navbar-expand-lg navbar-light bg-primary">
    <div class="container-fluid">
        <a class="navbar-brand" href="/" style="color: white">
            Admin
        </a>
        <button class="navbar-toggler" type="button">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarScroll">
            <ul class="navbar-nav me-auto my-2 my-lg-0 navbar-nav-scroll" style="--bs-scroll-height: 100px;">
                <li class="nav-item">
                    <a href="/home"  class="nav-link active" aria-current="page"style="color: white">
                        Home  |
                    </a>
                </li>
                <li class="nav-item">
                    <a href="/users" class="nav-link" style="color: white">
                        Users |
                    </a>
                </li>
                <li class="nav-item">
                    <a href="/tasks" class="nav-link" style="color: white">
                        Tasks |
                    </a>
                </li>
                <li class="nav-item">
                    <a href="/logout" class="nav-link" style="color: white">
                        Logout
                    </a>
                </li>
            </ul>
        </div>
    </div>
</nav>
`
}


func (l *Layout) NormalMenu() string  {
return `
<nav class="navbar navbar-expand-lg navbar-light bg-primary">
    <div class="container-fluid">
        <a class="navbar-brand" href="/" style="color: white">
            Normal
        </a>
        <button class="navbar-toggler" type="button">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarScroll">
            <ul class="navbar-nav me-auto my-2 my-lg-0 navbar-nav-scroll" style="--bs-scroll-height: 100px;">
                <li class="nav-item">
                    <a href="/home"  class="nav-link active" aria-current="page"style="color: white">
                        Home |
                    </a>
                </li>
                <li class="nav-item">
                    <a href="/user-data" class="nav-link" style="color: white">
                        My Data |
                    </a>
                </li>
                <li class="nav-item">
                    <a href="/tasks" class="nav-link" style="color: white">
                        My Tasks |
                    </a>
                </li>
                <li class="nav-item">
                    <a href="/logout" class="nav-link" style="color: white">
                        Logout
                    </a>
                </li>
            </ul>
        </div>
    </div>
</nav>
`
}


func (l *Layout) FontLayout() string  {
return `

`
}
    


func (l *Layout) BackLayout() string  {
return `

`
}

