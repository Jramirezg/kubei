// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetApplicationResourcesURL generates an URL for the get application resources operation
type GetApplicationResourcesURL struct {
	ApplicationID                               *string
	ApplicationsGte                             *int64
	ApplicationsIsNot                           []int64
	ApplicationsIs                              []int64
	ApplicationsLte                             *int64
	CurrentRuntimeScan                          *bool
	PackageID                                   *string
	PackagesGte                                 *int64
	PackagesIsNot                               []int64
	PackagesIs                                  []int64
	PackagesLte                                 *int64
	Page                                        int64
	PageSize                                    int64
	ReportingSBOMAnalyzersContainElements       []string
	ReportingSBOMAnalyzersDoesntContainElements []string
	ResourceHashContains                        []string
	ResourceHashEnd                             *string
	ResourceHashIsNot                           []string
	ResourceHashIs                              []string
	ResourceHashStart                           *string
	ResourceNameContains                        []string
	ResourceNameEnd                             *string
	ResourceNameIsNot                           []string
	ResourceNameIs                              []string
	ResourceNameStart                           *string
	ResourceTypeIs                              []string
	SortDir                                     *string
	SortKey                                     string
	VulnerabilitySeverityGte                    *string
	VulnerabilitySeverityLte                    *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetApplicationResourcesURL) WithBasePath(bp string) *GetApplicationResourcesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetApplicationResourcesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetApplicationResourcesURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/applicationResources"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var applicationIDQ string
	if o.ApplicationID != nil {
		applicationIDQ = *o.ApplicationID
	}
	if applicationIDQ != "" {
		qs.Set("applicationID", applicationIDQ)
	}

	var applicationsGteQ string
	if o.ApplicationsGte != nil {
		applicationsGteQ = swag.FormatInt64(*o.ApplicationsGte)
	}
	if applicationsGteQ != "" {
		qs.Set("applications[gte]", applicationsGteQ)
	}

	var applicationsIsNotIR []string
	for _, applicationsIsNotI := range o.ApplicationsIsNot {
		applicationsIsNotIS := swag.FormatInt64(applicationsIsNotI)
		if applicationsIsNotIS != "" {
			applicationsIsNotIR = append(applicationsIsNotIR, applicationsIsNotIS)
		}
	}

	applicationsIsNot := swag.JoinByFormat(applicationsIsNotIR, "")

	if len(applicationsIsNot) > 0 {
		qsv := applicationsIsNot[0]
		if qsv != "" {
			qs.Set("applications[isNot]", qsv)
		}
	}

	var applicationsIsIR []string
	for _, applicationsIsI := range o.ApplicationsIs {
		applicationsIsIS := swag.FormatInt64(applicationsIsI)
		if applicationsIsIS != "" {
			applicationsIsIR = append(applicationsIsIR, applicationsIsIS)
		}
	}

	applicationsIs := swag.JoinByFormat(applicationsIsIR, "")

	if len(applicationsIs) > 0 {
		qsv := applicationsIs[0]
		if qsv != "" {
			qs.Set("applications[is]", qsv)
		}
	}

	var applicationsLteQ string
	if o.ApplicationsLte != nil {
		applicationsLteQ = swag.FormatInt64(*o.ApplicationsLte)
	}
	if applicationsLteQ != "" {
		qs.Set("applications[lte]", applicationsLteQ)
	}

	var currentRuntimeScanQ string
	if o.CurrentRuntimeScan != nil {
		currentRuntimeScanQ = swag.FormatBool(*o.CurrentRuntimeScan)
	}
	if currentRuntimeScanQ != "" {
		qs.Set("currentRuntimeScan", currentRuntimeScanQ)
	}

	var packageIDQ string
	if o.PackageID != nil {
		packageIDQ = *o.PackageID
	}
	if packageIDQ != "" {
		qs.Set("packageID", packageIDQ)
	}

	var packagesGteQ string
	if o.PackagesGte != nil {
		packagesGteQ = swag.FormatInt64(*o.PackagesGte)
	}
	if packagesGteQ != "" {
		qs.Set("packages[gte]", packagesGteQ)
	}

	var packagesIsNotIR []string
	for _, packagesIsNotI := range o.PackagesIsNot {
		packagesIsNotIS := swag.FormatInt64(packagesIsNotI)
		if packagesIsNotIS != "" {
			packagesIsNotIR = append(packagesIsNotIR, packagesIsNotIS)
		}
	}

	packagesIsNot := swag.JoinByFormat(packagesIsNotIR, "")

	if len(packagesIsNot) > 0 {
		qsv := packagesIsNot[0]
		if qsv != "" {
			qs.Set("packages[isNot]", qsv)
		}
	}

	var packagesIsIR []string
	for _, packagesIsI := range o.PackagesIs {
		packagesIsIS := swag.FormatInt64(packagesIsI)
		if packagesIsIS != "" {
			packagesIsIR = append(packagesIsIR, packagesIsIS)
		}
	}

	packagesIs := swag.JoinByFormat(packagesIsIR, "")

	if len(packagesIs) > 0 {
		qsv := packagesIs[0]
		if qsv != "" {
			qs.Set("packages[is]", qsv)
		}
	}

	var packagesLteQ string
	if o.PackagesLte != nil {
		packagesLteQ = swag.FormatInt64(*o.PackagesLte)
	}
	if packagesLteQ != "" {
		qs.Set("packages[lte]", packagesLteQ)
	}

	pageQ := swag.FormatInt64(o.Page)
	if pageQ != "" {
		qs.Set("page", pageQ)
	}

	pageSizeQ := swag.FormatInt64(o.PageSize)
	if pageSizeQ != "" {
		qs.Set("pageSize", pageSizeQ)
	}

	var reportingSBOMAnalyzersContainElementsIR []string
	for _, reportingSBOMAnalyzersContainElementsI := range o.ReportingSBOMAnalyzersContainElements {
		reportingSBOMAnalyzersContainElementsIS := reportingSBOMAnalyzersContainElementsI
		if reportingSBOMAnalyzersContainElementsIS != "" {
			reportingSBOMAnalyzersContainElementsIR = append(reportingSBOMAnalyzersContainElementsIR, reportingSBOMAnalyzersContainElementsIS)
		}
	}

	reportingSBOMAnalyzersContainElements := swag.JoinByFormat(reportingSBOMAnalyzersContainElementsIR, "")

	if len(reportingSBOMAnalyzersContainElements) > 0 {
		qsv := reportingSBOMAnalyzersContainElements[0]
		if qsv != "" {
			qs.Set("reportingSBOMAnalyzers[containElements]", qsv)
		}
	}

	var reportingSBOMAnalyzersDoesntContainElementsIR []string
	for _, reportingSBOMAnalyzersDoesntContainElementsI := range o.ReportingSBOMAnalyzersDoesntContainElements {
		reportingSBOMAnalyzersDoesntContainElementsIS := reportingSBOMAnalyzersDoesntContainElementsI
		if reportingSBOMAnalyzersDoesntContainElementsIS != "" {
			reportingSBOMAnalyzersDoesntContainElementsIR = append(reportingSBOMAnalyzersDoesntContainElementsIR, reportingSBOMAnalyzersDoesntContainElementsIS)
		}
	}

	reportingSBOMAnalyzersDoesntContainElements := swag.JoinByFormat(reportingSBOMAnalyzersDoesntContainElementsIR, "")

	if len(reportingSBOMAnalyzersDoesntContainElements) > 0 {
		qsv := reportingSBOMAnalyzersDoesntContainElements[0]
		if qsv != "" {
			qs.Set("reportingSBOMAnalyzers[doesntContainElements]", qsv)
		}
	}

	var resourceHashContainsIR []string
	for _, resourceHashContainsI := range o.ResourceHashContains {
		resourceHashContainsIS := resourceHashContainsI
		if resourceHashContainsIS != "" {
			resourceHashContainsIR = append(resourceHashContainsIR, resourceHashContainsIS)
		}
	}

	resourceHashContains := swag.JoinByFormat(resourceHashContainsIR, "")

	if len(resourceHashContains) > 0 {
		qsv := resourceHashContains[0]
		if qsv != "" {
			qs.Set("resourceHash[contains]", qsv)
		}
	}

	var resourceHashEndQ string
	if o.ResourceHashEnd != nil {
		resourceHashEndQ = *o.ResourceHashEnd
	}
	if resourceHashEndQ != "" {
		qs.Set("resourceHash[end]", resourceHashEndQ)
	}

	var resourceHashIsNotIR []string
	for _, resourceHashIsNotI := range o.ResourceHashIsNot {
		resourceHashIsNotIS := resourceHashIsNotI
		if resourceHashIsNotIS != "" {
			resourceHashIsNotIR = append(resourceHashIsNotIR, resourceHashIsNotIS)
		}
	}

	resourceHashIsNot := swag.JoinByFormat(resourceHashIsNotIR, "")

	if len(resourceHashIsNot) > 0 {
		qsv := resourceHashIsNot[0]
		if qsv != "" {
			qs.Set("resourceHash[isNot]", qsv)
		}
	}

	var resourceHashIsIR []string
	for _, resourceHashIsI := range o.ResourceHashIs {
		resourceHashIsIS := resourceHashIsI
		if resourceHashIsIS != "" {
			resourceHashIsIR = append(resourceHashIsIR, resourceHashIsIS)
		}
	}

	resourceHashIs := swag.JoinByFormat(resourceHashIsIR, "")

	if len(resourceHashIs) > 0 {
		qsv := resourceHashIs[0]
		if qsv != "" {
			qs.Set("resourceHash[is]", qsv)
		}
	}

	var resourceHashStartQ string
	if o.ResourceHashStart != nil {
		resourceHashStartQ = *o.ResourceHashStart
	}
	if resourceHashStartQ != "" {
		qs.Set("resourceHash[start]", resourceHashStartQ)
	}

	var resourceNameContainsIR []string
	for _, resourceNameContainsI := range o.ResourceNameContains {
		resourceNameContainsIS := resourceNameContainsI
		if resourceNameContainsIS != "" {
			resourceNameContainsIR = append(resourceNameContainsIR, resourceNameContainsIS)
		}
	}

	resourceNameContains := swag.JoinByFormat(resourceNameContainsIR, "")

	if len(resourceNameContains) > 0 {
		qsv := resourceNameContains[0]
		if qsv != "" {
			qs.Set("resourceName[contains]", qsv)
		}
	}

	var resourceNameEndQ string
	if o.ResourceNameEnd != nil {
		resourceNameEndQ = *o.ResourceNameEnd
	}
	if resourceNameEndQ != "" {
		qs.Set("resourceName[end]", resourceNameEndQ)
	}

	var resourceNameIsNotIR []string
	for _, resourceNameIsNotI := range o.ResourceNameIsNot {
		resourceNameIsNotIS := resourceNameIsNotI
		if resourceNameIsNotIS != "" {
			resourceNameIsNotIR = append(resourceNameIsNotIR, resourceNameIsNotIS)
		}
	}

	resourceNameIsNot := swag.JoinByFormat(resourceNameIsNotIR, "")

	if len(resourceNameIsNot) > 0 {
		qsv := resourceNameIsNot[0]
		if qsv != "" {
			qs.Set("resourceName[isNot]", qsv)
		}
	}

	var resourceNameIsIR []string
	for _, resourceNameIsI := range o.ResourceNameIs {
		resourceNameIsIS := resourceNameIsI
		if resourceNameIsIS != "" {
			resourceNameIsIR = append(resourceNameIsIR, resourceNameIsIS)
		}
	}

	resourceNameIs := swag.JoinByFormat(resourceNameIsIR, "")

	if len(resourceNameIs) > 0 {
		qsv := resourceNameIs[0]
		if qsv != "" {
			qs.Set("resourceName[is]", qsv)
		}
	}

	var resourceNameStartQ string
	if o.ResourceNameStart != nil {
		resourceNameStartQ = *o.ResourceNameStart
	}
	if resourceNameStartQ != "" {
		qs.Set("resourceName[start]", resourceNameStartQ)
	}

	var resourceTypeIsIR []string
	for _, resourceTypeIsI := range o.ResourceTypeIs {
		resourceTypeIsIS := resourceTypeIsI
		if resourceTypeIsIS != "" {
			resourceTypeIsIR = append(resourceTypeIsIR, resourceTypeIsIS)
		}
	}

	resourceTypeIs := swag.JoinByFormat(resourceTypeIsIR, "")

	if len(resourceTypeIs) > 0 {
		qsv := resourceTypeIs[0]
		if qsv != "" {
			qs.Set("resourceType[is]", qsv)
		}
	}

	var sortDirQ string
	if o.SortDir != nil {
		sortDirQ = *o.SortDir
	}
	if sortDirQ != "" {
		qs.Set("sortDir", sortDirQ)
	}

	sortKeyQ := o.SortKey
	if sortKeyQ != "" {
		qs.Set("sortKey", sortKeyQ)
	}

	var vulnerabilitySeverityGteQ string
	if o.VulnerabilitySeverityGte != nil {
		vulnerabilitySeverityGteQ = *o.VulnerabilitySeverityGte
	}
	if vulnerabilitySeverityGteQ != "" {
		qs.Set("vulnerabilitySeverity[gte]", vulnerabilitySeverityGteQ)
	}

	var vulnerabilitySeverityLteQ string
	if o.VulnerabilitySeverityLte != nil {
		vulnerabilitySeverityLteQ = *o.VulnerabilitySeverityLte
	}
	if vulnerabilitySeverityLteQ != "" {
		qs.Set("vulnerabilitySeverity[lte]", vulnerabilitySeverityLteQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetApplicationResourcesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetApplicationResourcesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetApplicationResourcesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetApplicationResourcesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetApplicationResourcesURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetApplicationResourcesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
