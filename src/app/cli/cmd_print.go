package cli

import (
	"fmt"
	"klog/app"
	"klog/parser"
	"klog/service"
)

type Print struct {
	FilterArgs
	MultipleFilesArgs
	Sort bool `short:"s" name:"sort" help:"Sort output by date (from oldest to latest)"`
}

func (args *Print) Run(ctx *app.Context) error {
	rs, err := ctx.RetrieveRecords(args.File...)
	if err != nil {
		return prettifyError(err)
	}
	rs = service.FindFilter(rs, args.FilterArgs.toFilter())
	if args.Sort {
		rs = service.Sort(rs, true)
	}
	fmt.Println("\n" + parser.SerialiseRecords(rs, styler))
	return nil
}
