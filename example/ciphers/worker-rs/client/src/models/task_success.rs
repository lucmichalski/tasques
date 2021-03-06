/*
 * Tasques API
 *
 * A Task queue backed by Elasticsearch
 *
 * The version of the OpenAPI document: 0.0.1
 *
 * Generated by: https://openapi-generator.tech
 */

#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct TaskSuccess {
    #[serde(rename = "data", skip_serializing_if = "Option::is_none")]
    pub data: Option<serde_json::Value>,
}

impl TaskSuccess {
    pub fn new() -> TaskSuccess {
        TaskSuccess { data: None }
    }
}
