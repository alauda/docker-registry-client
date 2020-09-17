package registry

import "strings"

type repositoriesResponse struct {
	Repositories []string `json:"repositories"`
}

func (registry *Registry) Repositories() ([]string, error) {
	url := registry.url("/v2/_catalog")
	repos := make([]string, 0, 10)
	var err error //We create this here, otherwise url will be rescoped with :=
	var response repositoriesResponse
	for {
		if !strings.HasPrefix(url, "http") {
			url = registry.url(url)
		}
		registry.Logf("registry.repositories url=%s", url)
		url, err = registry.getPaginatedJSON(url, &response)
		registry.Logf("next.url url=%s", url)
		switch err {
		case ErrNoMorePages:
			repos = append(repos, response.Repositories...)
			return repos, nil
		case nil:
			repos = append(repos, response.Repositories...)
			continue
		default:
			return nil, err
		}
	}
}
