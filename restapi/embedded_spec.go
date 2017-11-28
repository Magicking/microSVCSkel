// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "MicroSVC API\n",
    "title": "MicroSVC",
    "version": "0.1.0"
  },
  "paths": {
    "/list": {
      "get": {
        "description": "Return various information",
        "summary": "Various information on the micro service",
        "operationId": "list",
        "responses": {
          "200": {
            "description": "Information of the micro service",
            "schema": {
              "$ref": "#/definitions/Information"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/push": {
      "post": {
        "description": "Push some data",
        "summary": "Push some data",
        "operationId": "push",
        "parameters": [
          {
            "type": "string",
            "name": "message",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Return some string",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unexpected error",
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
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "fields": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Information": {
      "type": "object",
      "properties": {
        "string": {
          "description": "Some string",
          "type": "string"
        }
      }
    }
  }
}`))
}