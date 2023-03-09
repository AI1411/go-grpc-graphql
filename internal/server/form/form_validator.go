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

	formError "github.com/AI1411/go-grpc-praphql/internal/server/form/error"
)

type formValidator struct {
	Form interface{}
}

func NewFormValidator(formRequest interface{}) *formValidator {
	return &formValidator{
		Form: formRequest,
	}
}

func (formValidator *formValidator) Validate(structLevelFuncs ...validator.StructLevelFunc) error {
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
	return nil
}
