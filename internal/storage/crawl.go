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
	"strings"

	"github.com/prometheus/common/log"
	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/yaml.v2"
)

func crawl(fs billy.Filesystem, d string) error {
	ls, err := fs.ReadDir(d)
	if err != nil {
		log.Error(err)
		return err
	}

	hadError := false
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

		hadError = hadError || processFile(fs, path) != nil
	}
	if hadError {
		return errors.New("Did not parse all files, check logs")
	}
	return nil
}

func processFile(fs billy.Filesystem, path string) error {
	fc, err := fs.Open(path)
	if err != nil {
		log.Errorf("%s\n%s", path, err)
		return err
	}
	defer fc.Close()

	st, err := fs.Stat(path)
	if err != nil {
		log.Errorf("%s\n%s", path, err)
		return err
	}

	bc := make([]byte, st.Size())
	_, err = fc.Read(bc)
	if err != nil && err != io.EOF {
		log.Errorf("%s\n%s", path, err)
		return err
	}

	article := Article{}
	err = yaml.Unmarshal(bc, &tip)
	if err != nil {
		log.Errorf("%s\n%s", path, err)
		return err
	}
	return nil
}
