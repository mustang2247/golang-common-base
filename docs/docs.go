// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mustang Kong",
            "email": "mustang2247@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/email/add": {
            "post": {
                "description": "添加 email 数据到数据库",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "email"
                ],
                "summary": "添加 email 数据到数据库",
                "parameters": [
                    {
                        "description": "新建email数据",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/email.EmailTextContent"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/email/delete/:contentId": {
            "delete": {
                "description": "删除 email 数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "email"
                ],
                "summary": "删除 email 数据",
                "responses": {
                    "0": {
                        "description": "{\"code\":0,\"message\":\"删除email成功\",\"data\":\"32\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/email/list": {
            "get": {
                "description": "获取 email 列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "email"
                ],
                "summary": "获取 email 列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "获取第几页的数据，默认为：1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页展示多少行，默认为：10",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "按照倒序或者顺序的方式排列，0或者-1为倒序，其他值为顺序",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "0": {
                        "description": "{\"code\":0,\"message\":\"获取分类列表成功\",\"data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/email/push": {
            "post": {
                "description": "推送 email 数据到数据库",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "email"
                ],
                "summary": "推送 email 数据到数据库",
                "parameters": [
                    {
                        "description": "新建email数据",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/email.EmailTextContent"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/email/update": {
            "put": {
                "description": "更新分类\t{\"name\":\"test1234\",\"key\":\"mus_test\",\"child\":{\"0-\":\"test\"}}\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "email"
                ],
                "summary": "更新 email 数据",
                "parameters": [
                    {
                        "description": "更新email数据",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/email.EmailTextContent"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "{\"code\":0,\"message\":\"更新email 数据成功\",\"data\":\"32\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "email.EmailTextContent": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "消息体",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "mail_cc": {
                    "description": "抄送",
                    "type": "string"
                },
                "mail_from": {
                    "description": "发送者",
                    "type": "string"
                },
                "mail_to": {
                    "description": "接受者",
                    "type": "string"
                },
                "subject": {
                    "description": "主题",
                    "type": "string"
                },
                "twf_created": {
                    "$ref": "#/definitions/jsonTime.JSONTime"
                },
                "twf_modified": {
                    "$ref": "#/definitions/jsonTime.JSONTime"
                }
            }
        },
        "jsonTime.JSONTime": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "golang-common-base API docs",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
