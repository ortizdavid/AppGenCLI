package pythonsamples

type UserTemplate struct {}


func (user *UserTemplate) AddTemplate() string {
return`{% extends "layouts/back-layout.html" %}

{% block title %} Add User {% endblock %}

{% block content %}

    <h2>Add Users</h2>

    <div class="card ">
        <div class="row">
            <form action="/users/add" method="POST" enctype="multipart/form-data" class="">
                <div class="card">
                    <div class="card-header">
                        Data
                    </div>
                    <div class="card-body">
                        <div class="col-xs-12 col-sm-12 col md-12">
                            <div class="form-group" >
                                <strong>User Name:</strong>
                                <input type="text" name="user_name" id="" class="form-control" placeholder="Type User Name" autofocus>
                            </div>
                        </div>
                    </div>
                    <div class="card-body">
                        <div class="col-xs-12 col-sm-12 col md-12">
                            <div class="form-group" style="width: 30%">
                                <strong>Role:</strong>
                                <select name="role_id" class="form-control" id="">
                                    <option value="">Select Role</option>
                                    {% for role in roles %}
                                        <option value="{{ role.role_id }}">{{ role.role_name }}</option>
                                    {% endfor %}
                                </select>
                            </div>
                        </div>
                    </div>
                    <div class="card-body">
                        <div class="col-xs-12 col-sm-12 col md-12">
                            <div class="form-group" style="width: 50%;">
                                <strong>Password:</strong>
                                <input type="password" name="password" value="" placeholder="Type Password" class="form-control">
                            </div>
                        </div>
                    </div>
                </div> 
                <div class="col-xs-12 col-sm-12 col md-12" style="padding-left: 10px;;">
                    <div class="form-group">
                        <br>
                        <button type="submit" class="btn btn-primary">
                            <i class=""></i>
                            Submit
                        </button>
                        <a href="/users" class="btn btn-danger">
                            <i class=""></i>
                            Cancel
                        </a>
                    </div>
                    <br>
                </div>
            </form>
        </div>
    </div>  

{% endblock %}` 
}


func (user *UserTemplate) EditTemplate() string {
return`

` 
}


func (user *UserTemplate) ShowTemplate() string {
return`{% extends "layouts/back-layout.html" %}

{% block title %} Users {% endblock %}

{% block content %}

    <h2>Users</h2>

    <div align="right">
        <a href="/users/add" class="btn btn-primary">Add New</a>
    </div>

    <div class="card" style="padding: 1%;">
        {% if num_registos == 0 %}
        <div class="alert-alert-warning">
             <h5>No records found!</h5>
        </div>
        {% else %}
            <table class="table table-hover">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Username</th>
                        <th>Role</th>
                        <th>Created at</th>
                        <th>Update at</th>
                        <th>Options</th>
                    </tr>
                </thead>
                <tbody>
                    {% for user in users %}
                        <tr>
                            <td>{{ user.user_id }}</td>
                            <td>{{ user.user_name }}</td>
                            <td>{{ user.role_name }}</td>
                            <td>{{ user.created_at }}</td>
                            <td>{{ user.updated_at }}</td>
                            <td>
                                <a href="/users/{{ user.user_id }}/details" class="btn btn-primary">
                                    Details
                                </a>
                            <td>
                        </tr>
                    {% endfor %}
                </tbody>
            </table>
        {% endif %}
    </div>
    
{% endblock %}` 
}



func (user *UserTemplate) SearchTemplate() string {
return`

` 
}


func (user *UserTemplate) SearchResultsTemplate() string {
return`

` 
}


func (user *UserTemplate) DetailsTemplate() string {
return`{% extends "layouts/back-layout.html" %}

{% block title %} Details {% endblock %}

{% block content %}

    <h3>User Details</h3>

    <div align="right">
        <a href="/users" class="btn btn-primary">Go Back</a>
    </div>
    
    <div class="card" style="padding: 1%;">
        <div class="card-header">
            <h5>User Data</h5>
        </div>
        <div class="card-body">
            {% if user.image is none %}
                <img style="width:10%" class="img img-rounded" src="/static/images/blank-user.png" alt="Image" class="card-img-left">
            {% else %}
                <img style="width:30%" class="img img-rounded" src="/static/uploads/imgs/{{ user.image }}" alt="Image" class="card-img-left">
            {% endif %}
            <br><br>
            <p><strong>Id: </strong>{{ user.user_id }}</p>
            <p><strong>Name: </strong>{{ user.user_name }}</p>
            <p><strong>Role: </strong>{{ user.role_name }}</p>
            <p><strong>Created at: </strong>{{ user.created_at }}</p>
            <p><strong>Updated at: </strong>{{ user.updated_at }}</p>
        </div>
    </div>

{% endblock %}` 
}


func (user *UserTemplate) UserDataTemplate() string {
return`
{% extends "layouts/back-layout.html" %}

{% block title %} User Data {% endblock %}

{% block content %}

    <h3>My Personal data</h3>

    <div class="card" style="padding: 1%;">
        <div class="card-header">
            <h5>Data</h5>
        </div>
        <div class="card-body">
            {% if data.image is none %}
                <img style="width:10%" class="img img-rounded" src="/static/images/blank-data.png" alt="Image" class="card-img-left"><br>
                <a href="/upload-image" class="btn btn-primary">Upload Image</a>
            {% else %}
                <img style="width:30%" class="img img-rounded" src="/static/uploads/imgs/{{ data.image }}" alt="Image" class="card-img-left">
            {% endif %}
            <br><br>
            <p><strong>Id: </strong>{{ data.user_id }}</p>
            <p><strong>Name: </strong>{{ data.user_name }}</p>
            <p><strong>Role: </strong>{{ data.role_name }}</p>
            <p><strong>Created at: </strong>{{ data.created_at }}</p>
            <p><strong>Updated at: </strong>{{ data.updated_at }}</p>
        </div>
    </div>

{% endblock %}` 
}

func (user *UserTemplate) UserUploadTemplate() string {
return`
{% extends "layouts/back-layout.html" %}

{% block title %} Upload Image {% endblock %}

{% block content %}

    <h2>Upload My Image</h2>

    <div class="card ">
        <div class="row">
            <form action="/upload-image" method="POST" enctype="multipart/form-data" class="">
                <div class="card">
                    <div class="card-header">
                        Data
                    </div>
                    <div class="card-body">
                        <div class="col-xs-12 col-sm-12 col md-12">
                            <div class="form-group" style="width: 60%;">
                                <strong>Image:</strong>
                                <input type="file" name="image" id="" class="form-control">
                            </div>
                        </div>
                    </div>
                </div> 
                <div class="col-xs-12 col-sm-12 col md-12" style="padding-left: 10px;;">
                    <div class="form-group">
                        <br>
                        <button type="submit" class="btn btn-primary">
                            <i class=""></i>
                            Submit
                        </button>
                        <a href="/user-data" class="btn btn-danger">
                            <i class=""></i>
                            Cancel
                        </a>
                    </div>
                    <br>
                </div>
            </form>
        </div>
    </div>  

{% endblock %}` 
}