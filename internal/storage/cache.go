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

import "fmt"

// CacheStorage -
type CacheStorage map[string]Tip

var (
	// Cache -
	Cache = CacheStorage{}
)

// Push -
func (cs CacheStorage) Push(tip Tip) {
	key := fmt.Sprintf("%s/%s/%s/%s/%s", tip.User, tip.Branch, tip.Phase, tip.Stage, tip.Article.Name)
	cs[key] = tip
}

// Get -
func (cs CacheStorage) Get(user, branch, phase, stage, article string) Tip {
	key := fmt.Sprintf("%s/%s/%s/%s/%s", user, branch, phase, stage, article)
	return cs[key]
}
