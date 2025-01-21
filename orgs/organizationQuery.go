package orgs

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
		} `graphql:"repositories(first: $first)"`
	} `graphql:"organization(login: $login)"`
}
