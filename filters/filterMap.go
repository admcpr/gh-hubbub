package filters

import (
	"gh-hubbub/repos"
)

type FilterMap map[string]Filter

func (FilterMap *FilterMap) FilterRepos(repoConfigs []repos.RepoConfig) []repos.RepoConfig {
	if FilterMap == nil {
		return repoConfigs
	}

	filteredConfigs := []repos.RepoConfig{}

	for _, repo := range repoConfigs {
		matches := true
		for _, filter := range *FilterMap {
			if !filter.Matches(repo.Properties[filter.GetName()]) {
				matches = false
				break
			}
		}
		if matches {
			filteredConfigs = append(filteredConfigs, repo)
		}
	}

	return filteredConfigs
}
