// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tag": {
            "get": {
                "description": "Show a tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tags"
                ],
                "summary": "Show Tag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tag identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ShowTagResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tags"
                ],
                "summary": "Create Tag",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateTagRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.CreateTagResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tags"
                ],
                "summary": "Delete Tag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tag identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.DeleteTagResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tags": {
            "get": {
                "description": "List all tags",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tags"
                ],
                "summary": "List Tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ListTagsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateTagRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "hour": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "handler.CreateTagResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.TagResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.DeleteTagResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.TagResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ErrorResponse": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ListTagsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.TagResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ShowTagResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.TagResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.TagResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "hour": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "tag": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
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
