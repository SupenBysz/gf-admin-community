package validator

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gvalid"
	"strings"
)

func RegisterServicePhone() {
	gvalid.RegisterRule("service-phone", func(ctx context.Context, in gvalid.RuleFuncInput) error {
		err := gvalid.New().Data(in.Value.String()).Rules("phone").Run(ctx)
		if err == nil {
			return nil
		}
		err = gvalid.New().Data(in.Value.String()).Rules("\\d{11}").Run(ctx)
		if (strings.HasPrefix(in.Value.String(), "400") || strings.HasPrefix(in.Value.String(), "800")) && err == nil {
			return nil
		}
		err = gvalid.New().Data(in.Value.String()).Rules("[0-9]{5}").Run(ctx)
		if strings.HasPrefix(in.Value.String(), "95") && err == nil {
			return nil
		}
		err = gvalid.New().Data(in.Value.String()).Rules("telephone").Run(ctx)
		if err == nil {
			return nil
		}

		if in.Message != "" {
			return gerror.New(in.Message)
		}

		return gerror.New("The ServicePhone value `" + in.Value.String() + "` is not a valid ServicePhone number")
	})
}
