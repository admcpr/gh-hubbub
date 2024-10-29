package queries

import "time"

type RepositoryQuery struct {
	Repository Repository `graphql:"repository(owner: $owner, name: $name)"`
}

type Repository struct {
	// Overview
	Id              string `graphql:"id" group:"overview" desc:"The Node ID of the Repository."`
	DatabaseID      int    `graphql:"databaseId" group:"overview" desc:"Identifies the primary key from the database."`
	Name            string `graphql:"name" group:"overview" desc:"The name of the repository."`
	NameWithOwner   string `graphql:"nameWithOwner" group:"overview" desc:"The repository's name with owner (e.g., octocat/hello-world)."`
	Url             string `graphql:"url" group:"overview" desc:"The HTTP URL for this repository."`
	ResourcePath    string `graphql:"resourcePath" group:"overview" desc:"The repository's URL path."`
	Description     string `graphql:"description" group:"overview" desc:"The description of the repository."`
	DescriptionHTML string `graphql:"descriptionHTML" group:"overview" desc:"The description of the repository rendered to HTML."`
	HomepageURL     string `graphql:"homepageUrl" group:"overview" desc:"The repository's URL."`

	// Status
	IsArchived       bool `graphql:"isArchived" group:"status" desc:"Indicates if the repository is archived."`
	IsDisabled       bool `graphql:"isDisabled" group:"status" desc:"Indicates if the repository is disabled."`
	IsEmpty          bool `graphql:"isEmpty" group:"status" desc:"Indicates if the repository is empty."`
	IsFork           bool `graphql:"isFork" group:"status" desc:"Identifies if the repository is a fork."`
	IsInOrganization bool `graphql:"isInOrganization" group:"status" desc:"Indicates if the repository is part of an organization."`
	IsLocked         bool `graphql:"isLocked" group:"status" desc:"Indicates if the repository is locked."`
	IsMirror         bool `graphql:"isMirror" group:"status" desc:"Identifies if the repository is a mirror."`
	IsPrivate        bool `graphql:"isPrivate" group:"status" desc:"Identifies if the repository is private."`
	IsTemplate       bool `graphql:"isTemplate" group:"status" desc:"Indicates if the repository is a template repository."`

	// Metrics
	DiskUsage      int `graphql:"diskUsage" group:"metrics" desc:"The number of kilobytes this repository occupies on disk."`
	ForkCount      int `graphql:"forkCount" group:"metrics" desc:"Returns how many forks there are of this repository in the whole network."`
	StargazerCount int `graphql:"stargazerCount" group:"metrics" desc:"Returns a count of how many stargazers there are on this repository."`
	// WatcherCount   int `graphql:"watcherCount" group:"metrics" desc:"The number of watchers this repository has."`
	// OpenIssueCount       int       `graphql:"openIssueCount" group:"metrics" desc:"The number of open issues in this repository."`
	// OpenPullRequestCount int       `graphql:"openPullRequestCount" group:"metrics" desc:"The number of open pull requests in this repository."`
	CreatedAt time.Time `graphql:"createdAt" group:"metrics" desc:"Identifies the date and time when the object was created."`
	UpdatedAt time.Time `graphql:"updatedAt" group:"metrics" desc:"Identifies the date and time when the object was last updated."`
	PushedAt  time.Time `graphql:"pushedAt" group:"metrics" desc:"Identifies when the repository was last pushed to."`

	// Repository Features
	HasIssuesEnabled              bool `graphql:"hasIssuesEnabled" group:"features" desc:"Indicates if the repository has issues feature enabled."`
	HasProjectsEnabled            bool `graphql:"hasProjectsEnabled" group:"features" desc:"Indicates if the repository has projects feature enabled."`
	HasWikiEnabled                bool `graphql:"hasWikiEnabled" group:"features" desc:"Indicates if the repository has wiki feature enabled."`
	HasDiscussionsEnabled         bool `graphql:"hasDiscussionsEnabled" group:"features" desc:"Indicates if the repository has discussions feature enabled."`
	HasVulnerabilityAlertsEnabled bool `graphql:"hasVulnerabilityAlertsEnabled" group:"features" desc:"Indicates if the repository has vulnerability alerts enabled."`

	// Merge Settings
	MergeCommitAllowed  bool `graphql:"mergeCommitAllowed" group:"merge" desc:"Whether merge commits are allowed on this repository."`
	RebaseMergeAllowed  bool `graphql:"rebaseMergeAllowed" group:"merge" desc:"Whether rebase-merging is allowed on this repository."`
	SquashMergeAllowed  bool `graphql:"squashMergeAllowed" group:"merge" desc:"Whether squash-merging is allowed on this repository."`
	AutoMergeAllowed    bool `graphql:"autoMergeAllowed" group:"merge" desc:"Whether auto-merge is allowed on this repository."`
	DeleteBranchOnMerge bool `graphql:"deleteBranchOnMerge" group:"merge" desc:"Whether to delete head branches when pull requests are merged."`

	// Branch Protection
	// BranchProtectionRules BranchProtectionRuleConnection `graphql:"branchProtectionRules" group:"branch" desc:"A list of branch protection rules for this repository."`

	// Security Features
	IsSecurityPolicyEnabled bool   `graphql:"isSecurityPolicyEnabled" group:"security" desc:"Whether this repository has a security policy."`
	SecurityPolicyURL       string `graphql:"securityPolicyUrl" group:"security" desc:"The URL for the security policy of this repository."`

	// Permissions and Access
	ViewerPermission        string `graphql:"viewerPermission" group:"permissions" desc:"The current user's permission level on the repository (READ, WRITE, ADMIN)."`
	ViewerCanAdminister     bool   `graphql:"viewerCanAdminister" group:"permissions" desc:"Indicates if the current viewer can administer this repository."`
	ViewerCanCreateProjects bool   `graphql:"viewerCanCreateProjects" group:"permissions" desc:"Indicates if the current viewer can create projects in this repository."`
	// ViewerCanFork           bool   `graphql:"viewerCanFork" group:"viewer" desc:"Whether the viewer can fork this repository."`
	ViewerCanSubscribe    bool `graphql:"viewerCanSubscribe" group:"permissions" desc:"Indicates if the current viewer can subscribe to this repository."`
	ViewerCanUpdateTopics bool `graphql:"viewerCanUpdateTopics" group:"permissions" desc:"Indicates if the current viewer can update topics in this repository."`
	ViewerHasStarred      bool `graphql:"viewerHasStarred" group:"permissions" desc:"Indicates if the current viewer has starred this starrable."`

	DefaultBranchRef struct {
		Name                 string
		BranchProtectionRule BranchProtectionRule `graphql:"branchProtectionRule"`
	} `graphql:"defaultBranchRef"`
	VulnerabilityAlerts struct {
		TotalCount int
	} `graphql:"vulnerabilityAlerts"`
	OpenPullRequests struct {
		TotalCount int
	} `graphql:"pullRequests(states: OPEN)"`
	// ClosedPullRequests struct {
	// 	TotalCount int
	// } `graphql:"pullRequests(states: CLOSED)"`
	PrimaryLanguage struct {
		Name  string
		Color string
	} `graphql:"primaryLanguage"`
	Languages struct {
		Nodes []struct {
			Name  string
			Color string
		} `graphql:"nodes"`
	} `graphql:"languages(first: 100)"`
	LicenseInfo struct {
		Name string
		Key  string
	} `graphql:"licenseInfo"`
	Issues struct {
		TotalCount int
	} `graphql:"issues"`
	Releases struct {
		TotalCount int
	} `graphql:"releases"`
}
