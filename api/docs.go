// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Harsh Mittal",
            "email": "harshmittal2210@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "get string by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ping"
                ],
                "summary": "Hello API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.JSONResp"
                        }
                    }
                }
            }
        },
        "/accounts/": {
            "get": {
                "description": "get all bank account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bank Accounts"
                ],
                "summary": "Get All Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Items Per Page",
                        "name": "items_per_page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_bytesByHarsh_go-my-info_models.PaginatedListResp-models_BankAccount"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "create new bank account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bank Accounts"
                ],
                "summary": "Create Account",
                "parameters": [
                    {
                        "description": "Bank Account Body",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddBankAccountReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.BankAccount"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            }
        },
        "/accounts/{account_id}": {
            "get": {
                "description": "get bank account details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bank Accounts"
                ],
                "summary": "Get Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bank Account ID",
                        "name": "account_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.BankAccount"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "update bank account details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bank Accounts"
                ],
                "summary": "Update Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bank Account ID",
                        "name": "account_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Bank Account Body",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateBankAccountReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            }
        },
        "/banks/list": {
            "get": {
                "description": "get all bank list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bank"
                ],
                "summary": "Get Bank List",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Items Per Page",
                        "name": "items_per_page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_bytesByHarsh_go-my-info_models.PaginatedListResp-models_Bank"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            }
        },
        "/banks/register": {
            "post": {
                "description": "create a new bank",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bank"
                ],
                "summary": "Create Bank",
                "parameters": [
                    {
                        "description": "Create Bank Body",
                        "name": "page",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateBankReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Bank"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            }
        },
        "/users/add": {
            "post": {
                "description": "create new user by admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create User By Admin",
                "parameters": [
                    {
                        "description": "User Body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserByAdminReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            }
        },
        "/users/list": {
            "get": {
                "description": "user list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get User List",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Items Per Page",
                        "name": "items_per_page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_bytesByHarsh_go-my-info_models.PaginatedListResp-models_User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "get login token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login User",
                "parameters": [
                    {
                        "description": "Login Body",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AuthReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.AuthResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "description": "get user details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "update user details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "update User",
                "parameters": [
                    {
                        "description": "User Body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateUserReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            }
        },
        "/users/me/password": {
            "put": {
                "description": "update user password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update User Password",
                "parameters": [
                    {
                        "description": "Password Body",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdatePasswordReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "create new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "User Body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            }
        },
        "/users/{username}": {
            "get": {
                "description": "get another user details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "create new user by admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create User By Admin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User Body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateUserReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete user from db by admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user from DB",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONerrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_bytesByHarsh_go-my-info_internal_database.BankType": {
            "type": "string",
            "enum": [
                "central",
                "cooperative",
                "commercial",
                "regional",
                "local",
                "specialized",
                "small_finance",
                "payments"
            ],
            "x-enum-varnames": [
                "BankTypeCentral",
                "BankTypeCooperative",
                "BankTypeCommercial",
                "BankTypeRegional",
                "BankTypeLocal",
                "BankTypeSpecialized",
                "BankTypeSmallFinance",
                "BankTypePayments"
            ]
        },
        "github_com_bytesByHarsh_go-my-info_models.PaginatedListResp-models_Bank": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Bank"
                    }
                },
                "has_more": {
                    "type": "boolean"
                },
                "items_per_page": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "total_count": {
                    "type": "integer"
                }
            }
        },
        "github_com_bytesByHarsh_go-my-info_models.PaginatedListResp-models_BankAccount": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.BankAccount"
                    }
                },
                "has_more": {
                    "type": "boolean"
                },
                "items_per_page": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "total_count": {
                    "type": "integer"
                }
            }
        },
        "github_com_bytesByHarsh_go-my-info_models.PaginatedListResp-models_User": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                },
                "has_more": {
                    "type": "boolean"
                },
                "items_per_page": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "total_count": {
                    "type": "integer"
                }
            }
        },
        "models.AddBankAccountReq": {
            "type": "object",
            "required": [
                "account_number",
                "account_type",
                "balance",
                "bank_id",
                "currency",
                "name"
            ],
            "properties": {
                "account_number": {
                    "type": "string"
                },
                "account_type": {
                    "type": "string"
                },
                "balance": {
                    "type": "string"
                },
                "bank_id": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.AuthReq": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.AuthResp": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "models.Bank": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "bank_type": {
                    "$ref": "#/definitions/github_com_bytesByHarsh_go-my-info_internal_database.BankType"
                },
                "contact_email": {
                    "type": "string"
                },
                "contact_phone": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "established_year": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.BankAccount": {
            "type": "object",
            "properties": {
                "account_number": {
                    "type": "string"
                },
                "account_type": {
                    "type": "string"
                },
                "balance": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.CreateBankReq": {
            "type": "object",
            "required": [
                "address",
                "bank_type",
                "contact_email",
                "established_year",
                "name"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "bank_type": {
                    "$ref": "#/definitions/github_com_bytesByHarsh_go-my-info_internal_database.BankType"
                },
                "contact_email": {
                    "type": "string"
                },
                "contact_phone": {
                    "type": "string"
                },
                "established_year": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.CreateUserByAdminReq": {
            "type": "object",
            "required": [
                "email",
                "is_active",
                "name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "is_superuser": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.CreateUserReq": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.JSONResp": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.JSONerrResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.UpdateBankAccountReq": {
            "type": "object",
            "required": [
                "account_number",
                "account_type",
                "balance",
                "currency",
                "name"
            ],
            "properties": {
                "account_number": {
                    "type": "string"
                },
                "account_type": {
                    "type": "string"
                },
                "balance": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.UpdatePasswordReq": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "password": {
                    "type": "string"
                }
            }
        },
        "models.UpdateUserReq": {
            "type": "object",
            "required": [
                "email",
                "name",
                "phone_num",
                "profile_img",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone_num": {
                    "type": "string"
                },
                "profile_img": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "is_superuser": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "profile_img": {
                    "type": "string"
                },
                "role": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
