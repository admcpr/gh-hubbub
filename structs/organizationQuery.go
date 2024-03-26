package structs

type OrganizationQuery struct {
	Organization struct {
		Id           string
		Login        string
		Url          string
		Repositories struct {
			Nodes []struct {
				Name string
				Url  string
			} `graphql:"nodes"`
			// Node RepositoryQuery `graphql:"node"`
		} `graphql:"repositories(first: $first)"`
	} `graphql:"organization(login: $login)"`
}
