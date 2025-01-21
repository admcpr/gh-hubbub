package filters

import "gh-hubbub/structs"

type FiltersMsg FilterMap

func (FilterMap *FilterMap) FilterRepos(repos []structs.RepoProperties) []structs.RepoProperties {
	if FilterMap == nil {
		return repos
	}

	filteredRepos := []structs.RepoProperties{}
	for _, repo := range repos {
		matches := true
		for _, filter := range *FilterMap {
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
