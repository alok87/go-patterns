package capture

import (
	"bytes"
	"io"
	"os"
)

// Stdout captures stdout.
func Stdout(f func()) string {
	c := &capturer{stdout: true}
	
	return c.capture(f)
}

// Stderr captures stderr.
func Stderr(f func()) string {
	c := &capturer{stderr: true}
	
	return c.capture(f)
}

// Output captures stdout and stderr.
func Output(f func()) string {
	c := &capturer{stdout: true, stderr: true}
	
	return c.capture(f)
}

// capturer has flags whether capture stdout/stderr or both.
type capturer struct {
	stdout bool
	stderr bool
}

func (c *capturer) capture(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	if c.stdout {
		stdout := os.Stdout
		os.Stdout = w
		defer func() {
			os.Stdout = stdout
		}()
	}

	if c.stderr {
		stderr := os.Stderr
		os.Stderr = w
		defer func() {
			os.Stderr = stderr
		}()
	}

	f()
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
