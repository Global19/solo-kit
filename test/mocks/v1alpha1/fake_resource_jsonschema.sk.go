// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

var FakeResourceJsonSchema = `
{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/testing.solo.io.FakeResource",
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
        "testing.solo.io.FakeResource": {
            "$schema": "http://json-schema.org/draft-04/schema#",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "metadata": {
                    "$ref": "#/definitions/core.solo.io.Metadata",
                    "additionalProperties": true,
                    "type": "object",
                    "title": "core.solo.io.Metadata"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "testing.solo.io.FakeResource"
        }
    }
}
`
