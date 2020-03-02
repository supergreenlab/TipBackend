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
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/yaml.v2"
)

var (
	pathRegexp = regexp.MustCompile("/([^/]+)/([^/]+)/([^/]+)/([^.]+).yml")
)

func crawl(base Tip, fs billy.Filesystem, d string) error {
	ls, err := fs.ReadDir(d)
	if err != nil {
		log.Error(err)
		return err
	}

	hadError := false
	for _, f := range ls {
		path := fmt.Sprintf("%s/%s", d, f.Name())
		if f.IsDir() {
			crawl(base, fs, path)
			continue
		}

		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		if !strings.HasSuffix(f.Name(), ".yml") {
			continue
		}

		pm := pathRegexp.FindAllStringSubmatch(path, -1)
		if len(pm) >= 1 && len(pm[0]) == 5 {
			article, err := processFile(fs, path)
			hadError = hadError || err != nil
			article.Name = pm[0][3]
			article.Lang = pm[0][4]
			Cache.Push(base.copyWith(pm[0][1], pm[0][2], pm[0][3], pm[0][4], article))
		}
	}
	if hadError {
		return errors.New("Did not parse all files, check logs")
	}
	return nil
}

func processFile(fs billy.Filesystem, path string) (Article, error) {
	article := Article{}
	fc, err := fs.Open(path)
	if err != nil {
		log.Errorf("%s\n%s", path, err)
		return article, err
	}
	defer fc.Close()

	st, err := fs.Stat(path)
	if err != nil {
		log.Errorf("%s\n%s", path, err)
		return article, err
	}

	bc := make([]byte, st.Size())
	_, err = fc.Read(bc)
	if err != nil && err != io.EOF {
		log.Errorf("%s\n%s", path, err)
		return article, err
	}

	err = yaml.Unmarshal(bc, &article)
	if err != nil {
		log.Errorf("%s\n%s", path, err)
		return article, err
	}
	return article, nil
}
