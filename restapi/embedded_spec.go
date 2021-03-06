package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import "encoding/json"

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "download",
    "title": "Download Test",
    "version": "1.0.0"
  },
  "host": "localhost:9098",
  "basePath": "/",
  "paths": {
    "/file": {
      "get": {
        "produces": [
          "application/octet-stream"
        ],
        "responses": {
          "200": {
            "description": "download a file",
            "schema": {
              "type": "file"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "file",
            "name": "attachment",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "The attachment has been uploaded."
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
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
}
