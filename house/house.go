package house

var parts = []string{
	"the horse and the hound and the horn\nthat belonged to",
	"the farmer sowing his corn\nthat kept",
	"the rooster that crowed in the morn\nthat woke",
	"the priest all shaven and shorn\nthat married",
	"the man all tattered and torn\nthat kissed",
	"the maiden all forlorn\nthat milked",
	"the cow with the crumpled horn\nthat tossed",
	"the dog\nthat worried",
	"the cat\nthat killed",
	"the rat\nthat ate",
	"the malt\nthat lay in",
}

func Song() string {
	song := "This is the house that Jack built."
	for verse := 1; verse < len(parts) + 1; verse++ {
		song += "\n\n"
		song += Verse("This is", parts[len(parts) - verse:], "the house that Jack built.", )
	}
	return song
}

func Verse(prefix string, phrases []string, postfix string) string {
	if len(phrases) == 0 {
		return Embed(prefix, postfix)
	}
	return Verse(Embed(prefix, phrases[0]), phrases[1:], postfix)
}

func Embed(l, p string) string {
	return l + " " + p
}
