package pprof

import (
	"context"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/pkg/profile"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/now"
	"github.com/michilu/boilerplate/service/slog"
)

const (
	op = "service/pprof"
)

func Run() error {
	const op = op + ".Run"
	runtime.SetBlockProfileRate(1)
	e := gin.Default()
	pprof.Register(e)
	err := e.Run(viper.GetString("service.pprof.addr"))
	if err != nil {
		const op = op + ".gin.Run"
		return &errs.Error{Op: op, Err: err}
	}
	return nil
}

func Profile(ctx context.Context) {
	const (
		op = op + ".Profile"
		s  = "assets/pprof"
	)
	before := s + "/cpu.pprof"
	d := viper.GetDuration("service.pprof.duration")
	if d == 0 {
		slog.Logger().Warn().Str("op", op).
			Err(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "service.pprof.duration"}).
			Dur("value", d).Msg("warnning")
		d = 5 * time.Minute
	}
	t := time.NewTicker(d)
	defer t.Stop()
	for {
		p := profile.Start(profile.ProfilePath(s))
		select {
		case <-ctx.Done():
			p.Stop()
			return
		case <-t.C:
		}
		p.Stop()
		after := before + "." + now.Now().UTC().Format(time.RFC3339)
		err := os.Rename(before, after)
		if err != nil {
			const op = op + ".os.Rename"
			slog.Logger().Error().Str("op", op).Err(err).Str("before", before).Str("after", after).Msg("error")
		} else {
			slog.Logger().Info().Str("op", op).Str("before", before).Str("after", after).Msg("rotated")
		}
	}
}
