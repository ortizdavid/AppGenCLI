package pythonsamples

type Helper struct {}


func (h *Helper) HttpCodes() string {
return `HTTP_CODE_OK = 200
HTTP_CODE_CREATED = 201
HTTP_CODE_ACEPTED = 202
HTTP_CODE_NO_CONTENT = 204
HTTP_CODE_MOVED_PERMANENTLY = 301
HTTP_CODE_FOUND = 302
HTTP_CODE_NOT_MODIFIED = 304
HTTP_CODE_BAD_REQUEST = 400
HTTP_CODE_UNAUTHORIZED = 401
HTTP_CODE_FORBIDEN = 403
HTTP_CODE_NOT_FOUND = 404
HTTP_CODE_METHOD_NOT_ALLOWED =  405
HTTP_CODE_REQUEST_ID_TO_LONG =  414
HTTP_CODE_INTERNAL_ERROR = 500
HTTP_CODE_NOT_IMPLEMENTED = 501
HTTP_CODE_BAD_GATEWAY = 502`
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


func (h *Helper) FileUploader() string {
return `import os
import uuid
from flask import request

class FileUploader:

    IMG_EXTENSIONS = {'.png', '.jpg', '.jpeg'}
    VIDEO_EXTENSIONS = {'.mp4', '.mkv', '.flv'}
    PDF_EXTENSIONS = {'.pdf'}
    errors = []
    

    def __init__(self):
        pass


    def get_file_extension(self, filename: str):
        return os.path.splitext(filename)[1]


    def upload_image(self, file_field_name: str, upload_dir: str):

        file = request.files[file_field_name]
        original_name = file.filename
        extension = self.get_file_extension(original_name)
        
        if file_field_name not in request.files:
            self.errors.append('No file part in the request')
            return None
        if original_name == '':
            self.errors.append('No file selected for uploading')
            return None
        if extension not in self.IMG_EXTENSIONS:
            self.errors.append(f'Invalid Image.. Allowed Formats {self.IMG_EXTENSIONS}')
            return None
        if len(self.errors) == 0:
            final_name = str(uuid.uuid4()) + extension
            file.save(os.path.join(upload_dir, final_name))
            return final_name
        
    def upload_video(self, file_field_name: str, upload_dir: str):

        file = request.files[file_field_name]
        original_name = file.filename
        extension = self.get_file_extension(original_name)
        
        if file_field_name not in request.files:
            self.errors.append('No file part in the request')
            return None
        if original_name == '':
            self.errors.append('No file selected for uploading')
            return None
        if extension not in self.VIDEO_EXTENSIONS:
            self.errors.append(f'Invalid Video.. Allowed Formats {self.VIDEO_EXTENSIONS}')
            return None
        if len(self.errors) == 0:
            final_name = str(uuid.uuid4()) + extension
            file.save(os.path.join(upload_dir, final_name))
            return final_name
        

    def upload_pdf(self, file_field_name: str, upload_dir: str):

        file = request.files[file_field_name]
        original_name = file.filename
        extension = self.get_file_extension(original_name)
        
        if file_field_name not in request.files:
            self.errors.append('No file part in the request')
            return None
        if original_name == '':
            self.errors.append('No file selected for uploading')
            return None
        if extension not in self.PDF_EXTENSIONS:
            self.errors.append(f'Invalid PDF.. Allowed formats {self.PDF_EXTENSIONS}')
            return None
        if len(self.errors) == 0:
            final_name = str(uuid.uuid4()) + extension
            file.save(os.path.join(upload_dir, final_name))
            return final_name`
}


func (h *Helper) PasswordHandler() string {
return `from werkzeug.security import generate_password_hash, check_password_hash

class PasswordHandler:

    def generate(password):
        return generate_password_hash(password)

    def check(encrypted_password, pure_password):
        if check_password_hash(encrypted_password, pure_password):
            return True
        else:
            return False`
}
