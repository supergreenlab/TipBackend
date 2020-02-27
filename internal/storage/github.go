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
	"io"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"gopkg.in/yaml.v2"
)

var (
	fs = memfs.New()
	st = memory.NewStorage()
	r  *git.Repository
)

// InitTreeFromRepo update the storge tree from github
func InitTreeFromRepo(repo string) {
	go func() {
		var err error
		r, err = git.Clone(st, fs, &git.CloneOptions{
			URL: repo,
		})
		if err != nil {
			log.Fatal(err)
		}

		crawl(fs, "")
	}()
}

// UpdateTreeFromRepo update the storge tree from github
func UpdateTreeFromRepo() {
	go func() {
		w, err := r.Worktree()
		if err != nil {
			log.Fatal(err)
		}

		err = w.Pull(nil)
		if err != nil {
			log.Fatal(err)
		}

		crawl(fs, "")
	}()
}

func crawl(fs billy.Filesystem, d string) {
	ls, err := fs.ReadDir(d)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range ls {
		path := fmt.Sprintf("%s/%s", d, f.Name())
		if f.IsDir() {
			crawl(fs, path)
			continue
		}

		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		if !strings.HasSuffix(f.Name(), ".yml") {
			continue
		}

		log.Println(path)

		fc, err := fs.Open(path)
		if err != nil {
			log.Fatal(err)
		}

		st, err := fs.Stat(path)
		if err != nil {
			log.Fatal(err)
		}

		bc := make([]byte, st.Size())
		log.Infof("stat %d bytes", st.Size())
		_, err = fc.Read(bc)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		tip := Tip{}
		err = yaml.Unmarshal(bc, &tip)
		if err != nil {
			log.Fatal(err)
		}

		log.Info(tip)
	}
}
