package storage

// Tip -
type Tip struct {
	Tags      []string   `yaml:"tags"`
	Reftime   string     `yaml:"reftime"`
	Intro     Section    `yaml:"intro"`
	Products  []Product  `yaml:"products"`
	Reminders []Reminder `yaml:"reminders"`
	Triggers  []Trigger  `yaml:"triggers"`
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
