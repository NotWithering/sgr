package sgr

const (
	Reset     string = "\033[0m"
	Bold      string = "\033[1m"
	Faint     string = "\033[2m"
	Italic    string = "\033[3m"
	Underline string = "\033[4m"
	SlowBlink string = "\033[5m"
	FastBlink string = "\033[6m"
	Invert    string = "\033[7m"
	Conceal   string = "\033[8m"
	Strike    string = "\033[9m"

	DefaultFont      string = "\033[10m"
	AlternativeFont1 string = "\033[12m"
	AlternativeFont2 string = "\033[13m"
	AlternativeFont3 string = "\033[14m"
	AlternativeFont4 string = "\033[15m"
	AlternativeFont5 string = "\033[16m"
	AlternativeFont6 string = "\033[17m"
	AlternativeFont7 string = "\033[18m"
	AlternativeFont8 string = "\033[19m"

	Fraktur          string = "\033[20m"
	DoublyUnderlined string = "\033[21m"
	NotBold          string = "\033[21m"
	NormalIntensity  string = "\033[22m"
	NotItalic        string = "\033[23m"
	NotBlackletter   string = "\033[23m"
	NotUnderlined    string = "\033[24m"
	NotBlinking      string = "\033[25m"
	Monospace        string = "\033[26m"
	NotInverted      string = "\033[27m"
	Reveal           string = "\033[28m"
	NotStriked       string = "\033[29m"

	FgBlack   string = "\033[30m"
	FgRed     string = "\033[31m"
	FgGreen   string = "\033[32m"
	FgYellow  string = "\033[33m"
	FgBlue    string = "\033[34m"
	FgMagenta string = "\033[35m"
	FgCyan    string = "\033[36m"
	FgWhite   string = "\033[37m"
	// FgColor string = "\033[38m"
	FgDefault string = "\033[39m"

	BgBlack   string = "\033[40m"
	BgRed     string = "\033[41m"
	BgGreen   string = "\033[42m"
	BgYellow  string = "\033[43m"
	BgBlue    string = "\033[44m"
	BgMagenta string = "\033[45m"
	BgCyan    string = "\033[46m"
	BgWhite   string = "\033[47m"
	// BgColor string = "\033[48m"
	BgDefault string = "\033[49m"

	NotMonospace string = "\033[50m"
	Frame        string = "\033[51m"
	Encircle     string = "\033[52m"
	Overline     string = "\033[53m"
	NotFramed    string = "\033[54m"
	NotEncircled string = "\033[54m"
	NotOverlined string = "\033[55m"
	// UnderlineColor string = "\033[58m"
	DefaultUnderlineColor string = "\033[59m"

	IdeogreamUnderline      string = "\033[60m"
	IdeogramDoubleUnderline string = "\033[61m"
	IdeogramOverline        string = "\033[62m"
	IdeogramDoubleOverline  string = "\033[63m"
	IdeogramStress          string = "\033[64m"
	NoIdeogram              string = "\033[65m"

	Superscript    string = "\033[73m"
	Subscript      string = "\033[74m"
	NotSuperscript string = "\033[75m"
	NotSubscript   string = "\033[75m"

	FgHiBlack   string = "\033[90m"
	FgHiRed     string = "\033[91m"
	FgHiGreen   string = "\033[92m"
	FgHiYellow  string = "\033[93m"
	FgHiBlue    string = "\033[94m"
	FgHiMagenta string = "\033[95m"
	FgHiCyan    string = "\033[96m"
	FgHiWhite   string = "\033[97m"

	BgHiBlack   string = "\033[100"
	BgHiRed     string = "\033[101m"
	BgHiGreen   string = "\033[102m"
	BgHiYellow  string = "\033[103m"
	BgHiBlue    string = "\033[104m"
	BgHiMagenta string = "\033[105m"
	BgHiCyan    string = "\033[106m"
	BgHiWhite   string = "\033[107m"
)
