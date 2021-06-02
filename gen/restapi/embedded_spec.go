// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "UserServer Swagger",
    "contact": {
      "name": "Contact the developer",
      "url": "https://t.me/viewsharp"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/user": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Create user",
        "operationId": "postUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserEditableData"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid format",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "tags": [
          "user"
        ],
        "summary": "Get user by id",
        "operationId": "getUser",
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
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "404": {
            "description": "User not found"
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "user"
        ],
        "summary": "Updated user",
        "operationId": "putUser",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserEditableData"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid format",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "User not found"
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "user"
        ],
        "summary": "Delete user",
        "operationId": "deleteUser",
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
            "description": "Success"
          },
          "404": {
            "description": "User not found"
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "required": [
        "id",
        "name",
        "birth_date"
      ],
      "properties": {
        "birth_date": {
          "type": "string",
          "format": "date"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "UserEditableData": {
      "type": "object",
      "required": [
        "name",
        "birth_date"
      ],
      "properties": {
        "birth_date": {
          "type": "string",
          "format": "date"
        },
        "name": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations about user",
      "name": "user"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "UserServer Swagger",
    "contact": {
      "name": "Contact the developer",
      "url": "https://t.me/viewsharp"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/user": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Create user",
        "operationId": "postUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserEditableData"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid format",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "tags": [
          "user"
        ],
        "summary": "Get user by id",
        "operationId": "getUser",
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
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "404": {
            "description": "User not found"
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "user"
        ],
        "summary": "Updated user",
        "operationId": "putUser",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserEditableData"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid format",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "User not found"
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "user"
        ],
        "summary": "Delete user",
        "operationId": "deleteUser",
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
            "description": "Success"
          },
          "404": {
            "description": "User not found"
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "required": [
        "id",
        "name",
        "birth_date"
      ],
      "properties": {
        "birth_date": {
          "type": "string",
          "format": "date"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "UserEditableData": {
      "type": "object",
      "required": [
        "name",
        "birth_date"
      ],
      "properties": {
        "birth_date": {
          "type": "string",
          "format": "date"
        },
        "name": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations about user",
      "name": "user"
    }
  ]
}`))
}
