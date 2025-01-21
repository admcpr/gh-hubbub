package filters

import (
	"gh-hubbub/repos"
)

type FiltersMsg FilterMap

func (FilterMap *FilterMap) FilterRepos(repos []repos.RepoProperties) []repos.RepoProperties {
	if FilterMap == nil {
		return repos
	}

	filteredRepos := []repos.RepoProperties{}
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
