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
	"sort"
	"strings"
)

// DefaultLang -
const DefaultLang = "en"

// CacheStorage -
type CacheStorage map[string]map[string]Tip

var (
	// Cache -
	Cache = CacheStorage{}
)

// Push -
func (cs CacheStorage) Push(tip Tip) {
	key := fmt.Sprintf("%s/%s/%s/%s/%s/%s", tip.User, tip.Repo, tip.Branch, tip.Phase, tip.Stage, tip.Article.Name)
	if _, ok := cs[key]; ok == false {
		cs[key] = map[string]Tip{}
	}
	cs[key][tip.Lang] = tip

}

// Get -
func (cs CacheStorage) Get(user, repo, branch, phase, stage, article, lang string) (Tip, bool) {
	key := fmt.Sprintf("%s/%s/%s/%s/%s/%s", user, repo, branch, phase, stage, article)
	c, ok := cs[key]
	if ok == false {
		return Tip{}, ok
	}
	t, ok := c[lang]
	return t, ok
}

// GetSlug -
func (cs CacheStorage) GetSlug(user, repo, branch, slug, lang string) (Tip, bool) {
	key := ""
	for k := range cs {
		if strings.HasPrefix(k, fmt.Sprintf("%s/%s/%s", user, repo, branch)) && strings.HasSuffix(k, slug) {
			key = k
			break
		}
	}
	if key == "" {
		return Tip{}, false
	}
	c, ok := cs[key]
	if ok == false {
		return Tip{}, false
	}
	t, ok := c[lang]
	return t, true
}

// List -
func (cs CacheStorage) List(from, to int, user, repo, branch, phase, stage, article, lang string) []Tip {
	keyPrefixComponent := []string{phase, stage, article}
	keyPrefix := fmt.Sprintf("%s/%s/%s", user, repo, branch)
	for _, v := range keyPrefixComponent {
		if v != "" {
			keyPrefix = fmt.Sprintf("%s/%s", keyPrefix, v)
		}
	}
	keys := make([]string, 0, len(cs))
	for k := range cs {
		if strings.HasPrefix(k, keyPrefix) {
			keys = append(keys, k)
		}
	}

	if from >= to {
		from = -1
		to = -1
	}

	if from < 0 {
		from = 0
	}
	if to < 0 || to > len(keys) {
		to = len(keys)
	}

	sort.Strings(keys)
	keys = keys[from:to]
	res := make([]Tip, len(keys))
	for i, v := range keys {
		t, ok := cs[v][lang]
		if ok == false {
			t = cs[v][DefaultLang]
		}
		res[i] = t
	}
	return res
}
