// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package webhook

const NewIssue = `{
	"object_kind":"issue",
	"event_type":"issue",
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
		"author_id":1,
		"closed_at":null,
		"confidential":false,
		"created_at":"2019-04-06 21:03:04 UTC",
		"description":"hello world!",
		"due_date":null,
		"id":181,
		"iid":1,
		"last_edited_at":null,
		"last_edited_by_id":null,
		"milestone_id":null,
		"moved_to_id":null,
		"project_id":24,
		"relative_position":1073742323,
		"state":"opened",
		"time_estimate":0,
		"title":"test new issue",
		"updated_at":"2019-04-06 21:03:04 UTC",
		"updated_by_id":null,
		"url":"http://localhost:3000/manland/webhook/issues/1",
		"total_time_spent":0,
		"human_total_time_spent":null,
		"human_time_estimate":null,
		"assignee_ids":[50],
		"assignee_id":50,
		"action":"open"
	},
	"labels":[],
	"changes":{
		"author_id":{
			"previous":null,
			"current":1
		},
		"created_at":{
			"previous":null,
			"current":"2019-04-06 21:03:04 UTC"
		},
		"description":{
			"previous":null,
			"current":"hello world!"
		},
		"id":{
			"previous":null,
			"current":181
		},
		"iid":{
			"previous":null,
			"current":1
		},
		"project_id":{
			"previous":null,
			"current":24
		},
		"relative_position":{
			"previous":null,
			"current":1073742323
		},
		"state":{
			"previous":null,
			"current":"opened"
		},
		"title":{
			"previous":null,
			"current":"test new issue"
		},
		"updated_at":{
			"previous":null,
			"current":"2019-04-06 21:03:04 UTC"
		},
		"assignees":{
			"previous":[],
			"current":[{
				"name":"manland",
				"username":"manland",
				"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
				}]
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
		"assignees":[{
			"name":"manland",
			"username":"manland",
			"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
			}]
			}`

const NewIssueUnassigned = `{
				"object_kind":"issue",
				"event_type":"issue",
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
					"author_id":1,
					"closed_at":null,
					"confidential":false,
					"created_at":"2019-04-06 21:13:03 UTC",
					"description":"Hello world",
					"due_date":null,
					"id":182,
					"iid":2,
					"last_edited_at":null,
					"last_edited_by_id":null,
					"milestone_id":null,
					"moved_to_id":null,
					"project_id":24,
					"relative_position":1073742823,
					"state":"opened",
					"time_estimate":0,
					"title":"new issue",
					"updated_at":"2019-04-06 21:13:03 UTC",
					"updated_by_id":null,
					"url":"http://localhost:3000/manland/webhook/issues/2",
					"total_time_spent":0,
					"human_total_time_spent":null,
					"human_time_estimate":null,
					"assignee_ids":[],
					"assignee_id":null,
					"action":"open"
				},
				"labels":[],
				"changes":{
					"author_id":{
						"previous":null,
						"current":1
					},
					"created_at":{
						"previous":null,
						"current":"2019-04-06 21:13:03 UTC"
					},
					"description":{
						"previous":null,
						"current":"Hello world"
					},
					"id":{
						"previous":null,
						"current":182
					},
					"iid":{
						"previous":null,
						"current":2
					},
					"project_id":{
						"previous":null,
						"current":24
					},
					"relative_position":{
						"previous":null,
						"current":1073742823
					},
					"state":{
						"previous":null,
						"current":"opened"
					},
					"title":{
						"previous":null,
						"current":"new issue"
					},
					"updated_at":{
						"previous":null,
						"current":"2019-04-06 21:13:03 UTC"
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
				}`

const CloseIssue = `{
					"object_kind":"issue",
					"event_type":"issue",
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
						"author_id":1,
						"closed_at":"2019-04-06 21:08:00 UTC",
						"confidential":false,
						"created_at":"2019-04-06 21:03:04 UTC",
						"description":"hello world!",
						"due_date":null,
						"id":181,
						"iid":1,
						"last_edited_at":null,
						"last_edited_by_id":null,
						"milestone_id":null,
						"moved_to_id":null,
						"project_id":24,
						"relative_position":1073742323,
						"state":"closed",
						"time_estimate":0,
						"title":"test new issue",
						"updated_at":"2019-04-06 21:08:00 UTC",
						"updated_by_id":null,
						"url":"http://localhost:3000/manland/webhook/issues/1",
						"total_time_spent":0,
						"human_total_time_spent":null,
						"human_time_estimate":null,
						"assignee_ids":[50],
						"assignee_id":50,
						"action":"close"
					},
					"labels":[],
					"changes":{
						"updated_at":{
							"previous":"2019-04-06 21:08:00 UTC",
							"current":"2019-04-06 21:08:00 UTC"
						},
						"assignees":{
							"previous":[],
							"current":[{
								"name":"manland",
								"username":"manland",
								"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
								}]
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
						"assignees":[{
							"name":"manland",
							"username":"manland",
							"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
							}]
							}`

const ReopenIssue = `{
											"object_kind":"issue",
											"event_type":"issue",
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
												"author_id":1,
												"closed_at":null,
												"confidential":false,
												"created_at":"2019-04-06 21:03:04 UTC",
												"description":"hello world!",
												"due_date":null,
												"id":181,
												"iid":1,
												"last_edited_at":null,
												"last_edited_by_id":null,
												"milestone_id":null,
												"moved_to_id":null,
												"project_id":24,
												"relative_position":1073742323,
												"state":"opened",
												"time_estimate":0,
												"title":"test new issue",
												"updated_at":"2019-04-09 21:36:00 UTC",
												"updated_by_id":null,
												"url":"http://localhost:3000/manland/webhook/issues/1",
												"total_time_spent":0,
												"human_total_time_spent":null,
												"human_time_estimate":null,
												"assignee_ids":[50],
												"assignee_id":50,
												"action":"reopen"
											},
											"labels":[],
											"changes":{
												"closed_at":{
													"previous":"2019-04-06 21:08:00 UTC",
													"current":null
												},
												"state":{
													"previous":"closed",
													"current":"opened"
												},
												"updated_at":{
													"previous":"2019-04-06 21:08:00 UTC",
													"current":"2019-04-09 21:36:00 UTC"
												},
												"assignees":{
													"previous":[],
													"current":[{
														"name":"manland",
														"username":"manland",
														"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
														}]
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
												"assignees":[{
													"name":"manland",
													"username":"manland",
													"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
													}]
													}`
