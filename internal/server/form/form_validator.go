package form

import (
	"reflect"

	jaLocale "github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	jaTranslations "github.com/go-playground/validator/v10/translations/ja"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AI1411/go-grpc-graphql/grpc"
	formError "github.com/AI1411/go-grpc-graphql/internal/server/form/error"
)

const (
	ValidationPrefecture = "prefecture"
	ValidationBloodType  = "bloodType"
)

type Validator struct {
	Form interface{}
}

func NewFormValidator(formRequest interface{}) *Validator {
	return &Validator{
		Form: formRequest,
	}
}

func (formValidator *Validator) Validate(structLevelFuncs ...validator.StructLevelFunc) error {
	validate := validator.New()
	localeJa := jaLocale.New()
	translator := ut.New(localeJa)
	trans, _ := translator.GetTranslator("ja")

	if err := jaTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
		return errors.WithStack(err)
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		fieldName := fld.Tag.Get("jaFieldName")
		if fieldName == "-" {
			return ""
		}
		return fieldName
	})

	err := registerValidations(validate)
	if err != nil {
		return err
	}

	for _, option := range structLevelFuncs {
		validate.RegisterStructValidation(option, formValidator.Form)
	}

	err = validate.Struct(formValidator.Form)
	if err == nil {
		return nil
	}

	var errMessages []*errdetails.BadRequest_FieldViolation
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		errMessages = append(errMessages, &errdetails.BadRequest_FieldViolation{
			Field:       e.StructField(),
			Description: formError.ConvertFieldErrorDescription(e.Translate(trans), e.Tag(), e.Field(), e.Param()),
		})
	}

	st := status.New(codes.InvalidArgument, "bad request")
	v := &errdetails.BadRequest{FieldViolations: errMessages}
	dt, _ := st.WithDetails(v)
	return dt.Err()
}

func registerValidations(validate *validator.Validate) error {
	err := validate.RegisterValidation(ValidationPrefecture, isPrefecture)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation(ValidationBloodType, isBloodType)
	if err != nil {
		return err
	}
	return nil
}

func isPrefecture(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().Int()
	switch grpc.Prefecture(fieldValue) {
	case grpc.Prefecture_PREFECTURE_NULL,
		grpc.Prefecture_HOKKAIDO,
		grpc.Prefecture_AOMORI,
		grpc.Prefecture_IWATE,
		grpc.Prefecture_MIYAGI,
		grpc.Prefecture_AKITA,
		grpc.Prefecture_YAMAGATA,
		grpc.Prefecture_FUKUSHIMA,
		grpc.Prefecture_IBARAKI,
		grpc.Prefecture_TOCHIGI,
		grpc.Prefecture_GUNMA,
		grpc.Prefecture_SAITAMA,
		grpc.Prefecture_CHIBA,
		grpc.Prefecture_TOKYO,
		grpc.Prefecture_KANAGAWA,
		grpc.Prefecture_NIIGATA,
		grpc.Prefecture_TOYAMA,
		grpc.Prefecture_ISHIKAWA,
		grpc.Prefecture_FUKUI,
		grpc.Prefecture_YAMANASHI,
		grpc.Prefecture_NAGANO,
		grpc.Prefecture_GIFU,
		grpc.Prefecture_SHIZUOKA,
		grpc.Prefecture_AICHI,
		grpc.Prefecture_MIE,
		grpc.Prefecture_SHIGA,
		grpc.Prefecture_KYOTO,
		grpc.Prefecture_OSAKA,
		grpc.Prefecture_HYOGO,
		grpc.Prefecture_NARA,
		grpc.Prefecture_WAKAYAMA,
		grpc.Prefecture_TOTTORI,
		grpc.Prefecture_SHIMANE,
		grpc.Prefecture_OKAYAMA,
		grpc.Prefecture_HIROSHIMA,
		grpc.Prefecture_YAMAGUCHI,
		grpc.Prefecture_TOKUSHIMA,
		grpc.Prefecture_KAGAWA,
		grpc.Prefecture_EHIME,
		grpc.Prefecture_KOCHI,
		grpc.Prefecture_FUKUOKA,
		grpc.Prefecture_SAGA,
		grpc.Prefecture_NAGASAKI,
		grpc.Prefecture_KUMAMOTO,
		grpc.Prefecture_OITA,
		grpc.Prefecture_MIYAZAKI,
		grpc.Prefecture_KAGOSHIMA,
		grpc.Prefecture_OKINAWA,
		grpc.Prefecture_OVERSEAS:
		return true
	}
	return false
}

func isBloodType(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().Int()
	switch grpc.BloodType(fieldValue) {
	case grpc.BloodType_BLOOD_TYPE_NULL,
		grpc.BloodType_A,
		grpc.BloodType_B,
		grpc.BloodType_O,
		grpc.BloodType_AB:
		return true
	}
	return false
}
