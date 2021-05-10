package journal

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"
)

// Priority of a journal message
type Priority int

const (
	PriEmerg Priority = iota
	PriAlert
	PriCrit
	PriErr
	PriWarning
	PriNotice
	PriInfo
	PriDebug
)

var (
	// This can be overridden at build-time:
	// https://github.com/golang/go/wiki/GcToolchainTricks#including-build-information-in-the-executable
	journalSocket = "/run/systemd/journal/socket"

	// unixConnPtr atomically holds the local unconnected Unix-domain socket.
	// Concrete safe pointer type: *net.UnixConn
	unixConnPtr unsafe.Pointer
	// onceConn ensures that unixConnPtr is initialized exactly once.
	onceConn sync.Once
)

func init() {
	onceConn.Do(initConn)
}

// Enabled checks whether the local systemd journal is available for logging.
func Enabled() bool {
	onceConn.Do(initConn)

	if (*net.UnixConn)(atomic.LoadPointer(&unixConnPtr)) == nil {
		return false
	}

	if _, err := net.Dial("unixgram", journalSocket); err != nil {
		return false
	}

	return true
}

// Send a message to the local systemd journal. vars is a map of journald
// fields to values.  Fields must be composed of uppercase letters, numbers,
// and underscores, but must not start with an underscore. Within these
// restrictions, any arbitrary field name may be used.  Some names have special
// significance: see the journalctl documentation
// (http://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html)
// for more details.  vars may be nil.
func Send(message string, priority Priority, vars map[string]string) error {
	conn := (*net.UnixConn)(atomic.LoadPointer(&unixConnPtr))
	if conn == nil {
		return errors.New("could not initialize socket to journald")
	}

	socketAddr := &net.UnixAddr{
		Name: journalSocket,
		Net:  "unixgram",
	}

	data := new(bytes.Buffer)
	appendVariable(data, "PRIORITY", strconv.Itoa(int(priority)))
	appendVariable(data, "MESSAGE", message)
	for k, v := range vars {
		appendVariable(data, k, v)
	}

	_, _, err := conn.WriteMsgUnix(data.Bytes(), nil, socketAddr)
	if err == nil {
		return nil
	}
	if !isSocketSpaceError(err) {
		return err
	}

	// Large log entry, send it via tempfile and ancillary-fd.
	file, err := tempFd()
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, data)
	if err != nil {
		return err
	}
	rights := syscall.UnixRights(int(file.Fd()))
	_, _, err = conn.WriteMsgUnix([]byte{}, rights, socketAddr)
	if err != nil {
		return err
	}

	return nil
}

// Print prints a message to the local systemd journal using Send().
func Print(priority Priority, format string, a ...interface{}) error {
	return Send(fmt.Sprintf(format, a...), priority, nil)
}

func appendVariable(w io.Writer, name, value string) {
	if err := validVarName(name); err != nil {
		fmt.Fprintf(os.Stderr, "variable name %s contains invalid character, ignoring\n", name)
	}
	if strings.ContainsRune(value, '\n') {
		/* When the value contains a newline, we write:
		 * - the variable name, followed by a newline
		 * - the size (in 64bit little endian format)
		 * - the data, followed by a newline
		 */
		fmt.Fprintln(w, name)
		binary.Write(w, binary.LittleEndian, uint64(len(value)))
		fmt.Fprintln(w, value)
	} else {
		/* just write the variable and value all on one line */
		fmt.Fprintf(w, "%s=%s\n", name, value)
	}
}

// validVarName validates a variable name to make sure journald will accept it.
// The variable name must be in uppercase and consist only of characters,
// numbers and underscores, and may not begin with an underscore:
// https://www.freedesktop.org/software/systemd/man/sd_journal_print.html
func validVarName(name string) error {
	if name == "" {
		return errors.New("Empty variable name")
	} else if name[0] == '_' {
		return errors.New("Variable name begins with an underscore")
	}

	for _, c := range name {
		if !(('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '_') {
			return errors.New("Variable name contains invalid characters")
		}
	}
	return nil
}

// isSocketSpaceError checks whether the error is signaling
// an "overlarge message" condition.
func isSocketSpaceError(err error) bool {
	opErr, ok := err.(*net.OpError)
	if !ok || opErr == nil {
		return false
	}

	sysErr, ok := opErr.Err.(*os.SyscallError)
	if !ok || sysErr == nil {
		return false
	}

	return sysErr.Err == syscall.EMSGSIZE || sysErr.Err == syscall.ENOBUFS
}

// tempFd creates a temporary, unlinked file under `/dev/shm`.
func tempFd() (*os.File, error) {
	file, err := ioutil.TempFile("/dev/shm/", "journal.XXXXX")
	if err != nil {
		return nil, err
	}
	err = syscall.Unlink(file.Name())
	if err != nil {
		return nil, err
	}
	return file, nil
}

// initConn initializes the global `unixConnPtr` socket.
// It is meant to be called exactly once, at program startup.
func initConn() {
	autobind, err := net.ResolveUnixAddr("unixgram", "")
	if err != nil {
		return
	}

	sock, err := net.ListenUnixgram("unixgram", autobind)
	if err != nil {
		return
	}

	atomic.StorePointer(&unixConnPtr, unsafe.Pointer(sock))
}