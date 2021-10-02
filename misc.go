package vera

import (
	"strconv"
	"strings"
)

const SlotMin = 1
const SlotMax = 64

// Installed returns whether VeraCrypt is installed or not, as well es the installed
// version, e.g. VeraCrypt 1.24-Update7
func Installed() (string, bool) {
	cmd, stdout, _ := newCommand(version)
	// cmd.Run will return "veracrypt: command not found" if veracrypt is not
	// installed on this system
	if err := cmd.Run(); err != nil {
		return "", false
	}

	// the stdout content has a newline, so we simply trim it away
	return strings.Trim(stdout.String(), "\n"), true
}

// slotValid returns an ErrParameterIncorrect error if the supplied slot is not valid (at least SlotMin and at
// most SlotMax). Using this func can prevent an unnecessary cli call
func slotValid(slot uint8) error {
	if slot < SlotMin || slot > SlotMax {
		msg := ErrParameterIncorrect.Error() + ": " + strconv.Itoa(int(slot))
		return Error{cause: ErrParameterIncorrect, message: msg}
	}

	return nil
}
