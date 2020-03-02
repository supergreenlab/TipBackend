/*
 * Copyright (C) 2020  SuperGreenLab <towelie@supergreenlab.com>
 * Author: Constantin Clauzel <constantin.clauzel@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package storage

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
)

var (
	repos = map[string]*git.Repository{}
)

// InitTreeFromRepo init the storage tree from github
func InitTreeFromRepo(repo string) {
	go func() {
		tip, err := NewTipForGithubRepo(repo)
		if err != nil {
			log.Error(err)
			return
		}

		key := fmt.Sprintf("%s/%s", tip.User, tip.Repo)
		r, err := git.Clone(st, fs, &git.CloneOptions{
			URL: repo,
		})
		if err != nil {
			log.Error(err)
			return
		}
		repos[key] = r

		err = startCrawl(tip, fs, "")
		if err != nil {
			log.Warning("Did not update tips tree")
		} else {
			log.Info("Updated tips tree")
		}
	}()
}

// UpdateTreeFromRepo update the storage tree from github
func UpdateTreeFromRepo(repo string) {
	go func() {
		tip, err := NewTipForGithubRepo(repo)
		if err != nil {
			log.Error(err)
			return
		}

		key := fmt.Sprintf("%s/%s", tip.User, tip.Repo)
		r, ok := repos[key]
		if ok != true {
			log.Warningf("Unknown repo %s, initializing...", key)
			InitTreeFromRepo(repo)
			return
		}

		w, err := r.Worktree()
		if err != nil {
			log.Error(err)
			return
		}

		err = w.Pull(&git.PullOptions{RemoteName: "origin"})
		if err != nil {
			log.Error(err)
			return
		}

		err = startCrawl(tip, fs, "")
		if err != nil {
			log.Warning("Did not update tips tree")
		} else {
			log.Info("Updated tips tree")
		}
	}()
}
