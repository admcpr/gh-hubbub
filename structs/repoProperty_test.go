package structs

import (
	"gh-hubbub/queries"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToProperties(t *testing.T) {
	repo := queries.Repository{
		Id:            "123",
		Name:          "test-repo",
		Description:   "A test repository",
		NameWithOwner: "test-repo-owner/test-repo",
	}

	properties := ToProperties(repo)

	if len(properties) != 50 {
		t.Fatalf("expected 50 properties, got %d", len(properties))
	}
}

func TestNewRepoProperties(t *testing.T) {
	expectedPropertyCount := 49
	expectedPropertyGroupCount := 5

	repo := queries.Repository{
		Id:            "123",
		Name:          "test-repo",
		Description:   "A test repository",
		NameWithOwner: "test-repo-owner/test-repo",
	}

	repoProperties := NewRepoProperties(repo)

	assert.Equal(t, expectedPropertyCount, len(repoProperties.Properties))
	assert.Equal(t, expectedPropertyGroupCount, len(repoProperties.PropertyGroups))
}
