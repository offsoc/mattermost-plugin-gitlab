// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package webhook

const OpenMergeRequest = `{
	"object_kind":"merge_request",
	"event_type":"merge_request",
	"user":{
		"name":"Administrator",
		"username":"root",
		"avatar_url":"https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
	},
	"project":{
		"id":24,
		"name":"webhook",
		"description":"",
		"web_url":"http://localhost:3000/manland/webhook",
		"avatar_url":null,
		"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
		"git_http_url":"http://localhost:3000/manland/webhook.git",
		"namespace":"manland",
		"visibility_level":20,
		"path_with_namespace":"manland/webhook",
		"default_branch":"master",
		"ci_config_path":null,
		"homepage":"http://localhost:3000/manland/webhook",
		"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
		"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
		"http_url":"http://localhost:3000/manland/webhook.git"
	},
	"object_attributes":{
		"assignee_id":50,
		"author_id":1,
		"created_at":"2019-04-03 21:07:32 UTC",
		"description":"test open merge request",
		"head_pipeline_id":57,
		"id":35,
		"iid":4,
		"last_edited_at":null,
		"last_edited_by_id":null,
		"merge_commit_sha":null,
		"merge_error":null,
		"merge_params":{
			"force_remove_source_branch":null
		},
		"merge_status":"unchecked",
		"merge_user_id":null,
		"merge_when_pipeline_succeeds":false,
		"milestone_id":null,
		"source_branch":"master",
		"source_project_id":25,
		"state":"opened",
		"target_branch":"master",
		"target_project_id":24,
		"time_estimate":0,
		"title":"Master",
		"updated_at":"2019-04-03 21:07:32 UTC",
		"updated_by_id":null,
		"url":"http://localhost:3000/manland/webhook/merge_requests/4",
		"source":{
			"id":25,
			"name":"webhook",
			"description":"",
			"web_url":"http://localhost:3000/root/webhook",
			"avatar_url":null,
			"git_ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
			"git_http_url":"http://localhost:3000/root/webhook.git",
			"namespace":"root",
			"visibility_level":20,
			"path_with_namespace":"root/webhook",
			"default_branch":"master",
			"ci_config_path":null,
			"homepage":"http://localhost:3000/root/webhook",
			"url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
			"ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
			"http_url":"http://localhost:3000/root/webhook.git"
		},
		"target":{
			"id":24,
			"name":"webhook",
			"description":"",
			"web_url":"http://localhost:3000/manland/webhook",
			"avatar_url":null,
			"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
			"git_http_url":"http://localhost:3000/manland/webhook.git",
			"namespace":"manland",
			"visibility_level":20,
			"path_with_namespace":"manland/webhook",
			"default_branch":"master",
			"ci_config_path":null,
			"homepage":"http://localhost:3000/manland/webhook",
			"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
			"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
			"http_url":"http://localhost:3000/manland/webhook.git"
		},
		"last_commit":{
			"id":"1fd967c14f8265a6056525c343d984ce56472d5c",
			"message":"Update README.md",
			"timestamp":"2019-04-03T21:04:58Z",
			"url":"http://localhost:3000/manland/webhook/commit/1fd967c14f8265a6056525c343d984ce56472d5c",
			"author":{
				"name":"Administrator",
				"email":"admin@example.com"
			}
		},
		"work_in_progress":false,
		"total_time_spent":0,
		"human_total_time_spent":null,
		"human_time_estimate":null,
		"assignee_ids": [
			50
		],
		"action":"open"
	},
	"labels":[],
	"changes":{
		"head_pipeline_id":{
			"previous":null,
			"current":57
		},
		"updated_at":{
			"previous":"2019-04-03 21:07:32 UTC",
			"current":"2019-04-03 21:07:32 UTC"
		},
		"assignee":{
			"previous":null,
			"current":{
				"name":"manland",
				"username":"manland",
				"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
			}
		},
		"total_time_spent":{
			"previous":null,
			"current":0
		}
	},
	"repository":{
		"name":"webhook",
		"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
		"description":"",
		"homepage":"http://localhost:3000/manland/webhook"
	},
	"assignee":{
		"name":"manland",
		"username":"manland",
		"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
	}
	}`

const ReopenMerge = `{
			"object_kind":"merge_request",
			"event_type":"merge_request",
			"user":{
				"name":"manland",
				"username":"manland",
				"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
			},
			"project":{
				"id":24,
				"name":"webhook",
				"description":"",
				"web_url":"http://localhost:3000/manland/webhook",
				"avatar_url":null,
				"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
				"git_http_url":"http://localhost:3000/manland/webhook.git",
				"namespace":"manland",
				"visibility_level":20,
				"path_with_namespace":"manland/webhook",
				"default_branch":"master",
				"ci_config_path":null,
				"homepage":"http://localhost:3000/manland/webhook",
				"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
				"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
				"http_url":"http://localhost:3000/manland/webhook.git"
			},
			"object_attributes":{
				"assignee_id":null,
				"author_id":1,
				"created_at":"2019-04-01 20:25:41 UTC",
				"description":"# Et voila\\n\\nC'est tro pbien",
				"head_pipeline_id":56,
				"id":34,
				"iid":1,
				"last_edited_at":null,
				"last_edited_by_id":null,
				"merge_commit_sha":null,
				"merge_error":null,
				"merge_params":{
					"force_remove_source_branch":null
				},
				"merge_status":"can_be_merged",
				"merge_user_id":null,
				"merge_when_pipeline_succeeds":false,
				"milestone_id":null,
				"source_branch":"master",
				"source_project_id":25,
				"state":"opened",
				"target_branch":"master",
				"target_project_id":24,
				"time_estimate":0,
				"title":"Update README.md",
				"updated_at":"2019-04-01 21:34:16 UTC",
				"updated_by_id":50,
				"url":"http://localhost:3000/manland/webhook/merge_requests/1",
				"source":{
					"id":25,
					"name":"webhook",
					"description":"",
					"web_url":"http://localhost:3000/root/webhook",
					"avatar_url":null,
					"git_ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
					"git_http_url":"http://localhost:3000/root/webhook.git",
					"namespace":"root",
					"visibility_level":20,
					"path_with_namespace":"root/webhook",
					"default_branch":"master",
					"ci_config_path":null,
					"homepage":"http://localhost:3000/root/webhook",
					"url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
					"ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
					"http_url":"http://localhost:3000/root/webhook.git"
				},
				"target":{
					"id":24,
					"name":"webhook",
					"description":"",
					"web_url":"http://localhost:3000/manland/webhook",
					"avatar_url":null,
					"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"git_http_url":"http://localhost:3000/manland/webhook.git",
					"namespace":"manland",
					"visibility_level":20,
					"path_with_namespace":"manland/webhook",
					"default_branch":"master",
					"ci_config_path":null,
					"homepage":"http://localhost:3000/manland/webhook",
					"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"http_url":"http://localhost:3000/manland/webhook.git"
				},
				"last_commit":{
					"id":"9efcb9dc895dc6dbf19fdd6f2bdbf3e778870eb2",
					"message":"Update README.md",
					"timestamp":"2019-04-01T20:44:25Z",
					"url":"http://localhost:3000/manland/webhook/commit/9efcb9dc895dc6dbf19fdd6f2bdbf3e778870eb2",
					"author":{
						"name":"Administrator",
						"email":"admin@example.com"
					}
				},
				"work_in_progress":false,
				"total_time_spent":0,
				"human_total_time_spent":null,
				"human_time_estimate":null,
				"action":"reopen"
			},
			"labels":[],
			"changes":{
				"state":{
					"previous":"closed",
					"current":"opened"
				},
				"updated_at":{
					"previous":"2019-04-01 21:34:13 UTC",
					"current":"2019-04-01 21:34:16 UTC"
				},
				"total_time_spent":{
					"previous":null,
					"current":0
				}
			},
			"repository":{
				"name":"webhook",
				"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
				"description":"",
				"homepage":"http://localhost:3000/manland/webhook"
			}
		}
		`

const CloseMergeRequestByAssignee = `{
			"object_kind":"merge_request",
			"event_type":"merge_request",
			"user":{
				"name":"manland",
				"username":"manland",
				"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
			},
			"project":{
				"id":24,
				"name":"webhook",
				"description":"",
				"web_url":"http://localhost:3000/manland/webhook",
				"avatar_url":null,
				"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
				"git_http_url":"http://localhost:3000/manland/webhook.git",
				"namespace":"manland",
				"visibility_level":20,
				"path_with_namespace":"manland/webhook",
				"default_branch":"master",
				"ci_config_path":null,
				"homepage":"http://localhost:3000/manland/webhook",
				"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
				"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
				"http_url":"http://localhost:3000/manland/webhook.git"
			},
			"object_attributes":{
				"assignee_id":50,
				"author_id":1,
				"created_at":"2019-04-03 21:07:32 UTC",
				"description":"test open merge request",
				"head_pipeline_id":57,
				"id":35,
				"iid":4,
				"last_edited_at":null,
				"last_edited_by_id":null,
				"merge_commit_sha":null,
				"merge_error":null,
				"merge_params":{
					"force_remove_source_branch":null
				},
				"merge_status":"can_be_merged",
				"merge_user_id":null,
				"merge_when_pipeline_succeeds":false,
				"milestone_id":null,
				"source_branch":"master",
				"source_project_id":25,
				"state":"closed",
				"target_branch":"master",
				"target_project_id":24,
				"time_estimate":0,
				"title":"Master",
				"updated_at":"2019-04-03 21:29:41 UTC",
				"updated_by_id":null,
				"url":"http://localhost:3000/manland/webhook/merge_requests/4",
				"source":{
					"id":25,
					"name":"webhook",
					"description":"",
					"web_url":"http://localhost:3000/root/webhook",
					"avatar_url":null,
					"git_ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
					"git_http_url":"http://localhost:3000/root/webhook.git",
					"namespace":"root",
					"visibility_level":20,
					"path_with_namespace":"root/webhook",
					"default_branch":"master",
					"ci_config_path":null,
					"homepage":"http://localhost:3000/root/webhook",
					"url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
					"ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
					"http_url":"http://localhost:3000/root/webhook.git"
				},
				"target":{
					"id":24,
					"name":"webhook",
					"description":"",
					"web_url":"http://localhost:3000/manland/webhook",
					"avatar_url":null,
					"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"git_http_url":"http://localhost:3000/manland/webhook.git",
					"namespace":"manland",
					"visibility_level":20,
					"path_with_namespace":"manland/webhook",
					"default_branch":"master",
					"ci_config_path":null,
					"homepage":"http://localhost:3000/manland/webhook",
					"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"http_url":"http://localhost:3000/manland/webhook.git"
				},
				"last_commit":{
					"id":"1fd967c14f8265a6056525c343d984ce56472d5c",
					"message":"Update README.md",
					"timestamp":"2019-04-03T21:04:58Z",
					"url":"http://localhost:3000/manland/webhook/commit/1fd967c14f8265a6056525c343d984ce56472d5c",
					"author":{
						"name":"Administrator",
						"email":"admin@example.com"
					}
				},
				"work_in_progress":false,
				"total_time_spent":0,
				"human_total_time_spent":null,
				"human_time_estimate":null,
				"action":"close"
			},
			"labels":[],
			"changes":{
				"state":{
					"previous":"opened",
					"current":"closed"
				},
				"updated_at":{
					"previous":"2019-04-03 21:07:32 UTC",
					"current":"2019-04-03 21:29:41 UTC"
				},
				"assignee":{
					"previous":null,
					"current":{
						"name":"manland",
						"username":"manland",
						"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
					}
				},
				"total_time_spent":{
					"previous":null,
					"current":0
				}
			},
			"repository":{
				"name":"webhook",
				"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
				"description":"",
				"homepage":"http://localhost:3000/manland/webhook"
			},
			"assignee":{
				"name":"manland",
				"username":"manland",
				"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
			}
			}`

const CloseMergeRequestByCreator = `{
				"object_kind":"merge_request",
				"event_type":"merge_request",
				"user":{
					"name":"Administrator",
					"username":"root",
					"avatar_url":"https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
				},
				"project":{
					"id":24,
					"name":"webhook",
					"description":"",
					"web_url":"http://localhost:3000/manland/webhook",
					"avatar_url":null,
					"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"git_http_url":"http://localhost:3000/manland/webhook.git",
					"namespace":"manland",
					"visibility_level":20,
					"path_with_namespace":"manland/webhook",
					"default_branch":"master",
					"ci_config_path":null,
					"homepage":"http://localhost:3000/manland/webhook",
					"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"http_url":"http://localhost:3000/manland/webhook.git"
				},
				"object_attributes":{
					"assignee_id":null,
					"author_id":1,
					"created_at":"2019-04-01 20:25:41 UTC",
					"description":"# Et voila\\n\\nC'est tro pbien",
					"head_pipeline_id":57,
					"id":34,
					"iid":1,
					"last_edited_at":null,
					"last_edited_by_id":null,
					"merge_commit_sha":null,
					"merge_error":null,
					"merge_params":{
						"force_remove_source_branch":null},
						"merge_status":"can_be_merged",
						"merge_user_id":null,
						"merge_when_pipeline_succeeds":false,
						"milestone_id":null,
						"source_branch":"master",
						"source_project_id":25,
						"state":"closed",
						"target_branch":"master",
						"target_project_id":24,
						"time_estimate":0,
						"title":"Update README.md",
						"updated_at":"2019-04-03 21:06:08 UTC",
						"updated_by_id":50,
						"url":"http://localhost:3000/manland/webhook/merge_requests/1",
						"source":{
							"id":25,
							"name":"webhook",
							"description":"",
							"web_url":"http://localhost:3000/root/webhook",
							"avatar_url":null,
							"git_ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
							"git_http_url":"http://localhost:3000/root/webhook.git",
							"namespace":"root",
							"visibility_level":20,
							"path_with_namespace":"root/webhook",
							"default_branch":"master",
							"ci_config_path":null,
							"homepage":"http://localhost:3000/root/webhook",
							"url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
							"ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
							"http_url":"http://localhost:3000/root/webhook.git"
						},
						"target":{
							"id":24,
							"name":"webhook",
							"description":"",
							"web_url":"http://localhost:3000/manland/webhook",
							"avatar_url":null,
							"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
							"git_http_url":"http://localhost:3000/manland/webhook.git",
							"namespace":"manland",
							"visibility_level":20,
							"path_with_namespace":"manland/webhook",
							"default_branch":"master",
							"ci_config_path":null,
							"homepage":"http://localhost:3000/manland/webhook",
							"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
							"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
							"http_url":"http://localhost:3000/manland/webhook.git"
						},
						"last_commit":{
							"id":"1fd967c14f8265a6056525c343d984ce56472d5c",
							"message":"Update README.md",
							"timestamp":"2019-04-03T21:04:58Z",
							"url":"http://localhost:3000/manland/webhook/commit/1fd967c14f8265a6056525c343d984ce56472d5c",
							"author":{
								"name":"Administrator",
								"email":"admin@example.com"
							}
						},
						"work_in_progress":false,
						"total_time_spent":0,
						"human_total_time_spent":null,
						"human_time_estimate":null,
						"action":"close"
					},
					"labels":[],
					"changes":{
						"state":{
							"previous":"opened",
							"current":"closed"
						},
						"updated_at":{
							"previous":"2019-04-03 21:05:03 UTC",
							"current":"2019-04-03 21:06:08 UTC"
						},
						"total_time_spent":{
							"previous":null,
							"current":0}
						},
						"repository":{
							"name":"webhook",
							"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
							"description":"",
							"homepage":"http://localhost:3000/manland/webhook"
						}
					}
					`
const RootUpdateAssigneeMergeRequest = `{
				"object_kind":"merge_request",
				"event_type":"merge_request",
				"user":{
					"name":"Administrator",
					"username":"root",
					"avatar_url":"https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
				},
				"project":{
					"id":24,
					"name":"webhook",
					"namespace":"manland",
					"visibility_level":20,
					"path_with_namespace":"manland/webhook",
					"http_url":"http://localhost:3000/manland/webhook.git"
				},
				"object_attributes":{
					"assignee_id":50,
					"author_id":1,
					"created_at":"2019-04-03 21:07:32 UTC",
					"description":"test open merge request",
					"id":35,
					"iid":4,
					"merge_status":"can_be_merged",
					"state":"opened",
					"title":"Master-2",
					"url":"http://localhost:3000/manland/webhook/merge_requests/4",
					"source":{
						"id":25,
						"name":"webhook",
						"namespace":"root",
						"visibility_level":20,
						"path_with_namespace":"root/webhook",
						"http_url":"http://localhost:3000/root/webhook.git"
					},
					"target":{
						"id":24,
						"name":"webhook",
						"namespace":"manland",
						"visibility_level":20,
						"path_with_namespace":"manland/webhook",
						"http_url":"http://localhost:3000/manland/webhook.git"
					},
					"last_commit":{
						"id":"1fd967c14f8265a6056525c343d984ce56472d5c",
						"message":"Update README.md",
						"timestamp":"2019-04-03T21:04:58Z",
						"url":"http://localhost:3000/manland/webhook/commit/1fd967c14f8265a6056525c343d984ce56472d5c",
						"author":{
							"name":"Administrator",
							"email":"admin@example.com"
						}
					},
					"assignee_ids": [50],
					"action":"update"
				},
				"changes":{
					"assignees":{
						"previous":[
							{
								"id": 100,
								"name": "user",
								"username": "user",
								"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
							}
						],
						"current":[
							{
								"id": 50,
								"name": "manland",
								"username": "manland",
								"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
							}
						]
					}
				},
				"assignees": [
					{
					"id": 50,
					"name": "manland",
					"username": "manland",
					"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
					}
				],
				"repository":{
					"name":"webhook",
					"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"description":"",
					"homepage":"http://localhost:3000/manland/webhook"
				}
				}`

const UserUpdateAssigneeToManlandMergeRequest = `{
				"object_kind":"merge_request",
				"event_type":"merge_request",
				"user":{
					"name":"user",
					"username":"user",
					"avatar_url":"https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
				},
				"project":{
					"id":24,
					"name":"webhook",
					"namespace":"manland",
					"visibility_level":20,
					"path_with_namespace":"manland/webhook",
					"http_url":"http://localhost:3000/manland/webhook.git"
				},
				"object_attributes":{
					"assignee_id":50,
					"author_id":1,
					"created_at":"2019-04-03 21:07:32 UTC",
					"description":"test open merge request",
					"id":35,
					"iid":4,
					"merge_status":"can_be_merged",
					"state":"opened",
					"title":"Master-2",
					"url":"http://localhost:3000/manland/webhook/merge_requests/4",
					"source":{
						"id":25,
						"name":"webhook",
						"namespace":"root",
						"visibility_level":20,
						"path_with_namespace":"root/webhook",
						"http_url":"http://localhost:3000/root/webhook.git"
					},
					"target":{
						"id":24,
						"name":"webhook",
						"namespace":"manland",
						"visibility_level":20,
						"path_with_namespace":"manland/webhook",
						"http_url":"http://localhost:3000/manland/webhook.git"
					},
					"last_commit":{
						"id":"1fd967c14f8265a6056525c343d984ce56472d5c",
						"message":"Update README.md",
						"timestamp":"2019-04-03T21:04:58Z",
						"url":"http://localhost:3000/manland/webhook/commit/1fd967c14f8265a6056525c343d984ce56472d5c",
						"author":{
							"name":"Administrator",
							"email":"admin@example.com"
						}
					},
					"assignee_ids": [50],
					"action":"update"
				},
				"changes":{
					"assignees":{
						"previous":[
							{
								"id": 1,
								"name": "Administrator",
								"username": "root",
								"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
							}
						],
						"current":[
							{
								"id": 50,
								"name": "manland",
								"username": "manland",
								"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
							}
						]
					}
				},
				"assignees": [
					{
					"id": 50,
					"name": "manland",
					"username": "manland",
					"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
					}
				],
				"repository":{
					"name":"webhook",
					"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"description":"",
					"homepage":"http://localhost:3000/manland/webhook"
				}
				}`

const UserUpdateAssigneeToUserMergeRequest = `{
				"object_kind":"merge_request",
				"event_type":"merge_request",
				"user":{
					"name":"user",
					"username":"user",
					"avatar_url":"https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
				},
				"project":{
					"id":24,
					"name":"webhook",
					"namespace":"manland",
					"visibility_level":20,
					"path_with_namespace":"manland/webhook",
					"http_url":"http://localhost:3000/manland/webhook.git"
				},
				"object_attributes":{
					"assignee_id":100,
					"author_id":1,
					"created_at":"2019-04-03 21:07:32 UTC",
					"description":"test open merge request",
					"id":35,
					"iid":4,
					"merge_status":"can_be_merged",
					"state":"opened",
					"title":"Master-2",
					"url":"http://localhost:3000/manland/webhook/merge_requests/4",
					"source":{
						"id":25,
						"name":"webhook",
						"namespace":"root",
						"visibility_level":20,
						"path_with_namespace":"root/webhook",
						"http_url":"http://localhost:3000/root/webhook.git"
					},
					"target":{
						"id":24,
						"name":"webhook",
						"namespace":"manland",
						"visibility_level":20,
						"path_with_namespace":"manland/webhook",
						"http_url":"http://localhost:3000/manland/webhook.git"
					},
					"last_commit":{
						"id":"1fd967c14f8265a6056525c343d984ce56472d5c",
						"message":"Update README.md",
						"timestamp":"2019-04-03T21:04:58Z",
						"url":"http://localhost:3000/manland/webhook/commit/1fd967c14f8265a6056525c343d984ce56472d5c",
						"author":{
							"name":"Administrator",
							"email":"admin@example.com"
						}
					},
					"assignee_ids": [100],
					"action":"update"
				},
				"changes":{
					"assignees":{
						"previous":[
							{
								"id": 50,
								"name": "manland",
								"username": "manland",
								"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
							}
						],
						"current":[
							{
								"id": 100,
								"name": "user",
								"username": "user",
								"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
							}
						]
					}
				},
				"assignees": [
					{
						"id": 100,
						"name": "user",
						"username": "user",
						"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
					}
				],
				"repository":{
					"name":"webhook",
					"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"description":"",
					"homepage":"http://localhost:3000/manland/webhook"
				}
				}`

const RootUpdateReviewerMergeRequest = `{
				"object_kind":"merge_request",
				"event_type":"merge_request",
				"user":{
					"name":"Administrator",
					"username":"root",
					"avatar_url":"https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
				},
				"project":{
					"id":24,
					"name":"webhook",
					"namespace":"manland",
					"visibility_level":20,
					"path_with_namespace":"manland/webhook",
					"http_url":"http://localhost:3000/manland/webhook.git"
				},
				"object_attributes":{
					"assignee_id":50,
					"author_id":1,
					"created_at":"2019-04-03 21:07:32 UTC",
					"description":"test open merge request",
					"id":35,
					"iid":4,
					"merge_status":"can_be_merged",
					"state":"opened",
					"title":"Master-2",
					"url":"http://localhost:3000/manland/webhook/merge_requests/4",
					"source":{
						"id":25,
						"name":"webhook",
						"namespace":"root",
						"visibility_level":20,
						"path_with_namespace":"root/webhook",
						"http_url":"http://localhost:3000/root/webhook.git"
					},
					"target":{
						"id":24,
						"name":"webhook",
						"namespace":"manland",
						"visibility_level":20,
						"path_with_namespace":"manland/webhook",
						"http_url":"http://localhost:3000/manland/webhook.git"
					},
					"last_commit":{
						"id":"1fd967c14f8265a6056525c343d984ce56472d5c",
						"message":"Update README.md",
						"timestamp":"2019-04-03T21:04:58Z",
						"url":"http://localhost:3000/manland/webhook/commit/1fd967c14f8265a6056525c343d984ce56472d5c",
						"author":{
							"name":"Administrator",
							"email":"admin@example.com"
						}
					},
					"reviewer_ids": [50],
					"action":"update"
				},
				"changes":{
					"reviewers":{
						"previous":[
							{
								"id": 100,
								"name": "user",
								"username": "user",
								"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
							}
						],
						"current":[
							{
								"id": 50,
								"name": "manland",
								"username": "manland",
								"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
							}
						]
					}
				},
				"reviewers": [
					{
					"id": 50,
					"name": "manland",
					"username": "manland",
					"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
					}
				],
				"repository":{
					"name":"webhook",
					"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"description":"",
					"homepage":"http://localhost:3000/manland/webhook"
				}
				}`

const MultipleEventsMergeRequest = `{
				"object_kind":"merge_request",
				"event_type":"merge_request",
				"user":{
					"name":"Administrator",
					"username":"root",
					"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
				},
				"project":{
					"id":24,
					"name":"webhook",
					"namespace":"manland",
					"visibility_level":20,
					"path_with_namespace":"manland/webhook",
					"http_url":"http://localhost:3000/manland/webhook.git"
				},
				"object_attributes":{
					"assignee_id":50,
					"author_id":1,
					"created_at":"2019-04-03 21:07:32 UTC",
					"description":"test open merge request",
					"id":35,
					"iid":4,
					"merge_status":"can_be_merged",
					"state":"opened",
					"title":"Master-2",
					"url":"http://localhost:3000/manland/webhook/merge_requests/4",
					"source":{
						"id":25,
						"name":"webhook",
						"namespace":"root",
						"visibility_level":20,
						"path_with_namespace":"root/webhook",
						"http_url":"http://localhost:3000/root/webhook.git"
					},
					"target":{
						"id":24,
						"name":"webhook",
						"namespace":"manland",
						"visibility_level":20,
						"path_with_namespace":"manland/webhook",
						"http_url":"http://localhost:3000/manland/webhook.git"
					},
					"last_commit":{
						"id":"1fd967c14f8265a6056525c343d984ce56472d5c",
						"message":"Update README.md",
						"timestamp":"2019-04-03T21:04:58Z",
						"url":"http://localhost:3000/manland/webhook/commit/1fd967c14f8265a6056525c343d984ce56472d5c",
						"author":{
							"name":"Administrator",
							"email":"admin@example.com"
						}
					},
					"assignee_ids": [50],
					"action":"update"
				},
				"labels":[
					{
						"id": 1,
						"title": "label-1"
					}
				],
				"changes":{
					"title":{
						"previous":"Master",
						"current":"Master-2"
					},
					"description":{
						"previous":"test open merge request",
						"current":"testing open merge request"
					},
					"last_edited_at":{
						"previous":"2019-04-04 21:02:42 UTC",
						"current":"2019-04-04 21:03:46 UTC"
					},
					"updated_at":{
						"previous":"2019-04-04 21:02:42 UTC",
						"current":"2019-04-04 21:03:46 UTC"
					},
					"labels":{
						"previous":[],
						"current":[{
							"id": 1,
							"title": "label-1"
						}]
					}
				},
				"assignees": [
					{
						"id": 50,
						"name": "manland",
						"username": "manland",
						"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\\u0026d=identicon"
					}
				],
				"repository":{
					"name":"webhook",
					"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"description":"",
					"homepage":"http://localhost:3000/manland/webhook"
				}
				}`

const MergeRequestMerged = `{
							"object_kind":"merge_request",
							"event_type":"merge_request",
							"user":{
								"name":"manland",
								"username":"manland",
								"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
							},
							"project":{
								"id":24,
								"name":"webhook",
								"description":"",
								"web_url":"http://localhost:3000/manland/webhook",
								"avatar_url":null,
								"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
								"git_http_url":"http://localhost:3000/manland/webhook.git",
								"namespace":"manland",
								"visibility_level":20,
								"path_with_namespace":"manland/webhook",
								"default_branch":"master",
								"ci_config_path":null,
								"homepage":"http://localhost:3000/manland/webhook",
								"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
								"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
								"http_url":"http://localhost:3000/manland/webhook.git"
							},
							"object_attributes":{
								"assignee_id":50,
								"author_id":1,
								"created_at":"2019-04-03 21:07:32 UTC",
								"description":"test open merge request",
								"head_pipeline_id":57,
								"id":35,
								"iid":4,
								"last_edited_at":null,
								"last_edited_by_id":null,
								"merge_commit_sha":"44213d704986b25e21e2d0f086d0c1503fd659e2",
								"merge_error":null,
								"merge_params":{
									"force_remove_source_branch":null
								},
								"merge_status":"can_be_merged",
								"merge_user_id":null,
								"merge_when_pipeline_succeeds":false,
								"milestone_id":null,
								"source_branch":"master",
								"source_project_id":25,
								"state":"merged",
								"target_branch":"master",
								"target_project_id":24,
								"time_estimate":0,
								"title":"Master",
								"updated_at":"2019-04-04 21:40:41 UTC",
								"updated_by_id":1,
								"url":"http://localhost:3000/manland/webhook/merge_requests/4",
								"source":{
									"id":25,
									"name":"webhook",
									"description":"",
									"web_url":"http://localhost:3000/root/webhook",
									"avatar_url":null,
									"git_ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
									"git_http_url":"http://localhost:3000/root/webhook.git",
									"namespace":"root",
									"visibility_level":20,
									"path_with_namespace":"root/webhook",
									"default_branch":"master",
									"ci_config_path":null,
									"homepage":"http://localhost:3000/root/webhook",
									"url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
									"ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
									"http_url":"http://localhost:3000/root/webhook.git"
								},
								"target":{
									"id":24,
									"name":"webhook",
									"description":"",
									"web_url":"http://localhost:3000/manland/webhook",
									"avatar_url":null,
									"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
									"git_http_url":"http://localhost:3000/manland/webhook.git",
									"namespace":"manland",
									"visibility_level":20,
									"path_with_namespace":"manland/webhook",
									"default_branch":"master",
									"ci_config_path":null,
									"homepage":"http://localhost:3000/manland/webhook",
									"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
									"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
									"http_url":"http://localhost:3000/manland/webhook.git"
								},
								"last_commit":{
									"id":"1fd967c14f8265a6056525c343d984ce56472d5c",
									"message":"Update README.md",
									"timestamp":"2019-04-03T21:04:58Z",
									"url":"http://localhost:3000/manland/webhook/commit/1fd967c14f8265a6056525c343d984ce56472d5c",
									"author":{
										"name":"Administrator",
										"email":"admin@example.com"
									}
								},
								"work_in_progress":false,
								"total_time_spent":0,
								"human_total_time_spent":null,
								"human_time_estimate":null,
								"action":"merge"
							},
							"labels":[],
							"changes":{
								"merge_commit_sha":{
									"previous":null,
									"current":"44213d704986b25e21e2d0f086d0c1503fd659e2"
								},
								"state":{
									"previous":"opened",
									"current":"merged"
								},
								"updated_at":{
									"previous":"2019-04-04 21:15:52 UTC",
									"current":"2019-04-04 21:40:41 UTC"
								},
								"assignee":{
									"previous":null,
									"current":{
										"name":"manland",
										"username":"manland",
										"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
									}
								},
								"total_time_spent":{
									"previous":null,
									"current":0
								}
							},
							"repository":{
								"name":"webhook",
								"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
								"description":"",
								"homepage":"http://localhost:3000/manland/webhook"
							},
							"assignee":{
								"name":"manland",
								"username":"manland",
								"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
							}
							}`

const ApproveMergeRequest = `{
			"object_kind":"merge_request",
			"event_type":"merge_request",
			"user":{
				"name":"manland",
				"username":"manland",
				"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
			},
			"project":{
				"id":24,
				"name":"webhook",
				"description":"",
				"web_url":"http://localhost:3000/manland/webhook",
				"avatar_url":null,
				"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
				"git_http_url":"http://localhost:3000/manland/webhook.git",
				"namespace":"manland",
				"visibility_level":20,
				"path_with_namespace":"manland/webhook",
				"default_branch":"master",
				"ci_config_path":null,
				"homepage":"http://localhost:3000/manland/webhook",
				"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
				"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
				"http_url":"http://localhost:3000/manland/webhook.git"
			},
			"object_attributes":{
				"assignee_id":50,
				"author_id":1,
				"created_at":"2019-04-03 21:07:32 UTC",
				"description":"test open merge request",
				"head_pipeline_id":57,
				"id":35,
				"iid":4,
				"last_edited_at":null,
				"last_edited_by_id":null,
				"merge_commit_sha":null,
				"merge_error":null,
				"merge_params":{
					"force_remove_source_branch":null
				},
				"merge_status":"can_be_merged",
				"merge_user_id":null,
				"merge_when_pipeline_succeeds":false,
				"milestone_id":null,
				"source_branch":"master",
				"source_project_id":25,
				"state":"opened",
				"target_branch":"master",
				"target_project_id":24,
				"time_estimate":0,
				"title":"Master",
				"updated_at":"2019-04-03 21:29:41 UTC",
				"updated_by_id":null,
				"url":"http://localhost:3000/manland/webhook/merge_requests/4",
				"source":{
					"id":25,
					"name":"webhook",
					"description":"",
					"web_url":"http://localhost:3000/root/webhook",
					"avatar_url":null,
					"git_ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
					"git_http_url":"http://localhost:3000/root/webhook.git",
					"namespace":"root",
					"visibility_level":20,
					"path_with_namespace":"root/webhook",
					"default_branch":"master",
					"ci_config_path":null,
					"homepage":"http://localhost:3000/root/webhook",
					"url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
					"ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
					"http_url":"http://localhost:3000/root/webhook.git"
				},
				"target":{
					"id":24,
					"name":"webhook",
					"description":"",
					"web_url":"http://localhost:3000/manland/webhook",
					"avatar_url":null,
					"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"git_http_url":"http://localhost:3000/manland/webhook.git",
					"namespace":"manland",
					"visibility_level":20,
					"path_with_namespace":"manland/webhook",
					"default_branch":"master",
					"ci_config_path":null,
					"homepage":"http://localhost:3000/manland/webhook",
					"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
					"http_url":"http://localhost:3000/manland/webhook.git"
				},
				"last_commit":{
					"id":"1fd967c14f8265a6056525c343d984ce56472d5c",
					"message":"Update README.md",
					"timestamp":"2019-04-03T21:04:58Z",
					"url":"http://localhost:3000/manland/webhook/commit/1fd967c14f8265a6056525c343d984ce56472d5c",
					"author":{
						"name":"Administrator",
						"email":"admin@example.com"
					}
				},
				"work_in_progress":false,
				"total_time_spent":0,
				"human_total_time_spent":null,
				"human_time_estimate":null,
				"action":"approved"
			},
			"labels":[],
			"changes":{
				"updated_at":{
					"previous":"2019-04-03 21:07:32 UTC",
					"current":"2019-04-03 21:29:41 UTC"
				}
			},
			"repository":{
				"name":"webhook",
				"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
				"description":"",
				"homepage":"http://localhost:3000/manland/webhook"
			},
			"assignee":{
				"name":"manland",
				"username":"manland",
				"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
			}
			}`
