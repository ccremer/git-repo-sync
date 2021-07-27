package github

import (
	"github.com/ccremer/greposync/cfg"
	"github.com/ccremer/greposync/core"
)

type (
	// LabelConverter converts core.GitRepositoryLabel to cfg.RepositoryLabel and vice-versa
	LabelConverter struct{}
)

func (LabelConverter) convertToEntity(labels []*cfg.RepositoryLabel) []core.GitRepositoryLabel {
	if labels == nil || len(labels) == 0 {
		return []core.GitRepositoryLabel{}
	}
	converted := make([]core.GitRepositoryLabel, len(labels))
	for i := range labels {
		converted[i] = labels[i]
	}
	return converted
}

func (LabelConverter) convertFromEntity(labels []core.GitRepositoryLabel) []*cfg.RepositoryLabel {
	if labels == nil || len(labels) == 0 {
		return []*cfg.RepositoryLabel{}
	}
	converted := make([]*cfg.RepositoryLabel, len(labels))
	for i := range labels {
		converted[i] = labels[i].(*cfg.RepositoryLabel)
	}
	return converted
}
