package core

import (
	"errors"
	"io/fs"
)

// Values is a key-value construct with arbitrary hierarchy.
type Values map[string]interface{}

// Template is a representation of a single template file.
//counterfeiter:generate . Template
type Template interface {
	// GetRelativePath returns the path to a template file relative to the template root directory.
	// The path is delimited with a forward slash ("/") and not OS-specific.
	GetRelativePath() string
	// GetFileMode returns the mode of the template file.
	GetFileMode() fs.FileMode
	// Render takes the given Values and returns a string that essentially is the content of a file.
	Render(values Values) (string, error)
}

// TemplateStore is a service responsible for fetching templates.
//counterfeiter:generate . TemplateStore
type TemplateStore interface {
	// FetchTemplates retrieves the templates or an error if one failed.
	FetchTemplates() ([]Template, error)

	// FetchPullRequestTemplate retrieves the template for pull requests.
	// It returns nil if there is no template found or configured.
	FetchPullRequestTemplate() (Template, error)
}

// ValueStore is a service centered around configuration values fetching and configuring templates.
type ValueStore interface {
	// FetchValuesForTemplate retrieves the Values for the given template.
	FetchValuesForTemplate(template Template, config *GitRepositoryProperties) (Values, error)
	// FetchUnmanagedFlag returns true if the given template should not be rendered.
	// The implementation may return ErrKeyNotFound if the flag is undefined, as the boolean 'false' is ambiguous.
	FetchUnmanagedFlag(template Template, config *GitRepositoryProperties) (bool, error)
	// FetchTargetPath returns an alternative output path for the given template relative to the Git repository.
	// An empty string indicates that there is no alternative path configured.
	FetchTargetPath(template Template, config *GitRepositoryProperties) (string, error)
	// FetchFilesToDelete returns a slice of paths that should be deleted in the Git repository.
	// The paths are relative to the Git directory.
	FetchFilesToDelete(config *GitRepositoryProperties) ([]string, error)
}

// ErrKeyNotFound is an error that indicates that a particular key was not found in the ValueStore.
var ErrKeyNotFound = errors.New("key not found")
