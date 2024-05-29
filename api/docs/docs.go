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
        "/bocchi/access_token/get": {
            "get": {
                "description": "get available access-token by refresh-token",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get_access-token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/access_token/verify": {
            "get": {
                "description": "verify access-token if it is expired or not",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "verify_access-token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/apply": {
            "get": {
                "description": "apply to join the party",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "apply_party",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "活动id",
                        "name": "party_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/apply/list": {
            "get": {
                "description": "show party's applicants",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "apply_list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "活动id",
                        "name": "party_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/apply/permit": {
            "get": {
                "description": "permit user to join the party",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "permit_join",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "活动id",
                        "name": "party_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "member_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/create": {
            "post": {
                "description": "create party",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create_party",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标题",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "介绍",
                        "name": "content",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "类型",
                        "name": "type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "活动省份",
                        "name": "province",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "活动城市",
                        "name": "city",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "开始时间(例:2006-01-02)",
                        "name": "start_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "结束时间(例:2006-01-02)",
                        "name": "end_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/info": {
            "get": {
                "responses": {}
            }
        },
        "/bocchi/party/itinerary/create": {
            "post": {
                "description": "create itinerary for party",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create_itinerary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标题",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "类型 1:route;2:activity;3:accommodation;4:eat;5:other",
                        "name": "action_type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "活动id",
                        "name": "party_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "坐标 1时不需要",
                        "name": "rectangle",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "路线 1时需要",
                        "name": "route_json",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "备注",
                        "name": "remark",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "计划起始时间(例:2006-01-02 15:40)",
                        "name": "schedule_start_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "计划结束时间(例:2006-01-02 16:40)",
                        "name": "schedule_end_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/itinerary/info": {
            "get": {
                "responses": {}
            }
        },
        "/bocchi/party/itinerary/merge": {
            "get": {
                "description": "merge itinerary into party",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "merge_itinerary",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "行程id",
                        "name": "itinerary_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "活动id",
                        "name": "party_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/itinerary/sequence/change": {
            "post": {
                "description": "show party's itineraries order by sequence",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "change_itinerary_sequence",
                "parameters": [
                    {
                        "type": "array",
                        "description": "行程id列表",
                        "name": "itinerary_id_list",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "description": "对应行程的顺序列表",
                        "name": "sequence_list",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/itinerary/show": {
            "get": {
                "description": "show party's itineraries order by sequence",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "show_party_itinerary",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "活动id",
                        "name": "party_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/members": {
            "get": {
                "description": "get members who have join the party",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get_party_members",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "活动id",
                        "name": "party_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/party/search": {
            "post": {
                "description": "search party",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "search_party",
                "parameters": [
                    {
                        "type": "string",
                        "description": "搜索内容",
                        "name": "content",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "活动类型",
                        "name": "party_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "活动省份",
                        "name": "province",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "活动城市",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "几天后开始",
                        "name": "start_time_duration",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "搜索方式(0:默认;1:按开始时间;2:按活动已有人数)",
                        "name": "search_type",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/poi/comment/create": {
            "post": {
                "description": "comment poi",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "comment_create",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "poi_id",
                        "name": "poi_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "正文",
                        "name": "content",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/poi/comment/delete": {
            "get": {
                "description": "delete your comment",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "comment_delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "评论id",
                        "name": "comment_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/poi/comment/list": {
            "get": {
                "description": "show poi's comments",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "comment_list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "poi_id",
                        "name": "poi_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/user/avatar/upload": {
            "put": {
                "description": "revise user's avatar",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "PutAvatar",
                "parameters": [
                    {
                        "type": "file",
                        "description": "头像",
                        "name": "avatar_file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/user/info": {
            "get": {
                "description": "get user's info",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/user/login/": {
            "post": {
                "description": "login to get your auth token",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "otp",
                        "name": "otp",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/user/register/": {
            "post": {
                "description": "userRegister",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/user/signature": {
            "post": {
                "description": "revise signature",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "signature",
                "parameters": [
                    {
                        "type": "string",
                        "description": "signature",
                        "name": "signature",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
            }
        },
        "/bocchi/user/switch2fa": {
            "post": {
                "description": "switch on/off 2fa mode",
                "consumes": [
                    "json/form"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "switch_2fa",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "关闭:0;开启:1",
                        "name": "action_type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "totp",
                        "name": "totp",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "access-token",
                        "name": "access-token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "refresh-token",
                        "name": "refresh-token",
                        "in": "header"
                    }
                ],
                "responses": {}
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
