// Copyright 2020 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package extended contains an extended set of terminal descriptions.
// Applications desiring to have a better chance of Just Working by
// default should include this package.  This will significantly increase
// the size of the program.
package extended

import (
	// The following imports just register themselves --
	// these are the terminal types we aggregate in this package.
	_ "go.mau.fi/tcell/terminfo/a/aixterm"
	_ "go.mau.fi/tcell/terminfo/a/alacritty"
	_ "go.mau.fi/tcell/terminfo/a/ansi"
	_ "go.mau.fi/tcell/terminfo/b/beterm"
	_ "go.mau.fi/tcell/terminfo/c/cygwin"
	_ "go.mau.fi/tcell/terminfo/d/dtterm"
	_ "go.mau.fi/tcell/terminfo/e/emacs"
	_ "go.mau.fi/tcell/terminfo/f/foot"
	_ "go.mau.fi/tcell/terminfo/g/gnome"
	_ "go.mau.fi/tcell/terminfo/h/hpterm"
	_ "go.mau.fi/tcell/terminfo/k/konsole"
	_ "go.mau.fi/tcell/terminfo/k/kterm"
	_ "go.mau.fi/tcell/terminfo/l/linux"
	_ "go.mau.fi/tcell/terminfo/p/pcansi"
	_ "go.mau.fi/tcell/terminfo/r/rxvt"
	_ "go.mau.fi/tcell/terminfo/s/screen"
	_ "go.mau.fi/tcell/terminfo/s/simpleterm"
	_ "go.mau.fi/tcell/terminfo/s/sun"
	_ "go.mau.fi/tcell/terminfo/t/termite"
	_ "go.mau.fi/tcell/terminfo/t/tmux"
	_ "go.mau.fi/tcell/terminfo/v/vt100"
	_ "go.mau.fi/tcell/terminfo/v/vt102"
	_ "go.mau.fi/tcell/terminfo/v/vt220"
	_ "go.mau.fi/tcell/terminfo/v/vt320"
	_ "go.mau.fi/tcell/terminfo/v/vt400"
	_ "go.mau.fi/tcell/terminfo/v/vt420"
	_ "go.mau.fi/tcell/terminfo/v/vt52"
	_ "go.mau.fi/tcell/terminfo/w/wy50"
	_ "go.mau.fi/tcell/terminfo/w/wy60"
	_ "go.mau.fi/tcell/terminfo/w/wy99_ansi"
	_ "go.mau.fi/tcell/terminfo/x/xfce"
	_ "go.mau.fi/tcell/terminfo/x/xterm"
	_ "go.mau.fi/tcell/terminfo/x/xterm_kitty"
	_ "go.mau.fi/tcell/terminfo/x/xterm_termite"
)
