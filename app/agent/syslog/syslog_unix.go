// +build !windows,!nacl,!plan9

package syslog

import (
	"context"
	"io"
	"log/syslog"
	"time"

	"github.com/go-pkgz/repeater"
)

// GetWriter returns syslog writer for non-win platform
func GetWriter(ctx context.Context, syslogHost, syslogPrefix, containerName string) (res io.WriteCloser, err error) {
	// try UDP syslog first
	e := repeater.NewDefault(10, time.Second).Do(ctx, func() error {
		res, err = syslog.Dial("udp4", syslogHost, syslog.LOG_WARNING|syslog.LOG_DAEMON, syslogPrefix+containerName)
		return err
	})

	// try TCP if UDP failed
	if e != nil {
		e = repeater.NewDefault(10, time.Second).Do(ctx, func() error {
			res, err = syslog.Dial("tcp4", syslogHost, syslog.LOG_WARNING|syslog.LOG_DAEMON, syslogPrefix+containerName)
			return err
		})
	}
	return res, e
}

// IsSupported indicates if the platform supports syslog
func IsSupported() bool {
	return true
}
