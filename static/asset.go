package static

import (
	"fmt"
	"text/template"

	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/util"
)

// EmailAsset represets asset of email templates
type EmailAsset struct {
	templates map[string]*template.Template
}

// GetTemplate returns email template
func (a *EmailAsset) GetTemplate(name string) (*template.Template, error) {
	tmpl, ok := a.templates[name]
	if !ok {
		return nil, errors.WithStack(fmt.Errorf("template %v not found", name))
	}
	return tmpl, nil
}

func (a *EmailAsset) addTemplate(name string, tmpl *template.Template) {
	a.templates[name] = tmpl
}

// LoadEmailTemplates loads templates
func LoadEmailTemplates() (*EmailAsset, error) {
	b, err := getEmailsBox()
	if err != nil {
		return nil, err
	}

	as := &EmailAsset{
		templates: make(map[string]*template.Template),
	}

	for _, n := range b.List() {
		t, err := b.FindString(n)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		tmpl, err := template.New(n).Parse(t)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		as.addTemplate(n, tmpl)
	}

	return as, nil
}

func getEmailsBox() (b *packr.Box, err error) {
	defer func() {
		if e := util.ErrorFromRecover(recover()); e != nil {
			err = e
		}
	}()

	b = packr.New("emails", "./emails")
	return b, nil
}
