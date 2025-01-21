package repos

import "time"

type Query struct {
	Repository Repository `graphql:"repository(owner: $owner, name: $name)"`
}

type Repository struct {
	// Overview
	Id              string `graphql:"id" group:"1⟭ Overview" desc:"The Node ID of the Repository."`
	DatabaseID      int    `graphql:"databaseId" group:"1⟭ Overview" desc:"Identifies the primary key from the database."`
	Name            string `graphql:"name" group:"1⟭ Overview" desc:"The name of the repository."`
	NameWithOwner   string `graphql:"nameWithOwner" group:"1⟭ Overview" desc:"The repository's name with owner (e.g., octocat/hello-world)."`
	Url             string `graphql:"url" group:"1⟭ Overview" desc:"The HTTP URL for this repository."`
	ResourcePath    string `graphql:"resourcePath" group:"1⟭ Overview" desc:"The repository's URL path."`
	Description     string `graphql:"description" group:"1⟭ Overview" desc:"The description of the repository."`
	DescriptionHTML string `graphql:"descriptionHTML" group:"1⟭ Overview" desc:"The description of the repository rendered to HTML."`
	HomepageURL     string `graphql:"homepageUrl" group:"1⟭ Overview" desc:"The repository's URL."`
	PrimaryLanguage struct {
		Name string `name:"Primary Language" group:"1⟭ Overview" desc:"The primary programming language of the repository."`
		// Color string
	} `graphql:"primaryLanguage"`
	LicenseInfo struct {
		Name string `name:"License Name" group:"1⟭ Overview" desc:"The license associated with the repository."`
		// Key  string
	} `graphql:"licenseInfo"`

	// Status
	IsArchived       bool `graphql:"isArchived" group:"2⟭ Status" desc:"Indicates if the repository is archived."`
	IsDisabled       bool `graphql:"isDisabled" group:"2⟭ Status" desc:"Indicates if the repository is disabled."`
	IsEmpty          bool `graphql:"isEmpty" group:"2⟭ Status" desc:"Indicates if the repository is empty."`
	IsFork           bool `graphql:"isFork" group:"2⟭ Status" desc:"Identifies if the repository is a fork."`
	IsInOrganization bool `graphql:"isInOrganization" group:"2⟭ Status" desc:"Indicates if the repository is part of an organization."`
	IsLocked         bool `graphql:"isLocked" group:"2⟭ Status" desc:"Indicates if the repository is locked."`
	IsMirror         bool `graphql:"isMirror" group:"2⟭ Status" desc:"Identifies if the repository is a mirror."`
	IsPrivate        bool `graphql:"isPrivate" group:"2⟭ Status" desc:"Identifies if the repository is private."`
	IsTemplate       bool `graphql:"isTemplate" group:"2⟭ Status" desc:"Indicates if the repository is a template repository."`

	// Metrics
	DiskUsage      int `graphql:"diskUsage" group:"3⟭ Metrics" desc:"The number of kilobytes this repository occupies on disk."`
	ForkCount      int `graphql:"forkCount" group:"3⟭ Metrics" desc:"Returns how many forks there are of this repository in the whole network."`
	StargazerCount int `graphql:"stargazerCount" group:"3⟭ Metrics" desc:"Returns a count of how many stargazers there are on this repository."`
	// WatcherCount   int `graphql:"watcherCount" group:"3⟭ Metrics" desc:"The number of watchers this repository has."`
	// OpenIssueCount       int       `graphql:"openIssueCount" group:"3⟭ Metrics" desc:"The number of open issues in this repository."`
	// OpenPullRequestCount int       `graphql:"openPullRequestCount" group:"3⟭ Metrics" desc:"The number of open pull requests in this repository."`
	CreatedAt time.Time `graphql:"createdAt" group:"3⟭ Metrics" desc:"Identifies the date and time when the object was created."`
	UpdatedAt time.Time `graphql:"updatedAt" group:"3⟭ Metrics" desc:"Identifies the date and time when the object was last updated."`
	PushedAt  time.Time `graphql:"pushedAt" group:"3⟭ Metrics" desc:"Identifies when the repository was last pushed to."`
	Issues    struct {
		TotalCount int `name:"Open Issues" group:"3⟭ Metrics" desc:"The number of open issues for this repository."`
	} `graphql:"issues"`
	Releases struct {
		TotalCount int `name:"Releases" group:"3⟭ Metrics" desc:"The number of releases for this repository."`
	} `graphql:"releases"`
	OpenPullRequests struct {
		TotalCount int `name:"Open Pull Requests" group:"3⟭ Metrics" desc:"The number of open pull requests for this repository."`
	} `graphql:"pullRequests(states: OPEN)"`

	// Repository Features
	HasIssuesEnabled              bool `graphql:"hasIssuesEnabled" group:"4⟭ Features" desc:"Indicates if the repository has issues feature enabled."`
	HasProjectsEnabled            bool `graphql:"hasProjectsEnabled" group:"4⟭ Features" desc:"Indicates if the repository has projects feature enabled."`
	HasWikiEnabled                bool `graphql:"hasWikiEnabled" group:"4⟭ Features" desc:"Indicates if the repository has wiki feature enabled."`
	HasDiscussionsEnabled         bool `graphql:"hasDiscussionsEnabled" group:"4⟭ Features" desc:"Indicates if the repository has discussions feature enabled."`
	HasVulnerabilityAlertsEnabled bool `graphql:"hasVulnerabilityAlertsEnabled" group:"4⟭ Features" desc:"Indicates if the repository has vulnerability alerts enabled."`

	// Merge Settings
	MergeCommitAllowed  bool `graphql:"mergeCommitAllowed" group:"5⟭ Merge" desc:"Whether merge commits are allowed on this repository."`
	RebaseMergeAllowed  bool `graphql:"rebaseMergeAllowed" group:"5⟭ Merge" desc:"Whether rebase-merging is allowed on this repository."`
	SquashMergeAllowed  bool `graphql:"squashMergeAllowed" group:"5⟭ Merge" desc:"Whether squash-merging is allowed on this repository."`
	AutoMergeAllowed    bool `graphql:"autoMergeAllowed" group:"5⟭ Merge" desc:"Whether auto-merge is allowed on this repository."`
	DeleteBranchOnMerge bool `graphql:"deleteBranchOnMerge" group:"5⟭ Merge" desc:"Whether to delete head branches when pull requests are merged."`
	DefaultBranchRef    struct {
		Name string `name:"Default Branch" group:"5⟭ Merge" desc:"The name of the default branch for this repository."`
		// BranchProtectionRule BranchProtectionRule `graphql:"branchProtectionRule"`
	} `graphql:"defaultBranchRef"`

	// Branch Protection
	// BranchProtectionRules BranchProtectionRuleConnection `graphql:"branchProtectionRules" group:"branch" desc:"A list of branch protection rules for this repository."`

	// Security Features
	IsSecurityPolicyEnabled bool   `graphql:"isSecurityPolicyEnabled" group:"6⟭ Security" desc:"Whether this repository has a security policy."`
	SecurityPolicyURL       string `graphql:"securityPolicyUrl" group:"6⟭ Security" desc:"The URL for the security policy of this repository."`
	VulnerabilityAlerts     struct {
		TotalCount int `name:"Vulnerability Alerts" group:"6⟭ Security" desc:"The number of vulnerability alerts for this repository."`
	} `graphql:"vulnerabilityAlerts"`

	// Permissions and Access
	ViewerPermission        string `graphql:"viewerPermission" group:"6⟭ Permissions" desc:"The current user's permission level on the repository (READ, WRITE, ADMIN)."`
	ViewerCanAdminister     bool   `graphql:"viewerCanAdminister" group:"6⟭ Permissions" desc:"Indicates if the current viewer can administer this repository."`
	ViewerCanCreateProjects bool   `graphql:"viewerCanCreateProjects" group:"6⟭ Permissions" desc:"Indicates if the current viewer can create projects in this repository."`
	// ViewerCanFork           bool   `graphql:"viewerCanFork" group:"viewer" desc:"Whether the viewer can fork this repository."`
	ViewerCanSubscribe    bool `graphql:"viewerCanSubscribe" group:"6⟭ Permissions" desc:"Indicates if the current viewer can subscribe to this repository."`
	ViewerCanUpdateTopics bool `graphql:"viewerCanUpdateTopics" group:"6⟭ Permissions" desc:"Indicates if the current viewer can update topics in this repository."`
	ViewerHasStarred      bool `graphql:"viewerHasStarred" group:"6⟭ Permissions" desc:"Indicates if the current viewer has starred this starrable."`

	// ClosedPullRequests struct {
	// 	TotalCount int
	// } `graphql:"pullRequests(states: CLOSED)"`

	// Languages struct {
	// 	Nodes []struct {
	// 		Name  string
	// 		Color string
	// 	} `graphql:"nodes"`
	// } `graphql:"languages(first: 100)"`
}
