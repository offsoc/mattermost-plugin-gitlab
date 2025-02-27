// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package webhook

const JobPending = `{
	"object_kind": "build",
	"ref": "gitlab-script-trigger",
	"tag": false,
	"before_sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
	"sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
	"build_id": 1977,
	"build_name": "test",
	"build_stage": "test",
	"build_status": "pending",
	"build_created_at": "2021-02-23T02:41:37.886Z",
	"build_started_at": null,
	"build_finished_at": null,
	"build_duration": null,
	"build_allow_failure": false,
	"build_failure_reason": "script_failure",
	"pipeline_id": 2366,
	"project_id": 380,
	"project_name": "gitlab-org/gitlab-test",
	"user": {
	  "id": 3,
	  "name": "User",
	  "email": "user@gitlab.com",
	  "avatar_url": "http://www.gravatar.com/avatar/e32bd13e2add097461cb96824b7a829c?s=80\u0026d=identicon"
	},
	"commit": {
	  "id": 2366,
	  "sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
	  "message": "test\n",
	  "author_name": "User",
	  "author_email": "user@gitlab.com",
	  "status": "pending",
	  "duration": null,
	  "started_at": null,
	  "finished_at": null
	},
	"repository": {
	  "name": "gitlab_test",
	  "description": "Atque in sunt eos similique dolores voluptatem.",
	  "homepage": "http://192.168.64.1:3005/gitlab-org/gitlab-test",
	  "git_ssh_url": "git@192.168.64.1:gitlab-org/gitlab-test.git",
	  "git_http_url": "http://192.168.64.1:3005/gitlab-org/gitlab-test.git",
	  "visibility_level": 20
	},
	"runner": {
	  "active": true,
	  "runner_type": "project_type",
	  "is_shared": false,
	  "id": 380987,
	  "description": "shared-runners-manager-6.gitlab.com",
	  "tags": [
		"linux",
		"docker"
	  ]
	},
	"environment": null
  }`

const JobRunning = `{
	"object_kind": "build",
	"ref": "gitlab-script-trigger",
	"tag": false,
	"before_sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
	"sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
	"build_id": 1977,
	"build_name": "test",
	"build_stage": "test",
	"build_status": "running",
	"build_created_at": "2021-02-23T02:41:37.886Z",
	"build_started_at": null,
	"build_finished_at": null,
	"build_duration": null,
	"build_allow_failure": false,
	"build_failure_reason": "script_failure",
	"pipeline_id": 2366,
	"project_id": 380,
	"project_name": "gitlab-org/gitlab-test",
	"user": {
	  "id": 3,
	  "name": "User",
	  "email": "user@gitlab.com",
	  "avatar_url": "http://www.gravatar.com/avatar/e32bd13e2add097461cb96824b7a829c?s=80\u0026d=identicon"
	},
	"commit": {
	  "id": 2366,
	  "sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
	  "message": "test\n",
	  "author_name": "User",
	  "author_email": "user@gitlab.com",
	  "status": "running",
	  "duration": null,
	  "started_at": null,
	  "finished_at": null
	},
	"repository": {
	  "name": "gitlab_test",
	  "description": "Atque in sunt eos similique dolores voluptatem.",
	  "homepage": "http://192.168.64.1:3005/gitlab-org/gitlab-test",
	  "git_ssh_url": "git@192.168.64.1:gitlab-org/gitlab-test.git",
	  "git_http_url": "http://192.168.64.1:3005/gitlab-org/gitlab-test.git",
	  "visibility_level": 20
	},
	"runner": {
	  "active": true,
	  "runner_type": "project_type",
	  "is_shared": false,
	  "id": 380987,
	  "description": "shared-runners-manager-6.gitlab.com",
	  "tags": [
		"linux",
		"docker"
	  ]
	},
	"environment": null
  }`

const JobSuccess = `{
	"object_kind": "build",
	"ref": "gitlab-script-trigger",
	"tag": false,
	"before_sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
	"sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
	"build_id": 1977,
	"build_name": "test",
	"build_stage": "test",
	"build_status": "success",
	"build_created_at": "2021-02-23T02:41:37.886Z",
	"build_started_at": null,
	"build_finished_at": null,
	"build_duration": null,
	"build_allow_failure": false,
	"build_failure_reason": "script_failure",
	"pipeline_id": 2366,
	"project_id": 380,
	"project_name": "gitlab-org/gitlab-test",
	"user": {
	  "id": 3,
	  "name": "User",
	  "email": "user@gitlab.com",
	  "avatar_url": "http://www.gravatar.com/avatar/e32bd13e2add097461cb96824b7a829c?s=80\u0026d=identicon"
	},
	"commit": {
	  "id": 2366,
	  "sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
	  "message": "test\n",
	  "author_name": "User",
	  "author_email": "user@gitlab.com",
	  "status": "success",
	  "duration": null,
	  "started_at": null,
	  "finished_at": null
	},
	"repository": {
	  "name": "gitlab_test",
	  "description": "Atque in sunt eos similique dolores voluptatem.",
	  "homepage": "http://192.168.64.1:3005/gitlab-org/gitlab-test",
	  "git_ssh_url": "git@192.168.64.1:gitlab-org/gitlab-test.git",
	  "git_http_url": "http://192.168.64.1:3005/gitlab-org/gitlab-test.git",
	  "visibility_level": 20
	},
	"runner": {
	  "active": true,
	  "runner_type": "project_type",
	  "is_shared": false,
	  "id": 380987,
	  "description": "shared-runners-manager-6.gitlab.com",
	  "tags": [
		"linux",
		"docker"
	  ]
	},
	"environment": null
  }`

const JobFailed = `{
	  "object_kind": "build",
	  "ref": "gitlab-script-trigger",
	  "tag": false,
	  "before_sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
	  "sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
	  "build_id": 1977,
	  "build_name": "test",
	  "build_stage": "test",
	  "build_status": "failed",
	  "build_created_at": "2021-02-23T02:41:37.886Z",
	  "build_started_at": null,
	  "build_finished_at": null,
	  "build_duration": null,
	  "build_allow_failure": false,
	  "build_failure_reason": "script_failure",
	  "pipeline_id": 2366,
	  "project_id": 380,
	  "project_name": "gitlab-org/gitlab-test",
	  "user": {
		"id": 3,
		"name": "User",
		"email": "user@gitlab.com",
		"avatar_url": "http://www.gravatar.com/avatar/e32bd13e2add097461cb96824b7a829c?s=80\u0026d=identicon"
	  },
	  "commit": {
		"id": 2366,
		"sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
		"message": "test\n",
		"author_name": "User",
		"author_email": "user@gitlab.com",
		"status": "failed",
		"duration": null,
		"started_at": null,
		"finished_at": null
	  },
	  "repository": {
		"name": "gitlab_test",
		"description": "Atque in sunt eos similique dolores voluptatem.",
		"homepage": "http://192.168.64.1:3005/gitlab-org/gitlab-test",
		"git_ssh_url": "git@192.168.64.1:gitlab-org/gitlab-test.git",
		"git_http_url": "http://192.168.64.1:3005/gitlab-org/gitlab-test.git",
		"visibility_level": 20
	  },
	  "runner": {
		"active": true,
		"runner_type": "project_type",
		"is_shared": false,
		"id": 380987,
		"description": "shared-runners-manager-6.gitlab.com",
		"tags": [
		  "linux",
		  "docker"
		]
	  },
	  "environment": null
	}`
