package lingua

// constants that are not pertaining to build tags

var empty struct{}

// NumberWords was generated with this python code
/*
	numberWords = {}

	simple = '''zero one two three four five six seven eight nine ten eleven twelve
	        thirteen fourteen fifteen sixteen seventeen eighteen nineteen
	        twenty'''.split()
	for i, word in zip(xrange(0, 20+1), simple):
	    numberWords[word] = i

	tense = '''thirty forty fifty sixty seventy eighty ninety hundred'''.split()
	for i, word in zip(xrange(30, 100+1, 10), tense):
		numberWords[word] = i

	larges = '''thousand million billion trillion quadrillion quintillion sextillion septillion'''.split()
	for i, word in zip(xrange(3, 24+1, 3), larges):
		numberWords[word] = 10**i
*/
var NumberWords = map[string]int{
	"nol":         0,
	"satu":        1,
	"dua":         2,
	"tiga":        3,
	"empat":       4,
	"lima":        5,
	"enam":        6,
	"tujuh":       7,
	"delapan":     8,
	"sembilan":    9,
	"sepuluh":     10,
	"sebelas":     11,
	"belas":       12,
	"puluh":       10,
	"seratus":     100,
	"ratus":       100,
	"ribu":        1000,
	"juta":        1000000,
	"milyar":      1000000000,
	"triliun":     1000000000000,
	"quadrillion": 1000000000000000,
	// "quintillion": 1000000000000000000,
	// "sextillion": 1000000000000000000000,
	// "septillion": 1000000000000000000000000,
}
