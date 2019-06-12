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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Micro Application Manifest Manager",
    "title": "Microfest",
    "contact": {
      "name": "J. Iqbal",
      "url": "https://github.com/LGUG2Z/microfest",
      "email": "jade@beamery.com"
    },
    "version": "0.1.0"
  },
  "paths": {
    "/backup": {
      "post": {
        "security": [
          {
            "APIKeyHeader": []
          }
        ],
        "produces": [
          "text/plain"
        ],
        "summary": "Backs up the database to a GCS bucket",
        "operationId": "PostBackup",
        "parameters": [
          {
            "type": "string",
            "description": "GCS bucket name",
            "name": "bucket",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Backup Created",
            "schema": {
              "type": "string"
            }
          },
          "401": {
            "description": "Unauthorized",
            "headers": {
              "WWW-Authenticate": {
                "type": "string",
                "description": "Authorization information is missing or invalid"
              }
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
    },
    "/manifest": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Gets the latest manifest",
        "operationId": "GetManifest",
        "parameters": [
          {
            "type": "string",
            "description": "The environment hostname",
            "name": "host",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "description": "The latest manifest",
              "type": "object"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "APIKeyHeader": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "text/plain"
        ],
        "summary": "Submits a patch to create a new manifest",
        "operationId": "PutManifest",
        "parameters": [
          {
            "type": "string",
            "description": "The environment hostname",
            "name": "host",
            "in": "query",
            "required": true
          },
          {
            "description": "The manifest patch to submit",
            "name": "microfest",
            "in": "body",
            "schema": {
              "type": "object",
              "maxProperties": 3,
              "required": [
                "release",
                "manifest",
                "updated"
              ],
              "properties": {
                "manifest": {
                  "type": "object"
                },
                "release": {
                  "type": "string"
                },
                "updated": {
                  "type": "array",
                  "items": {
                    "type": "string"
                  }
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Manifest Created",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Bad Request Body"
          },
          "401": {
            "description": "Unauthorized",
            "headers": {
              "WWW-Authenticate": {
                "type": "string",
                "description": "Authorization information is missing or invalid"
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "APIKeyHeader": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "text/plain"
        ],
        "summary": "Submits a new manifest",
        "operationId": "PostManifest",
        "parameters": [
          {
            "type": "string",
            "description": "The environment hostname",
            "name": "host",
            "in": "query",
            "required": true
          },
          {
            "description": "The manifest to submit",
            "name": "microfest",
            "in": "body",
            "schema": {
              "type": "object",
              "maxProperties": 3,
              "required": [
                "release",
                "manifest",
                "updated"
              ],
              "properties": {
                "manifest": {
                  "type": "object"
                },
                "release": {
                  "type": "string"
                },
                "updated": {
                  "type": "array",
                  "items": {
                    "type": "string"
                  }
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Manifest Created",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Bad Request Body"
          },
          "401": {
            "description": "Unauthorized",
            "headers": {
              "WWW-Authenticate": {
                "type": "string",
                "description": "Authorization information is missing or invalid"
              }
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
    "principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "APIKeyHeader": {
      "type": "apiKey",
      "name": "X-API-KEY",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Micro Application Manifest Manager",
    "title": "Microfest",
    "contact": {
      "name": "J. Iqbal",
      "url": "https://github.com/LGUG2Z/microfest",
      "email": "jade@beamery.com"
    },
    "version": "0.1.0"
  },
  "paths": {
    "/backup": {
      "post": {
        "security": [
          {
            "APIKeyHeader": []
          }
        ],
        "produces": [
          "text/plain"
        ],
        "summary": "Backs up the database to a GCS bucket",
        "operationId": "PostBackup",
        "parameters": [
          {
            "type": "string",
            "description": "GCS bucket name",
            "name": "bucket",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Backup Created",
            "schema": {
              "type": "string"
            }
          },
          "401": {
            "description": "Unauthorized",
            "headers": {
              "WWW-Authenticate": {
                "type": "string",
                "description": "Authorization information is missing or invalid"
              }
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
    },
    "/manifest": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Gets the latest manifest",
        "operationId": "GetManifest",
        "parameters": [
          {
            "type": "string",
            "description": "The environment hostname",
            "name": "host",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "description": "The latest manifest",
              "type": "object"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "APIKeyHeader": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "text/plain"
        ],
        "summary": "Submits a patch to create a new manifest",
        "operationId": "PutManifest",
        "parameters": [
          {
            "type": "string",
            "description": "The environment hostname",
            "name": "host",
            "in": "query",
            "required": true
          },
          {
            "description": "The manifest patch to submit",
            "name": "microfest",
            "in": "body",
            "schema": {
              "type": "object",
              "maxProperties": 3,
              "required": [
                "release",
                "manifest",
                "updated"
              ],
              "properties": {
                "manifest": {
                  "type": "object"
                },
                "release": {
                  "type": "string"
                },
                "updated": {
                  "type": "array",
                  "items": {
                    "type": "string"
                  }
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Manifest Created",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Bad Request Body"
          },
          "401": {
            "description": "Unauthorized",
            "headers": {
              "WWW-Authenticate": {
                "type": "string",
                "description": "Authorization information is missing or invalid"
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "APIKeyHeader": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "text/plain"
        ],
        "summary": "Submits a new manifest",
        "operationId": "PostManifest",
        "parameters": [
          {
            "type": "string",
            "description": "The environment hostname",
            "name": "host",
            "in": "query",
            "required": true
          },
          {
            "description": "The manifest to submit",
            "name": "microfest",
            "in": "body",
            "schema": {
              "type": "object",
              "maxProperties": 3,
              "required": [
                "release",
                "manifest",
                "updated"
              ],
              "properties": {
                "manifest": {
                  "type": "object"
                },
                "release": {
                  "type": "string"
                },
                "updated": {
                  "type": "array",
                  "items": {
                    "type": "string"
                  }
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Manifest Created",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Bad Request Body"
          },
          "401": {
            "description": "Unauthorized",
            "headers": {
              "WWW-Authenticate": {
                "type": "string",
                "description": "Authorization information is missing or invalid"
              }
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
    "principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "APIKeyHeader": {
      "type": "apiKey",
      "name": "X-API-KEY",
      "in": "header"
    }
  }
}`))
}
