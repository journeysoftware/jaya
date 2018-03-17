// Package github contains utility functions for building Slack messages from Github notifications
package github

import (
	"encoding/json"
	"fmt"
)

// IssuesPayload ...
type IssuesPayload struct {
	Action string `json:"action"`
	Issue  struct {
		URL         string `json:"url"`
		LabelsURL   string `json:"labels_url"`
		CommentsURL string `json:"comments_url"`
		EventsURL   string `json:"events_url"`
		HTMLURL     string `json:"html_url"`
		ID          int64  `json:"id, string, omitempty"`
		Number      int64  `json:"number, string, omitempty"`
		Title       string `json:"title"`
		User        struct {
			Login             string `json:"login"`
			ID                int64  `json:"id, string, omitempty"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedeventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			Siteadmin         bool   `json:"site_admin, string, omitempty"`
		}
		Labels []struct {
			ID      int64  `json:"id, string, omitempty"`
			URL     string `json:"url"`
			Name    string `json:"name"`
			Color   string `json:"color"`
			Default bool   `json:"default, string, omitempty"`
		}
		State     string `json:"state"`
		Locked    bool   `json:"locked, string, omitempty"`
		Assignee  string `json:"assignee"`
		Milestone string `json:"milestone"`
		Comments  int64  `json:"comments, string, omitempty"`
		Createdat string `json:"created_at"`
		Updatedat string `json:"updated_at"`
		Closedat  string `json:"closed_at"`
		Body      string `json:"body"`
	}
	Repository struct {
		ID       int64  `json:"id, string, omitempty"`
		Name     string `json:"name"`
		Fullname string `json:"full_name"`
		Owner    struct {
			Login             string `json:"login"`
			ID                int64  `json:"id, string, omitempty"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedeventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			Siteadmin         bool   `json:"site_admin, string, omitempty"`
		}
		Private          bool   `json:"private, string, omitempty"`
		HTMLURL          string `json:"html_url"`
		Description      string `json:"description"`
		Fork             bool   `json:"fork, string, omitempty"`
		URL              string `json:"url"`
		ForksURL         string `json:"forks_url"`
		KeysURL          string `json:"keys_url"`
		CollaboratorsURL string `json:"collaborators_url"`
		TeamsURL         string `json:"teams_url"`
		HooksURL         string `json:"hooks_url"`
		IssueeventsURL   string `json:"issue_events_url"`
		EventsURL        string `json:"events_url"`
		AssigneesURL     string `json:"assignees_url"`
		BranchesURL      string `json:"branches_url"`
		TagsURL          string `json:"tags_url"`
		BlobsURL         string `json:"blobs_url"`
		GittagsURL       string `json:"git_tags_url"`
		GitrefsURL       string `json:"git_refs_url"`
		TreesURL         string `json:"trees_url"`
		StatusesURL      string `json:"statuses_url"`
		LanguagesURL     string `json:"languages_url"`
		StargazersURL    string `json:"stargazers_url"`
		ContributorsURL  string `json:"contributors_url"`
		SubscribersURL   string `json:"subscribers_url"`
		SubscriptionURL  string `json:"subscription_url"`
		CommitsURL       string `json:"commits_url"`
		GitcommitsURL    string `json:"git_commits_url"`
		CommentsURL      string `json:"comments_url"`
		IssuecommentURL  string `json:"issue_comment_url"`
		ContentsURL      string `json:"contents_url"`
		CompareURL       string `json:"compare_url"`
		MergesURL        string `json:"merges_url"`
		ArchiveURL       string `json:"archive_url"`
		DownloadsURL     string `json:"downloads_url"`
		IssuesURL        string `json:"issues_url"`
		PullsURL         string `json:"pulls_url"`
		MilestonesURL    string `json:"milestones_url"`
		NotificationsURL string `json:"notifications_url"`
		LabelsURL        string `json:"labels_url"`
		ReleasesURL      string `json:"releases_url"`
		Createdat        string `json:"created_at"`
		Updatedat        string `json:"updated_at"`
		Pushedat         string `json:"pushed_at"`
		GitURL           string `json:"git_url"`
		SSHURL           string `json:"ssh_url"`
		CloneURL         string `json:"clone_url"`
		SvnURL           string `json:"svn_url"`
		Homepage         string `json:"homepage"`
		Size             int64  `json:"size, string, omitempty"`
		Stargazerscount  int64  `json:"stargazers_count, string, omitempty"`
		Watcherscount    int64  `json:"watchers_count, string, omitempty"`
		Language         string `json:"language"`
		Hasissues        bool   `json:"has_issues, string, omitempty"`
		Hasdownloads     bool   `json:"has_downloads, string, omitempty"`
		Haswiki          bool   `json:"has_wiki, string, omitempty"`
		Haspages         bool   `json:"has_pages, string, omitempty"`
		Forkscount       int64  `json:"forks_count, string, omitempty"`
		MirrorURL        string `json:"mirror_url"`
		Openissuescount  int64  `json:"open_issues_count, string, omitempty"`
		Forks            int64  `json:"forks, string, omitempty"`
		Openissues       int64  `json:"open_issues, string, omitempty"`
		Watchers         int64  `json:"watchers, string, omitempty"`
		Defaultbranch    string `json:"default_branch"`
	}
	Sender struct {
		Login             string `json:"login"`
		ID                int64  `json:"id, string, omitempty"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedeventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		Siteadmin         bool   `json:"site_admin, string, omitempty"`
	}
}

//Issues ...
func Issues(body []byte) string {
	var m = IssuesPayload{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println("error:", err)
		return "error"
	}
	action := m.Action
	name := m.Issue.User.Login
	repo := m.Repository.Name
	message := name + " has " + action + " an issue for " + repo + "."
	return message
}
