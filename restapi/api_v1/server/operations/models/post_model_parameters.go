// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostModelParams creates a new PostModelParams object
// with the default values initialized.
func NewPostModelParams() PostModelParams {
	var (
		versionDefault = string("2017-02-13")
	)
	return PostModelParams{
		Version: versionDefault,
	}
}

// PostModelParams contains all the bound params for the post model operation
// typically these are obtained from a http.Request
//
// swagger:parameters postModel
type PostModelParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The manifest providing configuration for the deep learning model, the training data and the training execution.
	  Required: true
	  In: formData
	*/
	Manifest runtime.File
	/*The deep learning model code as compressed archive (ZIP).
	  In: formData
	*/
	ModelDefinition *runtime.File
	/*The release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format.
	  Required: true
	  In: query
	  Default: "2017-02-13"
	*/
	Version string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PostModelParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return err
		} else if err := r.ParseForm(); err != nil {
			return err
		}
	}

	manifest, manifestHeader, err := r.FormFile("manifest")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "manifest", err))
	} else {
		o.Manifest = runtime.File{Data: manifest, Header: manifestHeader}
	}

	modelDefinition, modelDefinitionHeader, err := r.FormFile("model_definition")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "modelDefinition", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else {
		o.ModelDefinition = &runtime.File{Data: modelDefinition, Header: modelDefinitionHeader}
	}

	qVersion, qhkVersion, _ := qs.GetOK("version")
	if err := o.bindVersion(qVersion, qhkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostModelParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("version", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("version", "query", raw); err != nil {
		return err
	}

	o.Version = raw

	return nil
}