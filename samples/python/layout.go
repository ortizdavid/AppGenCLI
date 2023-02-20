package pythonsamples

type Layout struct {}


func (l *Layout) Header(appName string)  string {
return `<!DOCTYPE html>
<html lang="en">
<head>
    {% block head %}
        <link href="{{url_for('static', filename = 'css/style.css')}}" rel="stylesheet">
        <link href="{{url_for('static', filename = 'lib/bootstrap/css/bootstrap.min.css')}}" rel="stylesheet">
        <title>`+appName+` - {% block title %}{% endblock %}</title>
    {% endblock %}        
</head>
<body style="background-color: hsl(233, 44%, 96%);">`
}

func (l *Layout) Footer(appName string)  string {
return `<div id="footer" class="footer">
        {% block footer %}
            <div align="center">
               `+appName+` &copy; Copyright 2022 
            </div>
        {% endblock %}
    </div>
    <script src="{{url_for('static', filename = 'js/script.js')}}"></script>
    <script src="{{url_for('static', filename = 'js/jquery/jquery.min.js')}}"></script>
    <script src="{{url_for('static', filename = 'lib/bootstrap/js/bootstrap.min.js')}}"></script>
</body>
</html>`
}


func (l *Layout) BlockContent() string {
return `<div id="content" class="container">
        <br><br><br>
        {% block content %}
        
        {% endblock %} 
    </div>`
}


func (l *Layout) AdminMenu() string  {
return `<nav class="navbar navbar-expand-lg navbar-light bg-primary">
<div class="container-fluid">
    <a class="navbar-brand" href="/home" style="color: white">
        Admin
    </a>
    <button class="navbar-toggler" type="button">
        <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarScroll">
        <ul class="navbar-nav me-auto my-2 my-lg-0 navbar-nav-scroll" style="--bs-scroll-height: 100px;">
            <li class="nav-item">
                <a href="/home"  class="nav-link active" aria-current="page"style="color: white">
                    Home  
                </a>
            </li>
            <li class="nav-item">
                <a href="/users" class="nav-link" style="color: white">
                    Users 
                </a>
            </li>
            <li class="nav-item">
                <a href="/tasks" class="nav-link" style="color: white">
                    Tasks 
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
</nav>`
}


func (l *Layout) NormalMenu() string  {
return `<nav class="navbar navbar-expand-lg navbar-light bg-primary">
    <div class="container-fluid">
        <a class="navbar-brand" href="/home" style="color: white">
            Normal
        </a>
        <button class="navbar-toggler" type="button">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarScroll">
            <ul class="navbar-nav me-auto my-2 my-lg-0 navbar-nav-scroll" style="--bs-scroll-height: 100px;">
                <li class="nav-item">
                    <a href="/home"  class="nav-link active" aria-current="page"style="color: white">
                        Home 
                    </a>
                </li>
                <li class="nav-item">
                    <a href="/user-data" class="nav-link" style="color: white">
                        My Data 
                    </a>
                </li>
                <li class="nav-item">
                    <a href="/my-tasks" class="nav-link" style="color: white">
                        My Tasks 
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
</nav>`
}


func (l *Layout) FontLayout(appName string) string  {
return ``+l.Header(appName)+`
    <nav class="navbar navbar-expand-lg navbar-light bg-primary">
        <div class="container-fluid">
            <a class="navbar-brand" href="/" style="color: white">
                Flask MVC
            </a>
          <button class="navbar-toggler" type="button">
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse" id="navbarScroll">
            <ul class="navbar-nav me-auto my-2 my-lg-0 navbar-nav-scroll" style="--bs-scroll-height: 100px;">
              <li class="nav-item">
                <a href="/"  class="nav-link active" aria-current="page"style="color: white">
                    Home
                </a>
              </li>
              <li class="nav-item">
                <a href="/register" class="nav-link" style="color: white">
                    Register
                </a>
              </li>
              <li class="nav-item">
                <a href="/login" class="nav-link" style="color: white">
                    Login
                </a>
              </li>
            </ul>
          </div>
        </div>
    </nav>
    `+l.BlockContent()+`
    `+l.Footer(appName)+``
}


func (l *Layout) BackLayout(appName string) string  {
return ``+l.Header(appName)+`
    <!-- MENU ACCORDING USER ROLE -->
    {% if logged_user.role_name == 'administrator' %}
        {% include 'layouts/admin-menu.html' %}
    {% elif logged_user.role_name == 'normal' %}
        {% include 'layouts/normal-menu.html' %}
    {% else %}
        {% include 'error/404-menu.html' %}
    {% endif %}
    `+l.BlockContent()+`
`+l.Footer(appName)+``
}

