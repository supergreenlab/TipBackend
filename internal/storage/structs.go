package storage

// Tip -
type Tip struct {
	User    string             `json:"user"`
	Branch  string             `json:"branch"`
	Phase   string             `json:"phase"`
	Stage   string             `json:"stage"`
	Article map[string]Article `json:"article"`
}

func (t Tip) copyWith(phase, stage string) {
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
