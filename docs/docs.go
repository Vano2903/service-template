// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Vano2903",
            "url": "https://github.com/vano2903",
            "email": "davidevanoncini2003@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/all": {
            "get": {
                "description": "Get all user for an unauthorized user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all user",
                "operationId": "getAllUnauthorizedUser",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpserver.HttpSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/httpserver.HttpUnauthenticatedUser"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login user given email and password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login",
                "operationId": "LoginUser",
                "parameters": [
                    {
                        "description": "email and password",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpLoginUserPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpserver.HttpSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/httpserver.LoginUser.HttpLoginUserPostResponse"
                                        },
                                        "message": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    }
                }
            }
        },
        "/user/me": {
            "post": {
                "description": "Get authenticated user info from jwt",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user info",
                "operationId": "GetUserInfo",
                "parameters": [
                    {
                        "type": "string",
                        "default": "token=xxx.xxx.xxx",
                        "description": "jwt token",
                        "name": "Cookie",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpserver.HttpSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        },
                                        "message": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Register a new user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new user",
                "operationId": "CreateNewUser",
                "parameters": [
                    {
                        "description": "User Informations",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpNewUserPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpserver.HttpSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/httpserver.CreateNewUser.HttpNewUserPostResponse"
                                        },
                                        "message": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Get user from ID for unauthorized users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user from ID",
                "operationId": "getUnauthorizedUser",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpserver.HttpSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/httpserver.HttpUnauthenticatedUser"
                                        },
                                        "message": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpserver.HttpError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpserver.CreateNewUser.HttpNewUserPostResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "httpserver.HttpError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "details": {
                    "type": "string",
                    "example": "Bad Request With More Info"
                },
                "error_doc_url": {
                    "type": "string",
                    "example": "https://example.com/docs/errors/invalid_id"
                },
                "error_type": {
                    "type": "string",
                    "example": "invalid_id"
                },
                "instance": {
                    "type": "string",
                    "example": "/api/v1/users/1"
                },
                "is_error": {
                    "type": "boolean",
                    "example": true
                },
                "message": {
                    "type": "string",
                    "example": "Bad Request"
                }
            }
        },
        "httpserver.HttpLoginUserPost": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "httpserver.HttpNewUserPost": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "httpserver.HttpSuccess": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {},
                "is_error": {
                    "type": "boolean",
                    "example": false
                },
                "message": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "httpserver.HttpUnauthenticatedUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "pfp": {
                    "type": "string"
                }
            }
        },
        "httpserver.LoginUser.HttpLoginUserPostResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "` + "`" + `json:\"email\"` + "`" + `",
                    "type": "string"
                },
                "firstName": {
                    "description": "` + "`" + `json:\"first_name\"` + "`" + `",
                    "type": "string"
                },
                "id": {
                    "description": "` + "`" + `json:\"id\"` + "`" + `",
                    "type": "integer"
                },
                "isBanned": {
                    "description": "` + "`" + `json:\"is_banned\"` + "`" + `",
                    "type": "boolean"
                },
                "lastName": {
                    "description": "` + "`" + `json:\"last_name\"` + "`" + `",
                    "type": "string"
                },
                "password": {
                    "description": "` + "`" + `json:\"password\"` + "`" + `",
                    "type": "string"
                },
                "pfp": {
                    "description": "` + "`" + `json:\"pfp\"` + "`" + `",
                    "type": "string"
                },
                "role": {
                    "description": "` + "`" + `json:\"role\"` + "`" + `",
                    "type": "string"
                },
                "specialMagicalSecretField": {
                    "description": "` + "`" + `json:\"special_magical_secret_field\"` + "`" + `",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Go Service Template",
	Description:      "User Management Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
