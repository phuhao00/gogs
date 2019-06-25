// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/gogs/gogs/models/errors"
	"strconv"
	"strings"
)

var (
	ErrMissingIssueCollectionNumber = errors.New("No IssueCollection number specified")
)

// Issue represents an issue or pull request of repository.
type IssueCollection struct {
	UserId              int64
	IssueIds			string
	IssueId2Issue		map[int64]*Issue
	IssueId2Repo		map[int64]*Repository
}

func (issue *IssueCollection) BeforeInsert() {

}

func (issue *IssueCollection) BeforeUpdate() {

}

func (issue *IssueCollection) AfterSet(colName string, _ xorm.Cell) {
	switch colName {

	}
}
//
func (issueCollection *IssueCollection) loadAttributes(e Engine) (err error) {
	if issueCollection.UserId != 0 {
		issue2repos:=strings.Split(issueCollection.IssueIds,",")
		for _, value := range issue2repos {
			issueIds:=strings.Split(value,":")
			repoId,err:=strconv.ParseInt(issueIds[1], 0, 64)
			issueId,err:=strconv.ParseInt(issueIds[0], 0, 64)
			if err==nil {
				issueCollection.IssueId2Repo[issueId], err = getRepositoryByID(e, repoId)
			}else {
				return fmt.Errorf("getRepositoryByID [%d]: %v", issueId, err)
			}
		}
	}
	return nil
}


