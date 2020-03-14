// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-03-14 14:29:44.046628 +0800 CST m=+0.032415935

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
            "name": "夜央 Oh oh oh oh oh oh",
            "email": "zoujh99@qq.com"
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
        "/api/v1/createRealEstate": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新建房地产(管理员)",
                "parameters": [
                    {
                        "description": "realEstate",
                        "name": "realEstate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.RealEstateRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/createSelling": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "发起销售",
                "parameters": [
                    {
                        "description": "selling",
                        "name": "selling",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.SellingRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/createSellingByBuy": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "买家购买",
                "parameters": [
                    {
                        "description": "sellingByBuy",
                        "name": "sellingByBuy",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.SellingByBuyRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/hello": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "测试输出Hello",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/queryAccountList": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取账户信息",
                "parameters": [
                    {
                        "description": "account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.AccountRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/queryRealEstateList": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取房地产信息(空json{}可以查询所有，指定proprietor可以查询指定业主名下房产)",
                "parameters": [
                    {
                        "description": "realEstateQuery",
                        "name": "realEstateQuery",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.RealEstateQueryRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/querySellingList": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "查询销售(可查询所有，也可根据发起销售人查询)(发起的)",
                "parameters": [
                    {
                        "description": "sellingListQuery",
                        "name": "sellingListQuery",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.SellingListQueryRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/querySellingListByBuyer": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "根据参与销售人、买家(买家AccountId)查询销售(参与的)",
                "parameters": [
                    {
                        "description": "sellingListQueryByBuy",
                        "name": "sellingListQueryByBuy",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.SellingListQueryByBuyRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/updateSelling": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "更新销售状态（买家确认、买卖家取消）Status取值为 完成\"done\"、取消\"cancelled\" 当处于销售中状态，卖家要取消时，buyer为\"\"空",
                "parameters": [
                    {
                        "description": "updateSelling",
                        "name": "updateSelling",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UpdateSellingRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "v1.AccountIdBody": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "string"
                }
            }
        },
        "v1.AccountRequestBody": {
            "type": "object",
            "properties": {
                "args": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.AccountIdBody"
                    }
                }
            }
        },
        "v1.RealEstateQueryRequestBody": {
            "type": "object",
            "properties": {
                "proprietor": {
                    "description": "所有者(业主)(业主AccountId)",
                    "type": "string"
                }
            }
        },
        "v1.RealEstateRequestBody": {
            "type": "object",
            "properties": {
                "accountId": {
                    "description": "操作人ID",
                    "type": "string"
                },
                "livingSpace": {
                    "description": "生活空间",
                    "type": "number"
                },
                "proprietor": {
                    "description": "所有者(业主)(业主AccountId)",
                    "type": "string"
                },
                "totalArea": {
                    "description": "总面积",
                    "type": "number"
                }
            }
        },
        "v1.SellingByBuyRequestBody": {
            "type": "object",
            "properties": {
                "buyer": {
                    "description": "买家(买家AccountId)",
                    "type": "string"
                },
                "objectOfSale": {
                    "description": "销售对象(正在出售的房地产RealEstateID)",
                    "type": "string"
                },
                "seller": {
                    "description": "发起销售人、卖家(卖家AccountId)",
                    "type": "string"
                }
            }
        },
        "v1.SellingListQueryByBuyRequestBody": {
            "type": "object",
            "properties": {
                "buyer": {
                    "description": "买家(买家AccountId)",
                    "type": "string"
                }
            }
        },
        "v1.SellingListQueryRequestBody": {
            "type": "object",
            "properties": {
                "seller": {
                    "description": "发起销售人、卖家(卖家AccountId)",
                    "type": "string"
                }
            }
        },
        "v1.SellingRequestBody": {
            "type": "object",
            "properties": {
                "objectOfSale": {
                    "description": "销售对象(正在出售的房地产RealEstateID)",
                    "type": "string"
                },
                "price": {
                    "description": "价格",
                    "type": "number"
                },
                "salePeriod": {
                    "description": "智能合约的有效期(单位为天)",
                    "type": "integer"
                },
                "seller": {
                    "description": "发起销售人、卖家(卖家AccountId)",
                    "type": "string"
                }
            }
        },
        "v1.UpdateSellingRequestBody": {
            "type": "object",
            "properties": {
                "buyer": {
                    "description": "买家(买家AccountId)",
                    "type": "string"
                },
                "objectOfSale": {
                    "description": "销售对象(正在出售的房地产RealEstateID)",
                    "type": "string"
                },
                "seller": {
                    "description": "发起销售人、卖家(卖家AccountId)",
                    "type": "string"
                },
                "status": {
                    "description": "需要更改的状态",
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "基于区块链技术的房地产交易系统api文档",
	Description: "基于区块链技术的房地产交易系统api文档",
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
