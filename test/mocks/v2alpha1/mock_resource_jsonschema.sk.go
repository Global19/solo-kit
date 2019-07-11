// Code generated by solo-kit. DO NOT EDIT.

package v2alpha1

var MockResourceJsonSchema = `
{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/testing.solo.io.MockResource",
    "definitions": {
        "core.solo.io.Metadata": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "annotations": {
                    "additionalProperties": true,
                    "type": "object",
                    "title": "core.solo.io.Metadata.AnnotationsEntry"
                },
                "cluster": {
                    "type": "string"
                },
                "labels": {
                    "additionalProperties": true,
                    "type": "object",
                    "title": "core.solo.io.Metadata.LabelsEntry"
                },
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "resourceVersion": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "core.solo.io.Metadata"
        },
        "core.solo.io.Status": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "reason": {
                    "type": "string"
                },
                "reportedBy": {
                    "type": "string"
                },
                "state": {
                    "enum": [
                        "Pending",
                        0,
                        "Accepted",
                        1,
                        "Rejected",
                        2
                    ],
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        }
                    ]
                },
                "subresourceStatuses": {
                    "additionalProperties": true,
                    "type": "object",
                    "title": "core.solo.io.Status.SubresourceStatusesEntry"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "core.solo.io.Status"
        },
        "testing.solo.io.MockResource": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "data.json": {
                    "type": "string"
                },
                "metadata": {
                    "$ref": "#/definitions/core.solo.io.Metadata",
                    "additionalProperties": true,
                    "type": "object",
                    "title": "core.solo.io.Metadata"
                },
                "oneofOne": {
                    "type": "string"
                },
                "oneofTwo": {
                    "type": "boolean"
                },
                "someDumbField": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/core.solo.io.Status",
                    "additionalProperties": true,
                    "type": "object",
                    "title": "core.solo.io.Status"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "testing.solo.io.MockResource"
        }
    }
}
`
