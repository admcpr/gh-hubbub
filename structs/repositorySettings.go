package structs

import (
	"gh-hubbub/queries"
)

type RepositorySettings struct {
	Name         string
	Url          string
	SettingsTabs []SettingsTab
}

type SettingsTab struct {
	Name     string
	Settings []Setting
}

func NewRepository(r queries.Repository) RepositorySettings {
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
					NewSetting("Open pull requests", r.OpenPullRequests.TotalCount),
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

func GetListItem(repo RepositorySettings) ListItem {
	return NewListItem(repo.Name, repo.Url)
}
