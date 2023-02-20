package pythonsamples

type TaskTemplate struct {}


func (task *TaskTemplate) AddTemplate() string {
return`{% extends "layouts/back-layout.html" %}

{% block title %} Add Task {% endblock %}

{% block content %}

    <h2>Add Task</h2>

    <div class="card ">
        <div class="row">
            <form action="/tasks/add" method="POST" enctype="multipart/form-data" class="">
                <div class="card">
                    <div class="card-header">
                        Data
                    </div>
                    <div class="card-body">
                        <div class="col-xs-12 col-sm-12 col md-12">
                            <div class="form-group" >
                                <strong>Task Name:</strong>
                                <input type="text" name="task_name" id="" class="form-control" placeholder="Type Task Name" autofocus>
                            </div>
                        </div>
                    </div>
                    <div class="card-body">
                        <div class="col-xs-12 col-sm-12 col md-12">
                            <div class="form-group" style="width: 30%;">
                                <strong>Start Date:</strong>
                                <input type="date" name="start_date" id="" class="form-control">
                            </div>
                        </div>
                    </div>
                    <div class="card-body">
                        <div class="col-xs-12 col-sm-12 col md-12">
                            <div class="form-group" style="width: 30%;">
                                <strong>End Date:</strong>
                                <input type="date" name="end_date" id="" class="form-control">
                            </div>
                        </div>
                    </div>
                    <div class="card-body">
                        <div class="col-xs-12 col-sm-12 col md-12">
                            <div class="form-group" style="">
                                <strong>Description:</strong>
                                <textarea name="description" id="" cols="30" rows="5" placeholder="Type description" class="form-control"></textarea>
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
                        <a href="/my-tasks" class="btn btn-danger">
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


func (task *TaskTemplate) EditTemplate() string {
return`

` 
}


func (task *TaskTemplate) ShowTemplate() string {
return`{% extends "layouts/back-layout.html" %}

{% block title %} Tasks {% endblock %}

{% block content %}

    <h2>All Tasks</h2>

    <div class="card" style="padding: 1%;">
        {% if num_registos == 0 %}
        <div class="alert alert-warning">
            <h5>No records found!</h5>
        </div>
        {% else %}
            <table class="table table-hover">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Task</th>
                        <th>User</th>
                        <th>Start Date</th>
                        <th>End Date</th>
                        <th>Status</th>
                        <th>Opções</th>
                    </tr>
                </thead>
                <tbody>
                    {% for task in  tasks %}
                        <tr>
                            <td>{{ task.task_id }}</td>
                            <td>{{ task.user_name }}</td>
                            <td>{{ task.task_name }}</td>
                            <td>{{ task.start_date }}</td>
                            <td>{{ task.end_date }}</td>
                            <td>{{ task.status }}</td>
                            <td>
                                <a href="/tasks/{{ task.task_id }}/details" class="btn btn-primary">
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



func (task *TaskTemplate) SearchTemplate() string {
return`

` 
}




func (task *TaskTemplate) SearchResultsTemplate() string {
return`

` 
}


func (task *TaskTemplate) DetailsTemplate() string {
return`{% extends "layouts/back-layout.html" %}

{% block title %} Details {% endblock %}

{% block content %}

    <h3>Task Details</h3>

    <div align="right">
        <a href="/my-tasks" class="btn btn-primary">Go Back</a>
    </div>
    
    <div class="card" style="padding: 1%;">
        <div class="card-header">
            <h5>Task Data</h5>
        </div>
        <div class="card-body">
            <br><br>
            <p><strong>Id: </strong>{{ task.user_id }}</p>
            <p><strong>Task Name: </strong>{{ task.task_name }}</p>
            <p><strong>Start Date: </strong>{{ task.start_date }}</p>
            <p><strong>End Date: </strong>{{ task.end_date }}</p>
            <p><strong>User: </strong>{{ task.user_name }}</p>
            <p><strong>Role: </strong>{{ task.role_name }}</p>
            <p><strong>Created at: </strong>{{ task.created_at }}</p>
            <p><strong>Updated at: </strong>{{ task.updated_at }}</p>
            <p><strong>Status: </strong>{{ task.status }}</p>
            <p><strong>Description: </strong>{{ task.description }}</p>
        </div>
    </div>

{% endblock %}` 
}


func (task *TaskTemplate) UserTasksTemplate() string {
return`{% extends "layouts/back-layout.html" %}

{% block title %} My Tasks {% endblock %}

{% block content %}

    <h2>My Tasks</h2>

    <div align="right">
        <a href="/tasks/add" class="btn btn-primary">Add new</a>
    </div>

    <div class="card" style="padding: 1%;">
        {% if num_registos == 0 %}
        <div class="alert alert-warning">
            <h5>No records found!</h5>
        </div>
        {% else %}
            <table class="table table-hover">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Task</th>
                        <th>Start Date</th>
                        <th>End Date</th>
                        <th>Status</th>
                        <th>Opções</th>
                    </tr>
                </thead>
                <tbody>
                    {% for task in  tasks %}
                        <tr>
                            <td>{{ task.task_id }}</td>
                            <td>{{ task.task_name }}</td>
                            <td>{{ task.start_date }}</td>
                            <td>{{ task.end_date }}</td>
                            <td>{{ task.status }}</td>
                            <td>
                                <a href="/tasks/{{ task.task_id }}/details" class="btn btn-primary">
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

