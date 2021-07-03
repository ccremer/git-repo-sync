package repository

import (
	"errors"
	"strings"

	pipeline "github.com/ccremer/go-command-pipeline"
)

func (s *Service) SkipCheckout() pipeline.Predicate {
	return func(step pipeline.Step) bool {
		return s.Config.SkipReset || s.Config.CommitBranch == ""
	}
}

func (s *Service) CheckoutBranch() pipeline.ActionFunc {
	return func() pipeline.Result {
		out, stderr, err := s.execGitCommand(s.logArgs("checkout", s.Config.CommitBranch)...)
		if err != nil {
			return s.toResult(err, stderr)
		}
		s.p.DebugF(out)
		return pipeline.Result{}
	}
}

func (s *Service) GetDefaultBranch() pipeline.ActionFunc {
	return func() pipeline.Result {
		out, stderr, err := s.execGitCommand("remote", "show", "origin")
		if err != nil {
			return s.toResult(err, stderr)
		}
		lines := strings.Split(out, "\n")
		head := "HEAD branch: "
		for _, line := range lines {
			str := strings.TrimSpace(line)
			if strings.Contains(str, head) {
				s.Config.DefaultBranch = strings.TrimPrefix(str, head)
				return pipeline.Result{}
			}
		}
		s.p.WarnF("No default branch detected. Fall back to master")
		s.Config.DefaultBranch = "master"
		return pipeline.Result{}
	}
}

func (s *Service) remoteBranchExists(branch string) (bool, error) {
	out, stderr, err := s.execGitCommand("branch", "-r", "--list")
	if err != nil {
		return false, errors.New(stderr)
	}
	branches := strings.Split(out, "\n")
	for _, line := range branches {
		if strings.Contains(strings.TrimSpace(line), branch) {
			return true, nil
		}
	}
	return false, nil
}
