package repos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToProperties(t *testing.T) {
	repo := Repository{
		Id:            "123",
		Name:          "test-repo",
		Description:   "A test repository",
		NameWithOwner: "test-repo-owner/test-repo",
	}

	properties := ToProperties(repo)

	if len(properties) != 49 {
		t.Fatalf("expected 49 properties, got %d", len(properties))
	}
}

func TestNewRepoConfig(t *testing.T) {
	expectedPropertyCount := 49
	expectedPropertyGroupCount := 7

	repo := Repository{
		Id:            "123",
		Name:          "test-repo",
		Description:   "A test repository",
		NameWithOwner: "test-repo-owner/test-repo",
	}

	repoConfig := NewRepoConfig(repo)

	assert.Equal(t, expectedPropertyCount, len(repoConfig.Properties))
	assert.Equal(t, expectedPropertyGroupCount, len(repoConfig.PropertyGroups))
}
