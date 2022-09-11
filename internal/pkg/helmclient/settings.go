package helmclient

import (
	"context"
	"github.com/Masterminds/semver"
	"github.com/gofrs/flock"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/cmd/helm/search"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/helmpath"
	"helm.sh/helm/v3/pkg/repo"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	settingsOnce sync.Once
	settings     *helmSettings
)

const (
	SearchMaxScore = 25
)

type helmSettings struct {
	settings *cli.EnvSettings
}

func GetHelmSettings() *helmSettings {
	settingsOnce.Do(func() {
		settings = &helmSettings{settings: cli.New()}
	})
	return settings
}

func (settings *helmSettings) GetSettings() *cli.EnvSettings {
	return settings.settings
}

func (settings *helmSettings) InitRepos(repoEntryList []*repo.Entry) {
	for _, repoEntry := range repoEntryList {
		if err := settings.initRepo(repoEntry); err != nil {
			log.Println(err)
		}
	}
}

func (settings *helmSettings) initRepo(repoEntry *repo.Entry) error {
	// Ensure the file directory exists as it is required for file locking
	err := os.MkdirAll(filepath.Dir(settings.settings.RepositoryConfig), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	// Acquire a file lock for process synchronization
	fileLock := flock.New(strings.Replace(settings.settings.RepositoryConfig, filepath.Ext(settings.settings.RepositoryConfig), ".lock", 1))
	lockCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	locked, err := fileLock.TryLockContext(lockCtx, time.Second)
	if err == nil && locked {
		settings.safeCloser(fileLock, &err)
	}
	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(settings.settings.RepositoryConfig)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	var f repo.File
	if err := yaml.Unmarshal(b, &f); err != nil {
		return err
	}

	r, err := repo.NewChartRepository(repoEntry, getter.All(settings.settings))
	if err != nil {
		return err
	}

	if _, err := r.DownloadIndexFile(); err != nil {
		return err
	}

	f.Update(repoEntry)

	if err := f.WriteFile(settings.settings.RepositoryConfig, 0644); err != nil {
		return err
	}

	return nil
}

func (settings *helmSettings) safeCloser(fileLock *flock.Flock, err *error) {
	if fileErr := fileLock.Unlock(); fileErr != nil && *err == nil {
		*err = fileErr
		glog.Error(fileErr)
	}
}

func (settings *helmSettings) ApplyConstraint(version string, versions bool, res []*search.Result) ([]*search.Result, error) {
	if len(version) == 0 {
		return res, nil
	}

	constraint, err := semver.NewConstraint(version)
	if err != nil {
		return res, errors.Wrap(err, "an invalid version/constraint format")
	}

	data := res[:0]
	foundNames := map[string]bool{}
	for _, r := range res {
		if _, found := foundNames[r.Name]; found {
			continue
		}
		v, err := semver.NewVersion(r.Chart.Version)
		if err != nil || constraint.Check(v) {
			data = append(data, r)
			if !versions {
				foundNames[r.Name] = true // If user hasn't requested all versions, only show the latest that matches
			}
		}
	}

	return data, nil
}

func (settings *helmSettings) BuildSearchIndex(repoName, version string) (*search.Index, error) {
	i := search.NewIndex()
	f := filepath.Join(settings.settings.RepositoryCache, helmpath.CacheIndexFile(repoName))
	ind, err := repo.LoadIndexFile(f)
	if err != nil {
		log.Printf("WARNING: Repo %q is corrupt or missing. Try 'helm repo update'.", repoName)
		return nil, err
	}
	i.AddRepo(repoName, ind, len(version) > 0)
	return i, nil
}

func (settings *helmSettings) UpdateRepo(entry *repo.Entry) error {
	r, err := repo.NewChartRepository(entry, getter.All(settings.settings))
	if err != nil {
		return err
	}
	_, err = r.DownloadIndexFile()
	if err != nil {
		return err
	}

	return nil
}
