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
            "name": "Mazettt",
            "email": "martin.d-herouville@epitech.eu"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/invite/{id}": {
            "post": {
                "description": "Answer an invite from an owner with an invite link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Create a new tenant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pending contract ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Tenant user data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created user data",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Missing fields",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "404": {
                        "description": "Pending contract not found",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "409": {
                        "description": "Email already exists",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Create a new user with owner role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Create a new owner",
                "parameters": [
                    {
                        "description": "Owner user data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created user data",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Missing fields",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "409": {
                        "description": "Email already exists",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/auth/token": {
            "post": {
                "description": "Authenticate user with email and password",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Authenticate user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "password / refresh_token",
                        "name": "grant_type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User email",
                        "name": "username",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User password",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Refresh token",
                        "name": "refresh_token",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token data",
                        "schema": {}
                    },
                    "400": {
                        "description": "Invalid grant_type",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get user profile information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user profile",
                "responses": {
                    "200": {
                        "description": "User data",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "404": {
                        "description": "Cannot find user",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get all users information",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "List of users",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get user information by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User data",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "404": {
                        "description": "Cannot find user",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "db.Role": {
            "type": "string",
            "enum": [
                "owner",
                "tenant"
            ],
            "x-enum-varnames": [
                "RoleOwner",
                "RoleTenant"
            ]
        },
        "models.UserRequest": {
            "type": "object",
            "required": [
                "email",
                "firstname",
                "lastname",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "models.UserResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/db.Role"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "utils.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/utils.ErrorCode"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "utils.ErrorCode": {
            "type": "string",
            "enum": [
                "invalid-password",
                "cannot-fetch-user",
                "cannot-find-user",
                "cannot-create-user",
                "no-claims",
                "cannot-decode-user",
                "missing-fields",
                "cannot-hash-password",
                "email-already-exists",
                "test-error",
                "too-many-requests",
                "pending-contract-not-found",
                "user-same-email-as-pending-contract",
                "contract-already-exists"
            ],
            "x-enum-varnames": [
                "InvalidPassword",
                "CannotFetchUser",
                "CannotFindUser",
                "CannotCreateUser",
                "NoClaims",
                "CannotDecodeUser",
                "MissingFields",
                "CannotHashPassword",
                "EmailAlreadyExists",
                "TestError",
                "TooManyRequests",
                "PendingContractNotFound",
                "UserSameEmailAsPendingContract",
                "ContractAlreadyExist"
            ]
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Enter the token with the ` + "`" + `Bearer: ` + "`" + ` prefix, e.g. \"Bearer abcde12345\".",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3001",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Immotep API",
	Description:      "This is the API used by the Immotep application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
