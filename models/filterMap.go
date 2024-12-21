package models

import "gh-hubbub/structs"

type filtersMsg filterMap

func (filterMap *filterMap) filterRepos(repos []structs.RepoProperties) []structs.RepoProperties {
	if filterMap == nil {
		return repos
	}

	filteredRepos := []structs.RepoProperties{}
	for _, repo := range repos {
		matches := true
		for _, filter := range *filterMap {
			if !filter.Matches(repo.Properties[filter.GetName()]) {
				matches = false
				break
			}
		}
		if matches {
			filteredRepos = append(filteredRepos, repo)
		}
	}

	return filteredRepos
}
