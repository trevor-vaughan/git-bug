package repository

import (
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/filesystem"
)

// relativeWorktreesStorage hides extensions.relativeWorktrees from go-git v5,
// which refuses every extension but a few and so can't open a repository where
// `git worktree add --relative-paths` or worktree.useRelativePaths has been
// used. The extension only changes how git writes the links between a linked
// worktree and the repository, which git-bug neither reads nor writes.
type relativeWorktreesStorage struct {
	*filesystem.Storage
}

func (s relativeWorktreesStorage) Config() (*config.Config, error) {
	cfg, err := s.Storage.Config()
	if err != nil {
		return nil, err
	}
	if cfg.Raw != nil && cfg.Raw.HasSection("extensions") {
		cfg.Raw.Section("extensions").RemoveOption("relativeWorktrees")
	}
	return cfg, nil
}
