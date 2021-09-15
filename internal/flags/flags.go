package flags

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/pflag"
)

func ParseArgs(data interface{}, args []string) {
	flagset := pflag.NewFlagSet("options", pflag.ExitOnError)

	v := reflect.ValueOf(data).Elem()
	t := v.Type()

	// Klio flags
	flagset.CountP("verbose", "v", "More verbose output (-vv... to further increase verbosity)")
	flagset.StringP("log-level", "", "info", "Set logs level: disable, fatal, error, warn, info, verbose, debug, spam")
	flagset.MarkHidden("verbose")
	flagset.MarkHidden("log-level")

	for i := 0; i < t.NumField(); i++ {
		fv := v.Field(i)
		ft := t.Field(i)

		flag := ft.Tag.Get("flag")
		help := ft.Tag.Get("help")
		alias := ft.Tag.Get("alias")

		if flag == "" {
			continue
		}

		switch fv.Kind() {
		case reflect.String:
			flagset.StringVarP(fv.Addr().Interface().(*string), flag, alias, fv.String(), help)
		case reflect.Bool:
			flagset.BoolVarP(fv.Addr().Interface().(*bool), flag, alias, fv.Bool(), help)
		case reflect.Map:
			flagset.StringToStringVarP(fv.Addr().Interface().(*map[string]string), flag, alias, map[string]string{}, help)
		case reflect.Slice:
			flagset.StringSliceVarP(fv.Addr().Interface().(*[]string), flag, alias, []string{}, help)
		case reflect.Int:
			flagset.IntVarP(fv.Addr().Interface().(*int), flag, alias, int(fv.Int()), help)
		default:
			panic(fmt.Sprintf("unsupported kind %s for flag --%s", fv.Kind(), flag))
		}
	}

	flagset.Parse(args)

	for i := 0; i < t.NumField(); i++ {
		fv := v.Field(i)
		ft := t.Field(i)

		flag := ft.Tag.Get("flag")
		required := ft.Tag.Get("required")

		if flag != "" && required == "true" {
			if fv.IsZero() {
				exitWithError(flagset, fmt.Errorf("missing required flag: --%s", flag))

			}
		}
	}
}

func exitWithError(flagset *pflag.FlagSet, err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	fmt.Fprintf(os.Stderr, "Usage of options:\n")
	flagset.PrintDefaults()
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(2)
}
