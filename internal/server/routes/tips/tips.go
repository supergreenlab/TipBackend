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

package tips

import (
	"net/http"

	"github.com/SuperGreenLab/TipBackend/internal/storage"
	"github.com/juju/httprequest"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// ServeTips -
func ServeTips(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := storage.Cache.Get(
		p.ByName("user"),
		p.ByName("branch"),
		p.ByName("phase"),
		p.ByName("stage"),
		p.ByName("article"),
	)
	if err := httprequest.WriteJSON(w, http.StatusOK, c); err != nil {
		log.Error(err)
	}
}
