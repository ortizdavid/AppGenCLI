package pythonsamples

type AuthTemplate struct {}

var layout *Layout

func (auth *AuthTemplate) LoginTemplate(appName string) string {
return``+layout.Header(appName)+`
	<div class="row">
		<br><h2 class="text-center">Task Management App</h2>
		<div class="col-md-4 offset-md-4">
			<div class="login-form bg-light mt-4 p-4">
				<form action="{{ url_for('login') }}" method="POST" class="row g-3">
					<h4>Login</h4>
					{{ error_msg }}
					<div class="col-12">
						<label><strong>User Name:</strong> </label>
						<input type="text" name="user_name" class="form-control" placeholder="User Name">
					</div>
					<div class="col-12">
						<label><strong>Password: </strong> </label>
						<input type="password" name="password" class="form-control" placeholder="Password">
					</div>
					<div class="col-12">
						<button type="submit" class="btn btn-primary float-end">Login</button>
					</div>
				</form>
				<hr class="mt-4">
				<div class="col-12">
					<p class="text-center mb-0">Create account <a href="/register">Here</a></p>
				</div>
			</div>
		</div>
	</div>
`+layout.Footer(appName)+`` 
}


func (auth *AuthTemplate) HomeTemplate(appName string) string {
return`{% extends "layouts/back-layout.html" %}

{% block title %} Home {% endblock %}

{% block content %}

    <h2>Home Page</h2><br>

    <div class="alert alert-info">
        <h5>User: {{ logged_user.user_name}}</h5>
        <h5>Role: {{ logged_user.role_name }} </h5>
    </div>

{% endblock %}` 
}