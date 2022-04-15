// Generated automatically.  DO NOT HAND-EDIT.

package dtterm

import "go.mau.fi/tcell/terminfo"

func init() {

	// CDE desktop terminal
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:         "dtterm",
		Columns:      80,
		Lines:        24,
		Colors:       8,
		Bell:         "\a",
		Clear:        "\x1b[H\x1b[J",
		ShowCursor:   "\x1b[?25h",
		HideCursor:   "\x1b[?25l",
		AttrOff:      "\x1b[m\x0f",
		Underline:    "\x1b[4m",
		Bold:         "\x1b[1m",
		Dim:          "\x1b[2m",
		Blink:        "\x1b[5m",
		Reverse:      "\x1b[7m",
		SetFg:        "\x1b[3%p1%dm",
		SetBg:        "\x1b[4%p1%dm",
		SetFgBg:      "\x1b[3%p1%d;4%p2%dm",
		ResetFgBg:    "\x1b[39;49m",
		PadChar:      "\x00",
		AltChars:     "``aaffggjjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:     "\x0e",
		ExitAcs:      "\x0f",
		EnableAcs:    "\x1b(B\x1b)0",
		SetCursor:    "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:  "\b",
		CursorUp1:    "\x1b[A",
		KeyUp:        "\x1b[A",
		KeyDown:      "\x1b[B",
		KeyRight:     "\x1b[C",
		KeyLeft:      "\x1b[D",
		KeyInsert:    "\x1b[2~",
		KeyDelete:    "\x1b[3~",
		KeyBackspace: "\b",
		KeyPgUp:      "\x1b[5~",
		KeyPgDn:      "\x1b[6~",
		KeyF1:        "\x1b[11~",
		KeyF2:        "\x1b[12~",
		KeyF3:        "\x1b[13~",
		KeyF4:        "\x1b[14~",
		KeyF5:        "\x1b[15~",
		KeyF6:        "\x1b[17~",
		KeyF7:        "\x1b[18~",
		KeyF8:        "\x1b[19~",
		KeyF9:        "\x1b[20~",
		KeyF10:       "\x1b[21~",
		KeyF11:       "\x1b[23~",
		KeyF12:       "\x1b[24~",
		KeyF13:       "\x1b[25~",
		KeyF14:       "\x1b[26~",
		KeyF15:       "\x1b[28~",
		KeyF16:       "\x1b[29~",
		KeyF17:       "\x1b[31~",
		KeyF18:       "\x1b[32~",
		KeyF19:       "\x1b[33~",
		KeyF20:       "\x1b[34~",
		KeyHelp:      "\x1b[28~",
		AutoMargin:   true,
	})
}
