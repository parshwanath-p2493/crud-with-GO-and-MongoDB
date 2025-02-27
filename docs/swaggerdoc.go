package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

var doc = `{
  "swagger": "2.0",
  "info": {
    "title": "Go CRUD API",
    "description": "CRUD API using Gin and MongoDB",
    "version": "1.0"
  },
  "host": "{{.Host}}",
  "basePath": "{{.BasePath}}",
  "schemes": ["http"],
  "paths": {
    "/getdata": {
      "get": {
        "tags": ["CRUD"],
        "summary": "Get all records",
        "responses": {
          "200": {"description": "Success"},
          "500": {"description": "Server error"}
        }
      }
    },
    "/adddata": {
      "post": {
        "tags": ["CRUD"],
        "summary": "Add a new record",
        "parameters": [
          {
            "in": "body",
            "name": "data",
            "description": "Data object",
            "schema": {
              "type": "object",
              "properties": {
                "name": {"type": "string"},
                "email": {"type": "string"}
              }
            }
          }
        ],
        "responses": {
          "201": {"description": "Created"},
          "400": {"description": "Bad request"}
        }
      }
    },
    "/update/{id}": {
      "put": {
        "tags": ["CRUD"],
        "summary": "Update a record",
        "parameters": [
          {"in": "path", "name": "id", "required": true, "type": "string"},
          {"in": "body", "name": "data", "schema": {"type": "object", "properties": {"name": {"type": "string"}, "email": {"type": "string"}}}}
        ],
        "responses": {
          "200": {"description": "Updated"},
          "400": {"description": "Bad request"},
          "404": {"description": "Not found"}
        }
      }
    },
    "/delete/{id}": {
      "delete": {
        "tags": ["CRUD"],
        "summary": "Delete a record",
        "parameters": [{"in": "path", "name": "id", "required": true, "type": "string"}],
        "responses": {
          "200": {"description": "Deleted"},
          "404": {"description": "Not found"}
        }
      }
    }
  }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
