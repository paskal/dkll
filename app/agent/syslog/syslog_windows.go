// +build windows nacl plan9

package syslog

import (
	"context"
	"errors"
	"io"
)

func GetWriter(ctx context.Context, host, proto, prefix, containerName string) (io.WriteCloser, error) {
	return nil, errors.New("syslog is not supported on this os")
}

func IsSupported() bool {
	return false
}
