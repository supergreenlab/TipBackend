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

package ghook

import (
	"encoding/json"
	"net/http"

	"github.com/SuperGreenLab/TipServer/internal/storage"
	"github.com/google/go-github/v29/github"
	"github.com/julienschmidt/httprouter"
)

// ServeGithubHookHandler Github hook
func ServeGithubHookHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := github.ValidatePayload(r, []byte(*hookSecret))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	gp := gitPush{}
	if r.Body == nil {
		http.Error(w, "Missing body", 400)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&gp)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	storage.UpdateTreeFromRepo()

	w.Write([]byte("OK"))
}
