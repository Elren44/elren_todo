package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

//NewForm initializes a form struct
func NewForm(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Required checks for required fiels
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "Это поле не может быть пустым")
		}
	}
}

//Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//MinLenght checks for string minimum lenght
func (f *Form) MinLenght(field string, lenght int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < lenght {
		f.Errors.Add(field, fmt.Sprintf("Минимум %d символов", lenght))
		return false
	}
	return true
}

//EqualPasswords checks if passwords not equal
func (f *Form) EqualPasswords(r *http.Request) {
	if r.Form.Get("password") != r.Form.Get("password2") {
		f.Errors.Add("password2", "Пароли не совпадают")
	}
}

//IsEmail checks if form field is email
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Неправильный адрес почты")
	}
}

func (f *Form) ExistUser() {
	f.Errors.Add("email", "Пользователь с таким Email уже существует.")
}
