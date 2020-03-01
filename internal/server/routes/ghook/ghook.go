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

	"github.com/SuperGreenLab/TipBackend/internal/storage"
	"github.com/google/go-github/v29/github"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ServeGithubHookHandler Github hook
func ServeGithubHookHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	hs := viper.GetString("GithubHookSecret")
	payload, err := github.ValidatePayload(r, []byte(hs))
	if err != nil {
		log.Errorf("ValidatePayload error:\n%s", err)
		http.Error(w, err.Error(), 400)
		return
	}
	gp := gitPush{}
	err = json.Unmarshal(payload, &gp)
	if err != nil {
		logrus.Errorf("JSON decode error:\n%s", err)
		http.Error(w, err.Error(), 400)
		return
	}

	storage.UpdateTreeFromRepo()

	w.Write([]byte("OK"))
}
