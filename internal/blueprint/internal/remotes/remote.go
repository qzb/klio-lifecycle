package remotes

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/g2a-com/klio-logger-go"
)

func Fetch(projectDir string, url string, rev string) error {
	var err error
	repoDir, alreadyFetched := GetDir(projectDir, url, rev)

	// Cleanup repo dir (only when function exits prematurely)
	skipDefer := false
	defer func() {
		if !skipDefer {
			err := os.RemoveAll(repoDir)
			if err != nil {
				log.Errorf("failed to cleanup repository from cache, please remove it manually: rm -r %s", repoDir)
			}
		}
	}()

	if !alreadyFetched {
		// Prepare directory for repository
		err = os.MkdirAll(repoDir, 0755)
		if err != nil {
			return err
		}

		// Create empty repository
		log.Debugf("initializing git repository: %s", repoDir)
		err = runGitCmd(repoDir, "init")
		if err != nil {
			return err
		}

		// Add remote
		log.Spamf("adding remote: %s", url)
		err = runGitCmd(repoDir, "remote", "add", "origin", url)
		if err != nil {
			return err
		}
	}

	// Fetch from origin if not fully initialized or if revision is a branch
	if runGitCmd(repoDir, "symbolic-ref", "HEAD") == nil {
		// Fetch specified version
		log.Verbosef("fetching commit %s from %s", url, rev)
		err = runGitCmd(repoDir, "fetch", "--depth", "1", "origin", rev)
		if err != nil {
			return fmt.Errorf("failed to fetch git revision, take a note that only tags and FULL commit hashes are supported: %s", err)
		}

		// Checkout desired version
		log.Debugf("checking out newly fetched version")
		err = runGitCmd(repoDir, "reset", "--hard", "origin/"+rev)
		if err != nil {
			return err
		}
	}

	skipDefer = true

	return nil
}

func GetDir(projectDir string, uri string, rev string) (string, bool) {
	dir := filepath.Join(
		projectDir,
		".g2a",
		"cache",
		"cicd",
		"git-repositories",
		url.PathEscape(uri),
		url.PathEscape(rev),
	)

	_, err := os.Stat(dir)
	ok := err == nil

	return dir, ok
}

func runGitCmd(dir string, args ...string) error {
	l := log.StandardLogger()
	stderr := bytes.Buffer{}

	c := exec.Command("git", args...)
	c.Dir = dir
	c.Stdout = l.WithLevel(log.SpamLevel)
	c.Stderr = io.MultiWriter(l.WithLevel(log.VerboseLevel), &stderr)

	err := c.Run()
	switch err.(type) {
	case *exec.ExitError:
		return errors.New(stderr.String())
	default:
		return err
	}
}
