package pythonsamples

type Helper struct {}

func (h *Helper) HttpCodes() string {
return `HTTP_CODE_OK = 200
HTTP_CODE_CREATED = 201
HTTP_CODE_BAD_REQUEST = 400
HTTP_CODE_NOT_FOUND = 404
HTTP_CODE_INTERNAL_ERROR = 500`
}

func (h *Helper) Constants() string {
return `ADMIN_USER = 1
NORMAL_USER = 2

TASK_STATUS = [
	'Completed', 
	'In Progress', 
	'Pending'
]`
}