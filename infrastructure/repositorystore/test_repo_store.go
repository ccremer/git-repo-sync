package repositorystore

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ccremer/greposync/domain"
	giturls "github.com/whilp/git-urls"
)

// TestRepositoryStore implements domain.GitRepositoryStore but doesn't actually use or interact with Git repositories.
// Instead, the purpose is to "fake" Git repositories so that they can be used for testing templates.
type TestRepositoryStore struct {
	StoreConfig
	// TestOutputRootDir is the root dir where the test artifacts are rendered per repository.
	// This is used for comparing the diff.
	TestOutputRootDir string

	instrumentation *RepositoryStoreInstrumentation
}

// NewTestRepositoryStore creates a new TestRepositoryStore.
func NewTestRepositoryStore(instrumentation *RepositoryStoreInstrumentation) *TestRepositoryStore {
	return &TestRepositoryStore{
		instrumentation: instrumentation,
	}
}

// ErrNotSupported indicates that a method is not supported.
var ErrNotSupported = fmt.Errorf("not supported")

// FetchGitRepositories implements domain.GitRepositoryStore.
func (s *TestRepositoryStore) FetchGitRepositories() ([]*domain.GitRepository, error) {
	list := make([]*domain.GitRepository, 0)
	dirs, err := os.ReadDir(filepath.Clean(s.ParentDir))
	if err != nil {
		return list, err
	}

	for _, dirEntry := range dirs {
		if !dirEntry.IsDir() || strings.HasPrefix(dirEntry.Name(), ".") {
			continue
		}
		u, err := giturls.Parse(dirEntry.Name())
		if err != nil {
			return list, err
		}
		repo := domain.NewGitRepository(domain.FromURL(u), domain.NewFilePath(s.ParentDir, dirEntry.Name()))
		list = append(list, repo)
	}
	return list, nil
}

// Diff implements domain.GitRepositoryStore.
// Since this implementation is meant for testing local fake repositories, the diff will be computed against the files stored in TestRepositoryStore.TestOutputRootDir.
// The files in the repository's RootDir are the expected files, where "---" is the expected file content, and "+++" the actual content.
func (s *TestRepositoryStore) Diff(repository *domain.GitRepository, _ domain.DiffOptions) (string, error) {
	args := []string{"diff", "--no-index", repository.RootDir.String(), filepath.Join(s.TestOutputRootDir, repository.URL.Path)}
	cwd, _ := os.Getwd()
	fmt.Println(cwd)
	stdout, stderr, err := execGitCommand(domain.NewFilePath(cwd), s.instrumentation.logGitArguments(repository, args))
	if err != nil && stdout == "" { // if there's a diff, the exit code is still 1 (--exit-code) implied
		return "", mergeWithStdErr(err, stderr)
	}
	return stdout, nil
}

// Clone returns ErrNotSupported.
func (s *TestRepositoryStore) Clone(_ *domain.GitRepository) error {
	return ErrNotSupported
}

// Checkout returns ErrNotSupported.
func (s *TestRepositoryStore) Checkout(_ *domain.GitRepository) error {
	return ErrNotSupported
}

// Fetch returns ErrNotSupported.
func (s *TestRepositoryStore) Fetch(_ *domain.GitRepository) error {
	return ErrNotSupported
}

// Reset returns ErrNotSupported.
func (s *TestRepositoryStore) Reset(_ *domain.GitRepository) error {
	return ErrNotSupported
}

// Pull returns ErrNotSupported.
func (s *TestRepositoryStore) Pull(_ *domain.GitRepository) error {
	return ErrNotSupported
}

// Add returns ErrNotSupported.
func (s *TestRepositoryStore) Add(_ *domain.GitRepository) error {
	return ErrNotSupported
}

// Commit returns ErrNotSupported.
func (s *TestRepositoryStore) Commit(_ *domain.GitRepository, _ domain.CommitOptions) error {
	return ErrNotSupported
}

// Push returns ErrNotSupported.
func (s *TestRepositoryStore) Push(_ *domain.GitRepository, _ domain.PushOptions) error {
	return ErrNotSupported
}
