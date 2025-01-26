package filters

import (
	"gh-reponark/repo"
)

type FilterMap map[string]Filter

func (FilterMap *FilterMap) FilterRepos(repoConfigs []repo.RepoConfig) []repo.RepoConfig {
	if FilterMap == nil {
		return repoConfigs
	}

	filteredConfigs := []repo.RepoConfig{}

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
