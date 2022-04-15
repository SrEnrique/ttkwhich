package ttkwhich

import (
	"errors"
	"os"
	"runtime"
)

// Hello returns a greeting for the named person.
func Platform() string {
	// Return a greeting that embeds the name in a message.

	message := runtime.GOOS
	return message
}

// which search in default paths

func Which(cmd string) string {
	var result = ""
	default_path_linux := []string{
		"/usr/local/sbin",
		"/usr/local/bin",
		"/usr/sbin",
		"/usr/bin",
		"/sbin",
		"/bin",
		"/usr/games",
		"/usr/local/games",
		"/snap/bin"}

	if runtime.GOOS == "linux" {
		result = Which_linux(cmd, default_path_linux)
	} else if runtime.GOOS == "windows" {
		result = "Validacion windows"
	}

	return result

}

func Which_in_path(cmd string, paths []string) string {
	var result = ""

	if runtime.GOOS == "linux" {
		result = Which_linux(cmd, paths)
	} else if runtime.GOOS == "windows" {
		result = "Validacion windows"
	}

	return result

}

func Which_linux(cmd string, paths []string) string {
	cmd_path := ""
	for _, path := range paths {
		//fmt.Println(path + "/" + cmd)
		if _, err := os.Stat(path + "/" + cmd); err == nil {
			cmd_path = path + "/" + cmd
		} else if errors.Is(err, os.ErrNotExist) {
			// path/to/whatever does *not* exist

		} else {
			// Schrodinger: file may or may not exist. See err for details.
			// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		}
	}

	return cmd_path
}
