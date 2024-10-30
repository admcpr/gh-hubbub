package structs

import (
	"gh-hubbub/queries"
	"testing"
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
	repo := queries.Repository{
		Id:            "123",
		Name:          "test-repo",
		Description:   "A test repository",
		NameWithOwner: "test-repo-owner/test-repo",
	}

	repoProperties := NewRepoProperties(repo)

	if len(repoProperties.Properties) != 50 {
		t.Fatalf("expected 4 properties, got %d", len(repoProperties.Properties))
	}

	if len(repoProperties.PropertyGroups) != 8 {
		t.Fatalf("expected 4 property groups, got %d", len(repoProperties.PropertyGroups))
	}
}
