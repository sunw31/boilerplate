package slog

import (
	"context"
	"strconv"
	"strings"

	"cloud.google.com/go/logging"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/rs/zerolog"
	"github.com/valyala/fastjson"
	logpb "google.golang.org/genproto/googleapis/logging/v2"
	"google.golang.org/grpc/codes"
)

// NewStackdriverLogging returns a new StackdriverLoggingWriter.
func NewStackdriverLogging(
	ctx context.Context,
	projectID string,
	logID string,
	labels map[string]string,
	opts ...logging.LoggerOption,
) (*StackdriverLoggingWriter, *logging.Client, error) {
	const op = op + ".NewStackdriverLogging"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, nil, err
	}
	if projectID == "" {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'project' is ''"}
		return nil, nil, err
	}
	if logID == "" {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'project' is ''"}
		return nil, nil, err
	}
	v0 := &StackdriverLoggingWriter{
		projectID: projectID,
	}
	v1, err := logging.NewClient(ctx, v0.GetParentProjects())
	if err != nil {
		const op = op + ".logging.NewClient"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return nil, nil, err
	}
	err = v1.Ping(ctx)
	if err != nil {
		const op = op + ".logging.Client.Ping"
		err := &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
		return nil, nil, err
	}
	v0.Logger = v1.Logger(logID,
		// labels comes before opts so that any CommonLabels in opts take precedence.
		append([]logging.LoggerOption{logging.CommonLabels(labels)}, opts...)...,
	)
	return v0, v1, nil
}

// StackdriverLogging accepts pre-encoded JSON messages and writes them to Google Stackdriver Logging
// and maps Zerolog levels to Stackdriver levels.
// The labels argument is ignored if opts includes CommonLabels.
// The returned client should be closed before the program exits.
type StackdriverLoggingWriter struct {
	Logger          *logging.Logger
	parentProjects  string
	projectID       string
	traceIDTemplate string
}

// Write always returns len(p), nil.
func (w *StackdriverLoggingWriter) Write(p []byte) (int, error) {
	w.Logger.Log(*(NewEntry(p)))
	return len(p), nil
}

// WriteLevel implements zerolog.LevelWriter. It always returns len(p), nil.
func (w *StackdriverLoggingWriter) WriteLevel(level zerolog.Level, p []byte) (int, error) {
	severity := logging.Default

	// https://godoc.org/github.com/rs/zerolog#Level
	// https://godoc.org/cloud.google.com/go/logging#pkg-constants
	switch level {
	case zerolog.NoLevel:
		severity = logging.Default
	case zerolog.DebugLevel:
		severity = logging.Debug
	case zerolog.InfoLevel:
		severity = logging.Info
	case zerolog.WarnLevel:
		severity = logging.Warning
	case zerolog.ErrorLevel:
		severity = logging.Error
	case zerolog.FatalLevel:
		severity = logging.Critical
	case zerolog.PanicLevel:
		severity = logging.Alert
	}

	v := NewEntry(p)
	v.Severity = severity
	w.Logger.Log(*v)
	return len(p), nil
}

func (p *StackdriverLoggingWriter) Flush() error {
	return p.Logger.Flush()
}

// GetTraceIDTemplate returns a template string of the stackdriver traceID.
func (p *StackdriverLoggingWriter) GetTraceIDTemplate() string {
	return p.GetParentProjects() + "/traces/%s"
}

// GetParentProjects returns a string of parent projects.
// https://godoc.org/cloud.google.com/go/logging#NewClient
func (p *StackdriverLoggingWriter) GetParentProjects() string {
	return "projects/" + p.projectID
}

type rawJSON []byte

func (r rawJSON) MarshalJSON() ([]byte, error) { return []byte(r), nil }
func (r *rawJSON) UnmarshalJSON(b []byte) error {
	*r = rawJSON(b)
	return nil
}

func NewEntry(p []byte) *logging.Entry {
	v0 := logging.Entry{Payload: rawJSON(p)}
	v1, err := fastjson.ParseBytes(p)
	if err == nil {
		v0.SpanID = string(v1.GetStringBytes("spanID"))
		v0.Trace = string(v1.GetStringBytes("trace"))
		v0.TraceSampled = v1.GetBool("traceSampled")
		v2 := strings.SplitN(string(v1.GetStringBytes("caller")), ":", 2)
		v3 := strings.SplitAfterN(v2[0], "/github.com/", 2)
		v4 := &logpb.LogEntrySourceLocation{File: "github.com/" + v3[len(v3)-1]}
		if len(v2) == 2 {
			v5, err := strconv.ParseInt(v2[1], 10, 64)
			if err == nil {
				v4.Line = v5
			}
		}
		v0.SourceLocation = v4
	}
	return &v0
}