package kibela

import (
	"fmt"
	"sync"

	"github.com/Songmu/kibela/client"
	"golang.org/x/xerrors"
)

type kibela struct {
	cli *client.Client

	groups     map[string]ID
	groupsErr  error
	groupsOnce sync.Once
}

func newKibela() (*kibela, error) {
	cli, err := client.New(version)
	if err != nil {
		return nil, xerrors.Errorf("failed to newKibela: %w", err)
	}
	return &kibela{cli: cli}, nil
}

func (ki *kibela) fetchGroups() (map[string]ID, error) {
	ki.groupsOnce.Do(func() {
		if ki.groups != nil {
			return
		}
		groups, err := ki.getGroups()
		if inErr != nil {
			ki.groupsErr = xerrors.Errorf("failed to ki.setGroups: %w", err)
			return
		}
		groupMap := make(map[string]ID, len(groups))
		for _, g := range groups {
			groupMap[g.Name] = g.ID
		}
		ki.groups = groupMap
	})
	return ki.groups, ki.groupsErr
}

func (ki *kibela) fetchGroupID(name string) (ID, error) {
	groups, err := ki.fetchGroups()
	if err != nil {
		return "", xerrors.Errorf("failed to fetchGroupID while setGroupID: %w", err)
	}
	id, ok := groups[name]
	if !ok {
		return "", fmt.Errorf("group %q doesn't exists", name)
	}
	return id, nil
}
