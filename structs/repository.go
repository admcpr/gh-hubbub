package structs

import (
	"fmt"
	"reflect"
	"time"
)

type Setting struct {
	Name  string
	Value interface{}
	Type  reflect.Type
}

func (s Setting) String() string {
	formattedValue := ""

	switch value := s.Value.(type) {
	case bool:
		formattedValue = YesNo(value)
	case string:
		formattedValue = value
	case time.Time:
		formattedValue = value.Format("2006/01/02")
	case int:
		formattedValue = fmt.Sprint(value)
	}

	return formattedValue
}

func NewSetting(name string, value interface{}) Setting {
	return Setting{Name: name, Value: value, Type: reflect.TypeOf(value)}
}

type RepositorySettings struct {
	Name         string
	Url          string
	SettingsTabs []SettingsTab
}

type SettingsTab struct {
	Name     string
	Settings []Setting
}

func NewRepository(r Repository) RepositorySettings {
	rule := r.DefaultBranchRef.BranchProtectionRule

	return RepositorySettings{
		Name: r.Name,
		Url:  r.Url,
		SettingsTabs: []SettingsTab{
			{
				Name: "Overview",
				Settings: []Setting{
					NewSetting("Private", r.IsPrivate),
					NewSetting("Template", r.IsTemplate),
					NewSetting("Archived", r.IsArchived),
					NewSetting("Disabled", r.IsDisabled),
					NewSetting("Fork", r.IsFork),
					NewSetting("Last updated", r.UpdatedAt),
					NewSetting("Stars", r.StargazerCount),
					NewSetting("Wiki", r.HasWikiEnabled),
					NewSetting("Issues", r.HasIssuesEnabled),
					NewSetting("Projects", r.HasProjectsEnabled),
					NewSetting("Discussions", r.HasDiscussionsEnabled),
				},
			},
			{
				Name: "Pull Requests",
				Settings: []Setting{
					NewSetting("Allow merge commits", r.MergeCommitAllowed),
					NewSetting("Allow squash merging", r.SquashMergeAllowed),
					NewSetting("Allow rebase merging", r.RebaseMergeAllowed),
					NewSetting("Allow auto-merge", r.AutoMergeAllowed),
					NewSetting("Automatically delete head branches", r.DeleteBranchOnMerge),
					NewSetting("Open pull requests", r.PullRequests.TotalCount),
				},
			},
			{
				Name: "Default Branch",
				Settings: []Setting{
					NewSetting("Name", r.DefaultBranchRef.Name),
					NewSetting("Require approving reviews", rule.RequiresApprovingReviews),
					NewSetting("Number of approvals required", rule.RequiredApprovingReviewCount),
					NewSetting("Dismiss stale requests", rule.DismissesStaleReviews),
					NewSetting("Require review from Code Owners", rule.RequiresCodeOwnerReviews),
					NewSetting("Restrict who can dismiss pull request reviews", rule.RestrictsReviewDismissals),
					NewSetting("Require approval of the most recent reviewable push", rule.RequireLastPushApproval),
					NewSetting("Require status checks to pass before merging", rule.RequiresStatusChecks),
					NewSetting("Require conversation resolution before merging", rule.RequiresConversationResolution),
					NewSetting("Requires signed commits", rule.RequiresCommitSignatures),
					NewSetting("Require linear history", rule.RequiresLinearHistory),
					NewSetting("Require deployments to succeed before merging", rule.RequiresDeployments),
					NewSetting("Lock branch", rule.LockBranch),
					NewSetting("Do not allow bypassing the above settings", rule.IsAdminEnforced),
					NewSetting("Restrict who can push to matching branches", rule.RestrictsPushes),
					NewSetting("Allow force pushes", rule.AllowsForcePushes),
					NewSetting("Allow deletions", rule.AllowsDeletions),
				},
			},
			{
				Name: "Security",
				Settings: []Setting{
					NewSetting("Vulnerability alerts enabled", r.HasVulnerabilityAlertsEnabled),
					NewSetting("Vulnerability alert count", r.VulnerabilityAlerts.TotalCount),
				},
			},
		},
	}
}
