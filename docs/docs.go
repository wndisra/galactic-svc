// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/spaceship": {
            "get": {
                "description": "Get all spaceships.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spaceship"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Create new spaceship.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spaceship"
                ],
                "parameters": [
                    {
                        "description": "Request body (JSON)",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/spaceship.createRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/spaceship/{id}": {
            "get": {
                "description": "Fetch existing spaceship by a specific ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spaceship"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Spaceship ID (integer)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "Delete existing spaceship by a specific ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spaceship"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Spaceship ID (integer)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "patch": {
                "description": "Update existing spaceship by a specific ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spaceship"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Spaceship ID (integer)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body (JSON)",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/spaceship.updateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "spaceship.armamentReq": {
            "type": "object",
            "properties": {
                "qty": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "spaceship.createRequest": {
            "type": "object",
            "properties": {
                "armament": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/spaceship.armamentReq"
                    }
                },
                "class": {
                    "type": "string"
                },
                "crew": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "spaceship.updateRequest": {
            "type": "object",
            "properties": {
                "armament": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/spaceship.armamentReq"
                    }
                },
                "class": {
                    "type": "string"
                },
                "crew": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
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
	Title:            "Galactic Service APIs",
	Description:      "The server APIs documentation for Galactic.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
