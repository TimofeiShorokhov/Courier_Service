// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
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
        "/orders": {
            "put": {
                "description": "assign order to courier",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "UpdateOrder",
                "operationId": "UpdateOrder",
                "parameters": [
                    {
                        "description": "id courier/id order",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dao.Order"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dao.Order": {
            "type": "object",
            "properties": {
                "courier_id": {
                    "type": "integer"
                },
                "customer_address": {
                    "type": "string"
                },
                "delivery_service_id": {
                    "type": "integer"
                },
                "delivery_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "order_date": {
                    "type": "string"
                },
                "picked": {
                    "type": "boolean"
                },
                "restaurant_address": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "COURIER_SERVICE",
	Description:      "API Server for TodoList Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
