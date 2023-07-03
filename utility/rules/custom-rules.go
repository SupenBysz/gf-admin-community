package rules

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gvalid"
	"reflect"
)

func RequiredLicense() {
	gvalid.RegisterRule("required-license", func(ctx context.Context, in gvalid.RuleFuncInput) error {
		license := ghttp.RequestFromCtx(ctx).Get("license")
		if license == nil {
			return nil
		}

		licenseFields, _ := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         &sys_model.PersonLicense{},
			RecursiveOption: 0,
		})

		rulesMap := gmap.StrStrMap{}

		for _, field := range licenseFields {
			tag := field.Tag("v")

			tag = gstr.Replace(tag, "-license", "")

			if tag == "" || gstr.HasPrefix(tag, "#") {
				continue
			}

			rulesMap.Set(field.TagJsonName(), tag)

			data := gmap.NewStrAnyMapFrom(license.MapStrAny())

			tmpArr := gstr.Split(tag, "#")

			value := data.GetVar(field.TagJsonName()).Val()

			vObj := g.Validator()

			if value == nil && field.Field.Type == reflect.TypeOf("") {
				value = ""
			}

			vObj = vObj.Data(value).Rules(tmpArr[0])

			if len(tmpArr) == 2 {
				vObj = vObj.Messages(tmpArr[1])
			}

			if err := vObj.Run(context.Background()); err != nil {
				fmt.Println("check value err:", err)
				return err
			}
		}

		return nil
	})
}
