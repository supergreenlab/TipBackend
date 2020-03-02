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
	"regexp"
)

var repoRegExp = regexp.MustCompile("github.com/([^/]+)/([^.|$]+)")

// Tip -
type Tip struct {
	User    string  `json:"user"`
	Repo    string  `json:"repo"`
	Branch  string  `json:"branch"`
	Phase   string  `json:"phase"`
	Stage   string  `json:"stage"`
	Name    string  `json:"name"`
	Lang    string  `json:"lang"`
	Article Article `json:"article"`
}

// NewTipForGithubRepo -
func NewTipForGithubRepo(repo string) (Tip, error) {
	ms := repoRegExp.FindAllStringSubmatch(repo, -1)
	if len(ms) >= 1 && len(ms[0]) == 3 {
		return Tip{User: ms[0][1], Repo: ms[0][2], Branch: "master"}, nil
	}
	return Tip{}, errors.New("malformed github repo url")
}

func (t Tip) copyWith(phase, stage, name, lang string, article Article) Tip {
	return Tip{
		User:    t.User,
		Repo:    t.Repo,
		Branch:  t.Branch,
		Phase:   phase,
		Stage:   stage,
		Name:    name,
		Lang:    lang,
		Article: article,
	}
}

// Article -
type Article struct {
	Lang      string     `json:"lang"`
	Name      string     `json:"name"`
	Tags      []string   `yaml:"tags" json:"tags"`
	Reftime   string     `yaml:"reftime" json:"refTime"`
	Intro     Section    `yaml:"intro" json:"intro"`
	Sections  []Section  `yaml:"sections" json:"sections"`
	Products  []Product  `yaml:"products" json:"products"`
	Sources   []Source   `yaml:"sources" json:"sources"`
	Reminders []Reminder `yaml:"reminders" json:"reminders"`
	Triggers  []Trigger  `yaml:"triggers" json:"triggers"`
}

// Section -
type Section struct {
	Title       string    `yaml:"title"`
	Image       Image     `yaml:"image"`
	Text        string    `yaml:"text"`
	SubSections []Section `yaml:"subsections"`
	Link        Link      `yaml:"link"`
}

// Image -
type Image struct {
	URL    string `yaml:"url"`
	Layout string `yaml:"layout"`
}

// Link -
type Link struct {
	Title string `yaml:"title"`
	To    string `yaml:"to"`
}

// Source -
type Source struct {
	Title string `yaml:"title"`
	URL   string `yaml:"url"`
}

// Product -
type Product struct {
	Name string `yaml:"name"`
	Urls []struct {
		Zone  string `yaml:"zone"`
		Image string `yaml:"image"`
		URL   string `yaml:"url"`
	}
}

// Reminder -
type Reminder struct {
	Delay string `yaml:"delay"`
	Text  string `yaml:"text"`
	Next  string `yaml:"next"`
}

// Trigger -
type Trigger struct {
	ID     string `yaml:"id"`
	Metric string `yaml:"metric"`
	Value  string `yaml:"value"`
}
