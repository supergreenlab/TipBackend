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

package assets

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SuperGreenLab/TipBackend/internal/storage"
	"github.com/juju/httprequest"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// ServeAsset -
func ServeAsset(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path, ok := storage.Cache.GetPathFromSlug(p.ByName("user"), p.ByName("repo"), p.ByName("branch"), p.ByName("slug"))
	if ok == false {
		httprequest.WriteJSON(w, http.StatusNotFound, map[string]string{"status": "Not found"})
		return
	}

	path = strings.ReplaceAll(path, fmt.Sprintf("%s/%s/%s", p.ByName("user"), p.ByName("repo"), p.ByName("branch")), "")
	b, err := storage.GetFileAt(fmt.Sprintf("%s/assets/%s", path, p.ByName("file")))

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		log.Error(err)
	}
}
