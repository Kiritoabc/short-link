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
        "/openapis/": {
            "get": {
                "description": "短链接跳转",
                "tags": [
                    "短链接跳转"
                ],
                "summary": "短链接跳转",
                "parameters": [
                    {
                        "type": "string",
                        "description": "短链接",
                        "name": "shortLink",
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
                                    "$ref": "#/definitions/utils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/openapis/hello": {
            "get": {
                "description": "健康检查",
                "tags": [
                    "短链接健康检查"
                ],
                "summary": "健康检查",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/openapis/short": {
            "post": {
                "description": "生成短链接",
                "tags": [
                    "生成短链接"
                ],
                "summary": "生成短链接",
                "parameters": [
                    {
                        "type": "string",
                        "description": "原始的长链接",
                        "name": "url",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "utils.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 7
                },
                "data": {},
                "msg": {
                    "type": "string",
                    "example": "successed"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Short Link API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
