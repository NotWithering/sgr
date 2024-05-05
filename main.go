// Package SGR contains a large block of constants that attempts to implement every SGR Ansi escape code.
package sgr

const (
	Reset     string = "\x1b[0m" // All attributes become turned off
	Bold      string = "\x1b[1m" // As with faint, the color change is a PC (SCO / CGA) invention.
	Faint     string = "\x1b[2m" //	May be implemented as a light font weight like bold.
	Italic    string = "\x1b[3m" // Not widely supported. Sometimes treated as inverse or blink.
	Underline string = "\x1b[4m" // Style extensions exists for Kitty, VTE, mintty, iTerm2 and Konsole.
	Blink     string = "\x1b[5m" // Sets blinking to less than 150 times per minute
	SlowBlink string = "\x1b[5m" // Sets blinking to less than 150 times per minute
	FastBlink string = "\x1b[6m" // MS-DOS ANSI.SYS 150+ per minute; not widely supported
	Invert    string = "\x1b[7m" // Swap foreground and background colors; inconsistent emulation
	Conceal   string = "\x1b[8m" // Not widely supported.
	Strike    string = "\x1b[9m" // Characters legible but marked as if for deletion. Not supported in Terminal.app.

	DefaultFont      string = "\x1b[10m" // Primary font
	AlternativeFont1 string = "\x1b[12m" // 1st alternative font
	AlternativeFont2 string = "\x1b[13m" // 2nd alternative font
	AlternativeFont3 string = "\x1b[14m" // 3rd alternative font
	AlternativeFont4 string = "\x1b[15m" // 4th alternative font
	AlternativeFont5 string = "\x1b[16m" // 5th alternative font
	AlternativeFont6 string = "\x1b[17m" // 6th alternative font
	AlternativeFont7 string = "\x1b[18m" // 7th alternative font
	AlternativeFont8 string = "\x1b[19m" // 8th alternative font

	Fraktur          string = "\x1b[20m" // Rarely supported
	DoublyUnderlined string = "\x1b[21m" // Double-underline per ECMA-48 but instead disables bold intensity on several terminals, including in the Linux kernel's console before version 4.17.
	NotBold          string = "\x1b[21m" // Double-underline per ECMA-48 but instead disables bold intensity on several terminals, including in the Linux kernel's console before version 4.17.
	NormalIntensity  string = "\x1b[22m" // Neither bold nor faint; color changes where intensity is implemented as such.
	NotItalic        string = "\x1b[23m" // Neither italic nor blackletter
	NotBlackletter   string = "\x1b[23m" // Neither italic nor blackletter
	NotUnderlined    string = "\x1b[24m" // Neither singly nor doubly underlined
	NotBlinking      string = "\x1b[25m" // Turn blinking off
	Monospace        string = "\x1b[26m" // ITU t.61 and T.416, not known to be used on terminals
	NotInverted      string = "\x1b[27m" // Not reversed
	Reveal           string = "\x1b[28m" // Not concealed
	NotStriked       string = "\x1b[29m" // Not crossed out

	FgBlack   string = "\x1b[30m" // VGA (0, 0, 0)
	FgRed     string = "\x1b[31m" // VGA (170, 0, 0)
	FgGreen   string = "\x1b[32m" // VGA (0, 170, 0)
	FgYellow  string = "\x1b[33m" // VGA (170, 85, 0)
	FgBlue    string = "\x1b[34m" // VGA (0, 0, 170)
	FgMagenta string = "\x1b[35m" // VGA (170, 0, 170)
	FgCyan    string = "\x1b[36m" // VGA (0, 170, 170)
	FgWhite   string = "\x1b[37m" // VGA (170, 170, 170)

	FgColor    string = "\x1b[38;5;" // Set foreground color where n is a 4-bit color index
	FgColorRGB string = "\x1b[38;2;" // Set foreground color where R;G;B are from 0 to 255

	FgDefault string = "\x1b[39m" // Implementation defined (according to standard)

	BgBlack   string = "\x1b[40m" // VGA (0, 0, 0)
	BgRed     string = "\x1b[41m" // VGA (170, 0, 0)
	BgGreen   string = "\x1b[42m" // VGA (0, 170, 0)
	BgYellow  string = "\x1b[43m" // VGA (170, 85, 0)
	BgBlue    string = "\x1b[44m" // VGA (0, 0, 170)
	BgMagenta string = "\x1b[45m" // VGA (170, 0, 170)
	BgCyan    string = "\x1b[46m" // VGA (0, 170, 170)
	BgWhite   string = "\x1b[47m" // VGA (170, 170, 170)

	BgColor    string = "\x1b[48;5;" // Set background color where n is a 4-bit color index
	BgColorRGB string = "\x1b[48;2;" // Set background color where R;G;B are from 0 to 255

	BgDefault string = "\x1b[49m" // Implementation defined (according to standard)

	NotMonospace string = "\x1b[50m" // T.61 and T.416
	Frame        string = "\x1b[51m" // Implemented as "emoji variation selector" in mintty.
	Encircle     string = "\x1b[52m" // Implemented as "emoji variation selector" in mintty.
	Overline     string = "\x1b[53m" // Not supported in Terminal.app
	NotFramed    string = "\x1b[54m" // Neither framed nor encircled
	NotEncircled string = "\x1b[54m" // Neither framed nor encircled
	NotOverlined string = "\x1b[55m" // Turn underline off

	UnderlineColor    string = "\x1b[58;5;" // Set underline color where n is a 4-bit color index
	UnderlineColorRGB string = "\x1b[58;2;" // Set underline color where R;G;B are from 0 to 255

	DefaultUnderlineColor string = "\x1b[59m" // Not in standard; implemented in Kitty, VTE, mintty, and iTerm2.

	IdeogreamUnderline      string = "\x1b[60m" // Ideogram underline or right side line; Rarely supported
	IdeogramDoubleUnderline string = "\x1b[61m" // Ideogram double underline or double line on the right side; Rarely supported
	IdeogramOverline        string = "\x1b[62m" // Ideogram overline or left side line; Rarely supported
	IdeogramDoubleOverline  string = "\x1b[63m" // Ideogram double overline or double line of the left side; Rarely supported
	IdeogramStress          string = "\x1b[64m" // Ideogram stress marking; Rarely supported
	NoIdeogram              string = "\x1b[65m" // Resets the effects of all of 60 - 64

	Superscript    string = "\x1b[73m" // Implemented only in mintty
	Subscript      string = "\x1b[74m" // Implemented only in mintty
	NotSuperscript string = "\x1b[75m" // Neither superscript nor subscript; Implemented only in mintty
	NotSubscript   string = "\x1b[75m" // Neither superscript nor subscript; Implemented only in mintty

	FgHiBlack   string = "\x1b[90m" // VGA (85, 85, 85)
	FgHiRed     string = "\x1b[91m" // VGA (255, 85, 85)
	FgHiGreen   string = "\x1b[92m" // VGA (85, 255, 85)
	FgHiYellow  string = "\x1b[93m" // VGA (255, 255, 85)
	FgHiBlue    string = "\x1b[94m" // VGA (85, 85, 255)
	FgHiMagenta string = "\x1b[95m" // VGA (255, 85, 255)
	FgHiCyan    string = "\x1b[96m" // VGA (85, 255, 255)
	FgHiWhite   string = "\x1b[97m" // VGA (255, 255, 255)

	BgHiBlack   string = "\x1b[100m" // VGA (85, 85, 85)
	BgHiRed     string = "\x1b[101m" // VGA (255, 85, 85)
	BgHiGreen   string = "\x1b[102m" // VGA (85, 255, 85)
	BgHiYellow  string = "\x1b[103m" // VGA (255, 255, 85)
	BgHiBlue    string = "\x1b[104m" // VGA (85, 85, 255)
	BgHiMagenta string = "\x1b[105m" // VGA (255, 85, 255)
	BgHiCyan    string = "\x1b[106m" // VGA (85, 255, 255)
	BgHiWhite   string = "\x1b[107m" // VGA (255, 255, 255)

	M         string = "m" // \x1b[0m <- The M at the end to close off an ansi escape code
	Separator string = ";" // \x1b[38;5; <- The ; inbetween arguments

	Black   string = "0" // VGA (0, 0, 0)
	Red     string = "1" // VGA (170, 0, 0)
	Green   string = "2" // VGA (0, 170, 0)
	Yellow  string = "3" // VGA (170, 85, 0)
	Blue    string = "4" // VGA (0, 0, 170)
	Magenta string = "5" // VGA (170, 0, 170)
	Cyan    string = "6" // VGA (0, 170, 170)
	White   string = "7" // VGA (170, 170, 170)
)
