package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/slack-go/slack"
	cbpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
)

func TestWriteMessage(t *testing.T) {
	n := new(slackNotifier)
	b := &cbpb.Build{
		ProjectId: "my-project-id",
		Id:        "some-build-id",
		Substitutions: map[string]string{
			"BRANCH_NAME":               "test-deploy",
			"COMMIT_SHA":                "commit-test-test",
			"REF_NAME":                  "test-deploy",
			"REPO_NAME":                 "test-repository",
			"REVISION_ID":               "test-revison",
			"SHORT_SHA":                 "test",
			"TRIGGER_BUILD_CONFIG_PATH": "cloudbuild/cloudbuild.yaml",
			"TRIGGER_NAME":              "triggername",
			"_CLUSTER_NAME":             "sample-cluster",
			"_ENV":                      "sample",
			"_FRONT_DEPLOYMENT_YAML":    "sample.yaml",
			"_REGION":                   "asia-northeast1",
			"_ZONE":                     "asia-northeast1-a",
		},
		Status: cbpb.Build_SUCCESS,
		LogUrl: "https://some.example.com/log/url?foo=bar",
	}

	got, err := n.writeMessage(b)
	if err != nil {
		t.Fatalf("writeMessage failed: %v", err)
	}

	want := &slack.WebhookMessage{
		Attachments: []slack.Attachment{{
			Text: `Successfully deployed to sample environment!! 
			- Environment : sample
			- Branch : test-deploy
			- Deployed Commit : commit-test-test
			- Cluster : sample-cluster
			- Trriger : triggername`,
			Color: "good",
			Actions: []slack.AttachmentAction{{
				Text: "Open Build Details",
				Type: "button",
				URL:  "https://some.example.com/log/url?foo=bar&utm_campaign=google-cloud-build-notifiers&utm_medium=chat&utm_source=google-cloud-build",
			}},
		}},
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("writeMessage got unexpected diff: %s", diff)
	}
}
