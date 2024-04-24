// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "tiny beer",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/memo/create": {
            "post": {
                "description": "使用access_token获取视频列表资源",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "备忘录相关接口"
                ],
                "summary": "根据内容创建新的备忘录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "访问令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "备忘录创建参数",
                        "name": "{object}",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.MemoCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SimpleResponse"
                        }
                    }
                }
            }
        },
        "/memo/delete": {
            "delete": {
                "description": "使用access_token获取视频列表资源",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "备忘录相关接口"
                ],
                "summary": "根据Id删除备忘录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "访问令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "备忘录删除参数",
                        "name": "{object}",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.MemoDeleteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SimpleResponse"
                        }
                    }
                }
            }
        },
        "/memo/list": {
            "get": {
                "description": "使用access_token获取视频列表资源",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "备忘录相关接口"
                ],
                "summary": "获取备忘录列表资源",
                "parameters": [
                    {
                        "type": "string",
                        "description": "访问令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.MemoListResponse"
                        }
                    }
                }
            }
        },
        "/memo/update": {
            "put": {
                "description": "使用access_token获取视频列表资源",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "备忘录相关接口"
                ],
                "summary": "根据Id更新备忘录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "访问令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "备忘录删除参数",
                        "name": "{object}",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.MemoUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SimpleResponse"
                        }
                    }
                }
            }
        },
        "/user/auth": {
            "post": {
                "description": "使用access_token验证用户身份",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "验证用户身份",
                "parameters": [
                    {
                        "type": "string",
                        "description": "访问令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "完成用户登陆操作分发access_token和refresh_token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用于用户登录",
                "parameters": [
                    {
                        "description": "登录参数",
                        "name": "{object}",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.LoginResponse"
                        }
                    }
                }
            }
        },
        "/user/refresh": {
            "post": {
                "description": "使用refresh_token更新访问令牌",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "更新访问令牌",
                "parameters": [
                    {
                        "type": "string",
                        "description": "refresh_token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/video/list": {
            "get": {
                "description": "使用access_token获取视频列表资源",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "视频相关接口"
                ],
                "summary": "获取视频列表资源",
                "parameters": [
                    {
                        "type": "string",
                        "description": "访问令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.VideoListResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 3
                },
                "username": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 4
                }
            }
        },
        "model.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.MemoCreateRequest": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 2
                }
            }
        },
        "model.MemoDeleteRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.MemoItem": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.MemoListResponse": {
            "type": "object",
            "properties": {
                "memoes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.MemoItem"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.MemoUpdateRequest": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.SimpleResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.VideoItem": {
            "type": "object",
            "properties": {
                "image": {
                    "type": "string"
                },
                "intro": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.VideoListResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "videoes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.VideoItem"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:9999",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "个人主页",
	Description:      "用于个人文件管理",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}