package collectionsamples

type Postman struct {}

func (p *Postman) PostmanCollection(appName string) string {
return `{
	"info": {
		"_postman_id": "d5b15e98-ee83-4ecd-94b2-6a330eb421e9",
		"name": "`+appName+`",
		"description": "# Task Management App\n\n## Description\n\nThis API Allows user manage his own daily tasks.\n\n## Features\n\n- Add tasks\n- Add Users\n- Add User roles\n- Upload image\n- Login\n- Logout",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "Tasks",
			"item": [
				{
					"name": "Get All Tasks",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "{{FLASK_API_ROOT}}/tasks",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"tasks"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Task",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "{{FLASK_API_ROOT}}/tasks/2",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"tasks",
								"2"
							]
						}
					},
					"response": []
				},
				{
					"name": "Filter Task by Date",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "{{FLASK_API_ROOT}}/tasks/2023-01-01/2023-02-28",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"tasks",
								"2023-01-01",
								"2023-02-28"
							]
						}
					},
					"response": []
				},
				{
					"name": "Filter Task by User",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "{{FLASK_API_ROOT}}/tasks/user/2",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"tasks",
								"user",
								"2"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete Task",
					"request": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "{{FLASK_API_ROOT}}/tasks/10",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"tasks",
								"10"
							]
						}
					},
					"response": []
				},
				{
					"name": "Add Task",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"task_name\": \"Cooking Bean\",\r\n    \"user_id\": 1,\r\n    \"start_date\": \"2023-01-23\",\r\n    \"end_date\": \"2021-11-18\",\r\n    \"description\": \"Cooking at 23h O'clock\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{FLASK_API_ROOT}}/tasks",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"tasks"
							]
						}
					},
					"response": []
				},
				{
					"name": "Edit Task",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"task_name\": \"Cooking Fish\",\r\n    \"user_id\": 1,\r\n    \"start_date\": \"2023-03-01\",\r\n    \"end_date\": \"2021-01-12\",\r\n    \"description\": \"Cooking at 13:00\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{FLASK_API_ROOT}}/tasks/11",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"tasks",
								"11"
							]
						}
					},
					"response": []
				}
			],
			"description": "# Tasks Made by each user\n\n- Get All user tasks\n- Get a task by id\n- Delete a task\n- Filter Task by Date\n- Filter Task by User Id\n- Add a new Task\n- Edit an existing Task"
		},
		{
			"name": "Users",
			"item": [
				{
					"name": "Get All Users",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "{{FLASK_API_ROOT}}/users",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"users"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get User",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "{{FLASK_API_ROOT}}/users/7",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"users",
								"7"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete User",
					"request": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "{{FLASK_API_ROOT}}/users/6",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"users",
								"6"
							]
						}
					},
					"response": []
				},
				{
					"name": "Add User",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"user_name\": \"user11\",\r\n    \"password\": \"12345678\",\r\n    \"role_id\": 2,\r\n    \"image\": \"\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{FLASK_API_ROOT}}/users",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"users"
							]
						}
					},
					"response": []
				},
				{
					"name": "Upload Image",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "formdata",
							"formdata": [
								{
									"key": "image",
									"type": "file",
									"src": "/C:/Users/Ortiz/Downloads/WhatsApp Image 2023-02-13 at 10.49.44.jpeg"
								}
							],
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{FLASK_API_ROOT}}/users/4/upload",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"users",
								"4",
								"upload"
							]
						}
					},
					"response": []
				},
				{
					"name": "Edit User",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"user_name\": \"user06\",\r\n    \"password\": \"12345678\",\r\n    \"role_id\": 2,\r\n    \"image\": \"a.jpg\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{FLASK_API_ROOT}}/users/7",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"users",
								"7"
							]
						}
					},
					"response": []
				}
			],
			"description": "## Users that make a task\n\n- Get All Registered Users\n- Get an User By ID\n- Delete an User\n- Add an User\n- Update an existing User\n- Upload User Image"
		},
		{
			"name": "Roles",
			"item": [
				{
					"name": "Get Role",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"task_name\": \"Nothing\",\r\n    \"user_id\": 2,\r\n    \"start_date\": \"2028-09-11\",\r\n    \"end_date\": \"2029-09-11\",\r\n    \"description\": \"None\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{FLASK_API_ROOT}}/roles/1",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"roles",
								"1"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete  Role",
					"request": {
						"method": "DELETE",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"task_name\": \"Nothing\",\r\n    \"user_id\": 2,\r\n    \"start_date\": \"2028-09-11\",\r\n    \"end_date\": \"2029-09-11\",\r\n    \"description\": \"None\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{FLASK_API_ROOT}}/roles/3",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"roles",
								"3"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get All Roles",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"task_name\": \"Nothing\",\r\n    \"user_id\": 2,\r\n    \"start_date\": \"2028-09-11\",\r\n    \"end_date\": \"2029-09-11\",\r\n    \"description\": \"None\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{FLASK_API_ROOT}}/roles",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"roles"
							]
						}
					},
					"response": []
				},
				{
					"name": "Add Role",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"role_name\": \"Employee\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{FLASK_API_ROOT}}/roles",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"roles"
							]
						}
					},
					"response": []
				},
				{
					"name": "Edit Role",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"role_name\": \"employee\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{FLASK_API_ROOT}}/roles/3",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"roles",
								"3"
							]
						}
					},
					"response": []
				}
			],
			"description": "## User Roles Features\n\n- Add Role\n- Delete a Role\n- Get All Roles\n- Add a New Role\n- Edit an Existing Role"
		},
		{
			"name": "Auth",
			"item": [
				{
					"name": "Logout",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "{{FLASK_API_ROOT}}/auth/logout",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"auth",
								"logout"
							]
						}
					},
					"response": []
				},
				{
					"name": "Login",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"user_name\": \"admin01\",\r\n    \"password\": \"12345678\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{FLASK_API_ROOT}}/auth/login",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"auth",
								"login"
							]
						}
					},
					"response": []
				},
				{
					"name": "Logged User",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "{{FLASK_API_ROOT}}/auth/user",
							"host": [
								"{{FLASK_API_ROOT}}"
							],
							"path": [
								"auth",
								"user"
							]
						}
					},
					"response": []
				}
			],
			"description": "## Autentication Features\n\n- Login\n- Logout\n- Logged User data"
		},
		{
			"name": "Root",
			"item": [
				{
					"name": "API Root 1",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:5000",
							"host": [
								"localhost"
							],
							"port": "5000"
						}
					},
					"response": []
				},
				{
					"name": "API Root 2",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "{{FLASK_API_ROOT}}",
							"host": [
								"{{FLASK_API_ROOT}}"
							]
						}
					},
					"response": []
				},
				{
					"name": "Download Postman  Collection",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:5000/postman-collection",
							"host": [
								"localhost"
							],
							"port": "5000",
							"path": [
								"postman-collection"
							]
						}
					},
					"response": []
				},
				{
					"name": "Download Postman  Test Run",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:5000/postman-test",
							"host": [
								"localhost"
							],
							"port": "5000",
							"path": [
								"postman-test"
							]
						}
					},
					"response": []
				}
			],
			"description": "# First Endpoint \n\n- /api\n- /"
		}
	],
	"event": [
		{
			"listen": "prerequest",
			"script": {
				"type": "text/javascript",
				"exec": [
					""
				]
			}
		},
		{
			"listen": "test",
			"script": {
				"type": "text/javascript",
				"exec": [
					""
				]
			}
		}
	],
	"variable": [
		{
			"key": "FLASK_API_ROOT",
			"value": "localhost:5000/api",
			"type": "default"
		}
	]
}
`	
}


func (p *Postman) PostmanTestRun(appName string) string {
return `{
	"id": "cfb533e8-c494-4e06-b384-bb1c12434d7a",
	"name": "`+appName+`",
	"timestamp": "2023-02-21T08:09:55.403Z",
	"collection_id": "0-d5b15e98-ee83-4ecd-94b2-6a330eb421e9",
	"folder_id": 0,
	"environment_id": "048d60e4-cffd-4196-91a6-db5016295e2c",
	"totalPass": 0,
	"totalFail": 0,
	"results": [
		{
			"id": "9fad4c47-f490-4b74-b585-46d547d3920a",
			"name": "Get All Tasks",
			"url": "http://127.0.0.1:5000/api/tasks",
			"time": 352,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				352
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "b19828a1-7a5c-4e9b-8cd3-9954713f1104",
			"name": "Get Task",
			"url": "http://127.0.0.1:5000/api/tasks/2",
			"time": 12,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				12
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "71b7fc38-755a-47ec-9a8d-e90e87839737",
			"name": "Filter Task by Date",
			"url": "http://127.0.0.1:5000/api/tasks/2023-01-01/2023-02-28",
			"time": 16,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				16
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "954d583a-d4b8-48f7-a0ec-5343df6c7791",
			"name": "Filter Task by User",
			"url": "http://127.0.0.1:5000/api/tasks/user/2",
			"time": 10,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				10
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "87bfdd88-fa51-40ae-adc2-8bf9189c98f3",
			"name": "Delete Task",
			"url": "http://127.0.0.1:5000/api/tasks/10",
			"time": 7,
			"responseCode": {
				"code": 404,
				"name": "NOT FOUND"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				7
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "2d6ef6ba-118c-43b8-82fe-4686d3c1ddfb",
			"name": "Add Task",
			"url": "http://127.0.0.1:5000/api/tasks",
			"time": 29,
			"responseCode": {
				"code": 201,
				"name": "CREATED"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				29
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "7f4b0a05-40c5-483f-8634-bd850126c45d",
			"name": "Edit Task",
			"url": "http://127.0.0.1:5000/api/tasks/11",
			"time": 14,
			"responseCode": {
				"code": 201,
				"name": "CREATED"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				14
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "f2e790c1-3916-41f3-9d0c-090892111cf2",
			"name": "Get All Users",
			"url": "http://127.0.0.1:5000/api/users",
			"time": 14,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				14
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "1f23515a-bc13-4458-8aea-9bd5cf9c6a78",
			"name": "Get User",
			"url": "http://127.0.0.1:5000/api/users/7",
			"time": 9,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				9
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "b5c4ef53-bbc3-47c9-9b2a-18354de26b78",
			"name": "Delete User",
			"url": "http://127.0.0.1:5000/api/users/6",
			"time": 7,
			"responseCode": {
				"code": 404,
				"name": "NOT FOUND"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				7
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "9013c9a5-5255-442a-9d13-406ced3fc6fd",
			"name": "Add User",
			"url": "http://127.0.0.1:5000/api/users",
			"time": 8,
			"responseCode": {
				"code": 400,
				"name": "BAD REQUEST"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				8
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "7c904ae2-aa1f-4feb-b729-0183a0ee1298",
			"name": "Upload Image",
			"url": "http://127.0.0.1:5000/api/users/4/upload",
			"time": 7,
			"responseCode": {
				"code": 400,
				"name": "BAD REQUEST"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				7
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "db4ae604-f961-48ec-a191-c53315f12ec7",
			"name": "Edit User",
			"url": "http://127.0.0.1:5000/api/users/7",
			"time": 11,
			"responseCode": {
				"code": 201,
				"name": "CREATED"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				11
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "6f7b4b58-a311-4e68-8b10-683bff1a262f",
			"name": "Get Role",
			"url": "http://127.0.0.1:5000/api/roles/1",
			"time": 11,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				11
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "14c5b8ca-91ea-4ccb-a31d-7da14a307a81",
			"name": "Delete  Role",
			"url": "http://127.0.0.1:5000/api/roles/3",
			"time": 9,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				9
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "e7517354-d0b0-4c2d-97df-fee3246989ca",
			"name": "Get All Roles",
			"url": "http://127.0.0.1:5000/api/roles",
			"time": 7,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				7
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "d167469c-3d13-497f-beb0-cbf69c583d85",
			"name": "Add Role",
			"url": "http://127.0.0.1:5000/api/roles",
			"time": 12,
			"responseCode": {
				"code": 201,
				"name": "CREATED"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				12
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "2b1b22b0-066f-4d8d-bbb6-99c817ebbc48",
			"name": "Edit Role",
			"url": "http://127.0.0.1:5000/api/roles/3",
			"time": 10,
			"responseCode": {
				"code": 201,
				"name": "CREATED"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				10
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "e477d48c-b6a6-4f95-acf9-3c9e6ea351e2",
			"name": "Logout",
			"url": "http://127.0.0.1:5000/api/auth/logout",
			"time": 8,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				8
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "354dbce4-cb35-4d14-b45a-107f852abfa6",
			"name": "Login",
			"url": "http://127.0.0.1:5000/api/auth/login",
			"time": 11,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				11
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "0c89dc6e-9b0d-4438-b9b5-a31e4325ad31",
			"name": "Logged User",
			"url": "http://127.0.0.1:5000/api/auth/user",
			"time": 8,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				8
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "3bc8def3-af2a-495a-9e84-f116f8540e3f",
			"name": "API Root 1",
			"url": "http://localhost:5000/",
			"time": 8,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				8
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "ad22bf06-c171-49ba-9763-b68fae8905ab",
			"name": "API Root 2",
			"url": "http://127.0.0.1:5000/api",
			"time": 8,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				8
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "57e7bd10-1751-423f-a08a-a9c84e5bc6ea",
			"name": "Download Postman  Collection",
			"url": "http://localhost:5000/postman-collection",
			"time": 56,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				56
			],
			"allTests": [
				{}
			]
		},
		{
			"id": "90078d73-5463-44a9-b65a-8a1fe33a2739",
			"name": "Download Postman  Test Run",
			"url": "http://localhost:5000/postman-test",
			"time": 16,
			"responseCode": {
				"code": 200,
				"name": "OK"
			},
			"tests": {},
			"testPassFailCounts": {},
			"times": [
				16
			],
			"allTests": [
				{}
			]
		}
	],
	"count": 1,
	"totalTime": 660,
	"collection": {
		"requests": [
			{
				"id": "9fad4c47-f490-4b74-b585-46d547d3920a",
				"method": "GET"
			},
			{
				"id": "b19828a1-7a5c-4e9b-8cd3-9954713f1104",
				"method": "GET"
			},
			{
				"id": "71b7fc38-755a-47ec-9a8d-e90e87839737",
				"method": "GET"
			},
			{
				"id": "954d583a-d4b8-48f7-a0ec-5343df6c7791",
				"method": "GET"
			},
			{
				"id": "87bfdd88-fa51-40ae-adc2-8bf9189c98f3",
				"method": "DELETE"
			},
			{
				"id": "2d6ef6ba-118c-43b8-82fe-4686d3c1ddfb",
				"method": "POST"
			},
			{
				"id": "7f4b0a05-40c5-483f-8634-bd850126c45d",
				"method": "PUT"
			},
			{
				"id": "f2e790c1-3916-41f3-9d0c-090892111cf2",
				"method": "GET"
			},
			{
				"id": "1f23515a-bc13-4458-8aea-9bd5cf9c6a78",
				"method": "GET"
			},
			{
				"id": "b5c4ef53-bbc3-47c9-9b2a-18354de26b78",
				"method": "DELETE"
			},
			{
				"id": "9013c9a5-5255-442a-9d13-406ced3fc6fd",
				"method": "POST"
			},
			{
				"id": "7c904ae2-aa1f-4feb-b729-0183a0ee1298",
				"method": "POST"
			},
			{
				"id": "db4ae604-f961-48ec-a191-c53315f12ec7",
				"method": "PUT"
			},
			{
				"id": "6f7b4b58-a311-4e68-8b10-683bff1a262f",
				"method": "GET"
			},
			{
				"id": "14c5b8ca-91ea-4ccb-a31d-7da14a307a81",
				"method": "DELETE"
			},
			{
				"id": "e7517354-d0b0-4c2d-97df-fee3246989ca",
				"method": "GET"
			},
			{
				"id": "d167469c-3d13-497f-beb0-cbf69c583d85",
				"method": "POST"
			},
			{
				"id": "2b1b22b0-066f-4d8d-bbb6-99c817ebbc48",
				"method": "PUT"
			},
			{
				"id": "e477d48c-b6a6-4f95-acf9-3c9e6ea351e2",
				"method": "GET"
			},
			{
				"id": "354dbce4-cb35-4d14-b45a-107f852abfa6",
				"method": "POST"
			},
			{
				"id": "0c89dc6e-9b0d-4438-b9b5-a31e4325ad31",
				"method": "GET"
			},
			{
				"id": "3bc8def3-af2a-495a-9e84-f116f8540e3f",
				"method": "GET"
			},
			{
				"id": "ad22bf06-c171-49ba-9763-b68fae8905ab",
				"method": "GET"
			},
			{
				"id": "57e7bd10-1751-423f-a08a-a9c84e5bc6ea",
				"method": "GET"
			},
			{
				"id": "90078d73-5463-44a9-b65a-8a1fe33a2739",
				"method": "GET"
			}
		]
	}
}
	
`	
}