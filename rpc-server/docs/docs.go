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
        "/chain/balance/{address}": {
            "get": {
                "description": "查詢指定帳戶餘額\n\nError code list\n- 400: PARAMETER_INVALID\n- 500: INTERNAL_PROCESS",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chain"
                ],
                "summary": "查詢指定帳戶餘額",
                "parameters": [
                    {
                        "type": "string",
                        "default": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
                        "example": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
                        "description": "帳戶地址",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/BalanceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorMessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorMessageResponse"
                        }
                    }
                }
            }
        },
        "/chain/blocks/height": {
            "get": {
                "description": "查詢最新區塊高度\n\nError code list\n- 500: INTERNAL_PROCESS",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chain"
                ],
                "summary": "查詢最新區塊高度",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/BlockResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorMessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorMessageResponse"
                        }
                    }
                }
            }
        },
        "/chain/faucet/{address}": {
            "post": {
                "description": "取得 0.02 測試幣 (模擬發送交易)\n\nError code list\n- 400: PARAMETER_INVALID\n- 500: INTERNAL_PROCESS",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chain"
                ],
                "summary": "取得 0.02 測試幣 (模擬發送交易)",
                "parameters": [
                    {
                        "type": "string",
                        "default": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
                        "example": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
                        "description": "接收帳戶地址",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/SuccessMessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorMessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorMessageResponse"
                        }
                    }
                }
            }
        },
        "/chain/tx": {
            "get": {
                "description": "查詢交易資料\n\nError code list\n- 400: PARAMETER_INVALID\n- 500: INTERNAL_PROCESS",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chain"
                ],
                "summary": "查詢交易資料",
                "parameters": [
                    {
                        "minimum": 10,
                        "type": "integer",
                        "default": 10,
                        "example": 10,
                        "description": "Page size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "0x8De0c53FC169BA09F111aA4170697e8CF42CCbBe",
                        "example": "0x8De0c53FC169BA09F111aA4170697e8CF42CCbBe",
                        "description": "Account address",
                        "name": "address",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/TxResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorMessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorMessageResponse"
                        }
                    }
                }
            }
        },
        "/chain/tx/{address}/sync": {
            "post": {
                "description": "更新指定地址的交易資料進資料庫，再次觸發會等待上次同步完畢，避免重複觸發\n\nError code list\n- 400: PARAMETER_INVALID\n- 500: INTERNAL_PROCESS",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chain"
                ],
                "summary": "更新指定地址的交易資料進資料庫，再次觸發會等待上次同步完畢，避免重複觸發",
                "parameters": [
                    {
                        "type": "string",
                        "default": "0x8De0c53FC169BA09F111aA4170697e8CF42CCbBe",
                        "example": "0x8De0c53FC169BA09F111aA4170697e8CF42CCbBe",
                        "description": "Account address",
                        "name": "address",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/SuccessMessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorMessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorMessageResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "BalanceResponse": {
            "description": "Balance response",
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number",
                    "example": 0.02
                },
                "wei": {
                    "type": "integer",
                    "example": 20000000000000000
                }
            }
        },
        "BlockResponse": {
            "description": "Block response",
            "type": "object",
            "properties": {
                "height": {
                    "type": "integer",
                    "example": 1030
                }
            }
        },
        "ErrorMessageResponse": {
            "description": "Error message response for 4xx and 5xx errors",
            "type": "object",
            "properties": {
                "code": {
                    "description": "HTTP status code",
                    "type": "integer",
                    "example": 400
                },
                "name": {
                    "description": "Error name (key)",
                    "type": "string",
                    "example": "PARAMETER_INVALID"
                }
            }
        },
        "SuccessMessageResponse": {
            "description": "Success message response for 200",
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "name": {
                    "type": "string",
                    "example": "SUCCESS"
                }
            }
        },
        "TxResponse": {
            "description": "Transaction response",
            "type": "object",
            "properties": {
                "from": {
                    "type": "string",
                    "example": "0x8de0c53fc169ba09f111aa4170697e8cf42ccbbe"
                },
                "height": {
                    "type": "integer",
                    "example": 1
                },
                "to": {
                    "type": "string",
                    "example": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
                },
                "wei": {
                    "type": "integer",
                    "example": 1000000000000000000
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
	Title:            "Ethereum Private Chain RPC Server",
	Description:      "Ethereum Private Chain RPC Server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
