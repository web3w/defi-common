package ulog

import (
	"fmt"

	"github.com/getsentry/raven-go"
	"github.com/pkg/errors"
)

//-------------------------------- Error ---------------------------------------

func ReportError(args ...interface{}) {
	if raven.ProjectID() == "" {
		return
	}
	// asynchronous send
	if err, ok := args[0].(error); ok {
		raven.CaptureError(err, nil)
	} else {
		raven.CaptureError(errors.New(fmt.Sprint(args...)), nil)
	}
}

func ReportErrorf(format string, args ...interface{}) {
	if raven.ProjectID() == "" {
		return
	}
	raven.CaptureError(fmt.Errorf(format, args...), nil)
}

func ReportErrorln(args ...interface{}) {
	if raven.ProjectID() == "" {
		return
	}
	raven.CaptureError(errors.New(fmt.Sprintln(args...)), nil)
}

func ReportErrorw(msg string, args ...interface{}) {
	if raven.ProjectID() == "" {
		return
	}
	// TODO: Consider calling raven.CaptureError directly, passing args as tags. The same for
	//       ReportPanicw and ReportFatalw.
	raven.CaptureError(errors.New(fmt.Sprintln(append([]interface{}{msg}, args...))), nil)
}

//-------------------------------- Panic ---------------------------------------

func ReportPanic(args ...interface{}) {
	if raven.ProjectID() == "" {
		return
	}

	// synchronous send
	if err, ok := args[0].(error); ok {
		raven.CaptureErrorAndWait(err, nil)
	} else {
		raven.CaptureErrorAndWait(errors.New(fmt.Sprint(args...)), nil)
	}
}

func ReportPanicf(format string, args ...interface{}) {
	if raven.ProjectID() == "" {
		return
	}
	raven.CaptureErrorAndWait(fmt.Errorf(format, args), nil)
}

func ReportPanicln(args ...interface{}) {
	if raven.ProjectID() == "" {
		return
	}
	raven.CaptureErrorAndWait(errors.New(fmt.Sprintln(args...)), nil)
}

func ReportPanicw(msg string, args ...interface{}) {
	if raven.ProjectID() == "" {
		return
	}
	raven.CaptureErrorAndWait(errors.New(fmt.Sprintln(append([]interface{}{msg}, args...))), nil)
}

//-------------------------------- Fatal ---------------------------------------

var (
	ReportFatal   = ReportPanic
	ReportFatalf  = ReportPanicf
	ReportFatalln = ReportPanicln
	ReportFatalw  = ReportPanicw
)

//-------------------------------- Warning -------------------------------------

var (
	ReportWarning   = ReportError
	ReportWarningf  = ReportErrorf
	ReportWarningln = ReportErrorln
	ReportWarningw  = ReportErrorw
)
