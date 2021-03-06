package twelve

const testVersion = 1

var stuff = []struct {
	day   string
	thing string
}{
	{"first", "a Partridge in a Pear Tree" },
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

func Song() string {
	song := ""
	for i := 1; i <= len(stuff); i++ {
		song += Verse(i) + "\n"
	}
	return song
}

func Verse(verse int) string {

	things := ""
	for i := verse; i > 0; i-- {
		things += stuff[i - 1].thing
		if i == 2 {
			things += ", and "
		} else if i > 2 {
			things += ", "
		}

	}

	return "On the " + stuff[verse - 1].day + " day of Christmas my true love gave to me, " + things + "."
}

