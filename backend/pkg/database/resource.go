// Copyright © 2022 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package database

import (
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/cisco-open/kubei/api/server/models"
	"github.com/cisco-open/kubei/api/server/restapi/operations"
	"github.com/cisco-open/kubei/backend/pkg/types"
	runtime_scan_models "github.com/cisco-open/kubei/runtime_scan/api/server/models"
	"github.com/cisco-open/kubei/shared/pkg/utils/slice"
)

const (
	resourceTableName = "resources"
	resourceViewName  = "resources_view"

	// NOTE: when changing one of the column names change also the gorm label in Resource.
	columnResourceID        = "id"
	columnResourceName      = "name"
	columnResourceHash      = "hash"
	columnResourceType      = "type"
	columnResourceAnalyzers = "reporting_analyzers"

	// NOTE: when changing one of the column names change also the gorm label in ResourceView.
	columnResourceViewApplications = "applications"
	columnResourceViewPackages     = "packages"
)

type Resource struct {
	ID string `gorm:"primarykey" faker:"-"` // consists of the resource hash

	Hash               string             `json:"hash,omitempty" gorm:"column:hash" faker:"oneof: hash1, hash2, hash3"`
	Name               string             `json:"name,omitempty" gorm:"column:name" faker:"oneof: resource1, resource2, resource3"`
	Type               types.ResourceType `json:"type,omitempty" gorm:"column:type" faker:"oneof: IMAGE, DIRECTORY, FILE"`
	SbomID             string             `json:"sbom_id,omitempty" gorm:"column:sbom_id" faker:"oneof: smobID1, smobID2, smobID3"`
	ReportingAnalyzers string             `json:"reporting_analyzers,omitempty" gorm:"column:reporting_analyzers" faker:"oneof: |analyzer1|, |analyzer1||analyzer2|"`
	Packages           []Package          `json:"packages,omitempty" gorm:"many2many:resource_packages;" faker:"-"`
}

type ResourceView struct {
	Resource
	Applications int `json:"applications,omitempty" gorm:"column:applications"`
	Packages     int `json:"packages,omitempty" gorm:"column:packages"`
	SeverityCounters
}

type GetApplicationResourcesParams struct {
	operations.GetApplicationResourcesParams
	// List of application IDs that were affected by the last runtime scan.
	RuntimeScanApplicationIDs []string
}

//go:generate $GOPATH/bin/mockgen -destination=./mock_resource_table.go -package=database github.com/cisco-open/kubei/backend/pkg/database ResourceTable
type ResourceTable interface {
	Create(resource *Resource) error
	GetApplicationResourcesAndTotal(params GetApplicationResourcesParams) ([]ResourceView, int64, error)
	GetApplicationResource(id string) (*models.ApplicationResourceEx, error)
	GetDBResource(id string, shouldGetRelationships bool) (*Resource, error)
	Count(filters *CountFilters) (int64, error)
	GetMostVulnerable(limit int) ([]*models.ApplicationResource, error)
}

type ResourceTableHandler struct {
	resourcesTable *gorm.DB
	resourcesView  *gorm.DB
	licensesView   *gorm.DB
	IDsView        IDsView
}

func (Resource) TableName() string {
	return resourceTableName
}

func CreateResourceFromVulnerabilityScan(resourceVulnerabilityScan *types.ResourceVulnerabilityScan, params *TransactionParams) *Resource {
	pkgIDToPackage := make(map[string]Package)

	resource := CreateResource(resourceVulnerabilityScan.Resource)

	for _, pkgVul := range resourceVulnerabilityScan.PackageVulnerabilities {
		vulnerability := CreateVulnerability(pkgVul, params)
		pkg := CreatePackage(pkgVul.Package, []Vulnerability{vulnerability})

		// update fix version on transaction params
		if pkgVul.FixVersion != "" {
			params.FixVersions[CreatePkgVulID(pkg.ID, vulnerability.ID)] = pkgVul.FixVersion
		}

		if p, ok := pkgIDToPackage[pkg.ID]; ok {
			// package exist, only append vulnerabilities
			p.Vulnerabilities = append(p.Vulnerabilities, pkg.Vulnerabilities...)
			pkgIDToPackage[pkg.ID] = p
		} else {
			pkgIDToPackage[pkg.ID] = *pkg
		}

		// update scanners on transaction params
		if len(pkgVul.Scanners) > 0 {
			resourcePkgID := CreateResourcePkgID(resource.ID, pkg.ID)
			params.Scanners[resourcePkgID] = slice.RemoveStringDuplicates(append(params.Scanners[resourcePkgID], pkgVul.Scanners...))
		}
	}

	pkgs := make([]Package, 0, len(pkgIDToPackage))
	for _, p := range pkgIDToPackage {
		pkgs = append(pkgs, p)
	}

	return resource.WithPackages(pkgs)
}

func CreateResourceFromRuntimeContentAnalysis(resourceContentAnalysis *runtime_scan_models.ResourceContentAnalysis, params *TransactionParams) *Resource {
	var resourceAnalyzers []string

	resource := CreateResource(types.ResourceInfoFromRuntimeScan(resourceContentAnalysis.Resource))

	pkgs := make([]Package, len(resourceContentAnalysis.Packages))
	for i, p := range resourceContentAnalysis.Packages {
		pkg := CreatePackageFromRuntimeContentAnalysis(p.Package)
		pkgs[i] = *pkg

		// update analyzers on transaction params
		if len(p.Analyzers) > 0 {
			params.Analyzers[CreateResourcePkgID(resource.ID, pkg.ID)] = p.Analyzers
			resourceAnalyzers = append(resourceAnalyzers, p.Analyzers...)
		}
	}

	return resource.WithPackages(pkgs).WithAnalyzers(slice.RemoveStringDuplicates(resourceAnalyzers))
}

func CreateResourceFromContentAnalysis(resourceContentAnalysis *models.ResourceContentAnalysis, params *TransactionParams) *Resource {
	var resourceAnalyzers []string

	resource := CreateResource(types.ResourceInfoFromBackendAPI(resourceContentAnalysis.Resource))

	pkgs := make([]Package, len(resourceContentAnalysis.Packages))
	for i, p := range resourceContentAnalysis.Packages {
		pkg := CreatePackageFromContentAnalysis(p.Package)
		pkgs[i] = *pkg

		// update analyzers on transaction params
		if len(p.Analyzers) > 0 {
			params.Analyzers[CreateResourcePkgID(resource.ID, pkg.ID)] = p.Analyzers
			resourceAnalyzers = append(resourceAnalyzers, p.Analyzers...)
		}
	}

	return resource.WithPackages(pkgs).WithAnalyzers(slice.RemoveStringDuplicates(resourceAnalyzers))
}

func CreateResource(info *types.ResourceInfo) *Resource {
	return &Resource{
		ID:   CreateResourceID(info),
		Hash: info.ResourceHash,
		Name: info.ResourceName,
		Type: info.ResourceType,
		// SbomID:             "", TODO: SBOM support
	}
}

func (r *Resource) WithPackages(packages []Package) *Resource {
	r.Packages = packages
	return r
}

func (r *Resource) WithAnalyzers(analyzers []string) *Resource {
	r.ReportingAnalyzers = ArrayToDBArray(analyzers)
	return r
}

func CreateResourceID(info *types.ResourceInfo) string {
	return uuid.NewV5(uuid.Nil, info.ResourceHash).String()
}

func (r *ResourceTableHandler) Create(resource *Resource) error {
	if err := r.resourcesTable.Create(resource).Error; err != nil {
		return fmt.Errorf("failed to create resource: %v", err)
	}
	return nil
}

func (r *ResourceTableHandler) setResourcesFilters(params GetApplicationResourcesParams) (*gorm.DB, error) {
	tx := r.resourcesView

	// name filters
	tx = FilterIs(tx, columnResourceName, params.ResourceNameIs)
	tx = FilterIsNot(tx, columnResourceName, params.ResourceNameIsNot)
	tx = FilterContains(tx, columnResourceName, params.ResourceNameContains)
	tx = FilterStartsWith(tx, columnResourceName, params.ResourceNameStart)
	tx = FilterEndsWith(tx, columnResourceName, params.ResourceNameEnd)

	// hash filters
	tx = FilterIs(tx, columnResourceHash, params.ResourceHashIs)
	tx = FilterIsNot(tx, columnResourceHash, params.ResourceHashIsNot)
	tx = FilterContains(tx, columnResourceHash, params.ResourceHashContains)
	tx = FilterStartsWith(tx, columnResourceHash, params.ResourceHashStart)
	tx = FilterEndsWith(tx, columnResourceHash, params.ResourceHashEnd)

	// type filter
	tx = FilterIs(tx, columnResourceType, params.ResourceTypeIs)

	// analyzers filter
	tx = FilterArrayContains(tx, columnResourceAnalyzers, params.ReportingSBOMAnalyzersContainElements)
	tx = FilterArrayDoesntContain(tx, columnResourceAnalyzers, params.ReportingSBOMAnalyzersDoesntContainElements)

	// applications filter
	tx = FilterIsNumber(tx, columnResourceViewApplications, params.ApplicationsIs)
	tx = FilterIsNotNumber(tx, columnResourceViewApplications, params.ApplicationsIsNot)
	tx = FilterGte(tx, columnResourceViewApplications, params.ApplicationsGte)
	tx = FilterLte(tx, columnResourceViewApplications, params.ApplicationsLte)

	// packages filter
	tx = FilterIsNumber(tx, columnResourceViewPackages, params.PackagesIs)
	tx = FilterIsNotNumber(tx, columnResourceViewPackages, params.PackagesIsNot)
	tx = FilterGte(tx, columnResourceViewPackages, params.PackagesGte)
	tx = FilterLte(tx, columnResourceViewPackages, params.PackagesLte)

	// vulnerabilities filter
	tx = SeverityFilterGte(tx, columnSeverityCountersHighestSeverity, params.VulnerabilitySeverityGte)
	tx = SeverityFilterLte(tx, columnSeverityCountersHighestSeverity, params.VulnerabilitySeverityLte)

	// system filter
	ids, err := r.getResourceIDs(params)
	if err != nil {
		return nil, fmt.Errorf("failed to get resource IDs: %v", err)
	}
	tx = FilterIs(tx, columnResourceID, ids)

	return tx, nil
}

func ApplicationResourceFromDB(view *ResourceView) *models.ApplicationResource {
	return &models.ApplicationResource{
		Applications:           uint32(view.Applications),
		ID:                     view.ID,
		Packages:               uint32(view.Packages),
		ReportingSBOMAnalyzers: DBArrayToArray(view.ReportingAnalyzers),
		ResourceHash:           view.Hash,
		ResourceName:           view.Name,
		ResourceType:           types.ResourceTypeToModels(view.Type),
		Vulnerabilities:        getVulnerabilityCount(view.SeverityCounters),
	}
}

func (r *ResourceTableHandler) GetApplicationResource(id string) (*models.ApplicationResourceEx, error) {
	var resourcesView ResourceView

	if err := r.resourcesView.Where(resourceViewName+"."+columnResourceID+" = ?", id).First(&resourcesView).Error; err != nil {
		return nil, fmt.Errorf("failed to get resource by id %q: %v", id, err)
	}

	licenses, err := r.getLicenses(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get licenses by id %q: %v", id, err)
	}

	return &models.ApplicationResourceEx{
		ApplicationResource: ApplicationResourceFromDB(&resourcesView),
		Licenses:            licenses,
	}, nil
}

func (r *ResourceTableHandler) GetDBResource(id string, shouldGetRelationships bool) (*Resource, error) {
	var resource Resource

	tx := r.resourcesTable.
		Where(resourceTableName+"."+columnResourceID+" = ?", id)

	if shouldGetRelationships {
		tx.Preload("Packages.Vulnerabilities").Preload(clause.Associations)
	}

	if err := tx.First(&resource).Error; err != nil {
		return nil, fmt.Errorf("failed to get resource by id %q: %v", id, err)
	}

	return &resource, nil
}

func (r *ResourceTableHandler) GetMostVulnerable(limit int) ([]*models.ApplicationResource, error) {
	tx := r.resourcesView

	var views []ResourceView

	sortOrder, err := createVulnerabilitiesColumnSortOrder("desc")
	if err != nil {
		return nil, fmt.Errorf("failed to create sort order: %v", err)
	}

	if err := tx.Order(sortOrder).Limit(limit).Scan(&views).Error; err != nil {
		return nil, err
	}

	ret := make([]*models.ApplicationResource, len(views))
	for i := range views {
		ret[i] = ApplicationResourceFromDB(&views[i])
	}

	return ret, nil
}

func (r *ResourceTableHandler) Count(filters *CountFilters) (int64, error) {
	var count int64
	var err error

	tx := r.resourcesView

	tx, err = r.setCountFilters(tx, filters)
	if err != nil {
		return 0, fmt.Errorf("failed to set count filters: %v", err)
	}

	if err := tx.Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count total: %v", err)
	}
	return count, nil
}

func (r *ResourceTableHandler) GetApplicationResourcesAndTotal(params GetApplicationResourcesParams) ([]ResourceView, int64, error) {
	var count int64
	var resources []ResourceView

	tx, err := r.setResourcesFilters(params)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to set filters: %v", err)
	}

	// get total item count with the set filters
	if err := tx.Count(&count).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count total: %v", err)
	}

	sortOrder, err := createApplicationResourcesSortOrder(params.SortKey, params.SortDir)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create sort order: %v", err)
	}

	// get specific page ordered items with the current filters
	if err := tx.Scopes(Paginate(params.Page, params.PageSize)).
		Order(sortOrder).
		Find(&resources).Error; err != nil {
		return nil, 0, err
	}

	return resources, count, nil
}

func createApplicationResourcesSortOrder(sortKey string, sortDir *string) (string, error) {
	if models.ApplicationResourcesSortKey(sortKey) == models.ApplicationResourcesSortKeyVulnerabilities {
		return createVulnerabilitiesColumnSortOrder(*sortDir)
	}

	sortKeyColumnName, err := getApplicationResourcesSortKeyColumnName(sortKey)
	if err != nil {
		return "", fmt.Errorf("failed to get sort key column name: %v", err)
	}

	return fmt.Sprintf("%v %v", sortKeyColumnName, strings.ToLower(*sortDir)), nil
}

func getApplicationResourcesSortKeyColumnName(key string) (string, error) {
	switch models.ApplicationResourcesSortKey(key) {
	case models.ApplicationResourcesSortKeyResourceName:
		return columnResourceName, nil
	case models.ApplicationResourcesSortKeyResourceHash:
		return columnResourceHash, nil
	case models.ApplicationResourcesSortKeyResourceType:
		return columnResourceType, nil
	case models.ApplicationResourcesSortKeyApplications:
		return columnResourceViewApplications, nil
	case models.ApplicationResourcesSortKeyPackages:
		return columnResourceViewPackages, nil
	case models.ApplicationResourcesSortKeyVulnerabilities:
		return "", fmt.Errorf("unsupported key (%v)", key)
	}

	return "", fmt.Errorf("unknown sort key (%v)", key)
}

func (r *ResourceTableHandler) getResourceIDs(params GetApplicationResourcesParams) ([]string, error) {
	ids, err := r.IDsView.GetIDs(r.createGetIDsParams(params), true)
	if err != nil {
		return nil, fmt.Errorf("failed to get IDs: %v", err)
	}

	return ids, nil
}

func (r *ResourceTableHandler) createGetIDsParams(params GetApplicationResourcesParams) GetIDsParams {
	retParams := GetIDsParams{
		LookupIDType: ResourceIDType,
	}

	// system filters - only one is allowed
	if params.CurrentRuntimeScan != nil && *params.CurrentRuntimeScan {
		retParams.FilterIDType = ApplicationIDType
		retParams.FilterIDs = params.RuntimeScanApplicationIDs
	} else if params.ApplicationID != nil {
		retParams.FilterIDType = ApplicationIDType
		retParams.FilterIDs = []string{*params.ApplicationID}
	} else if params.PackageID != nil {
		retParams.FilterIDType = PackageIDType
		retParams.FilterIDs = []string{*params.PackageID}
	}

	return retParams
}

func (r *ResourceTableHandler) getLicenses(id string) ([]string, error) {
	var licenses []string

	if err := r.licensesView.
		Select("distinct "+columnLicensesViewLicense).
		Where(columnLicensesViewResourceID+" = ?", id).
		Find(&licenses).Error; err != nil {
		return nil, fmt.Errorf("failed to get licenses: %v", err)
	}

	return licenses, nil
}

func UpdateResourceAnalyzers(resources []Resource, resourcePkgIDToAnalyzers map[ResourcePkgID][]string) []Resource {
	for i, resource := range resources {
		var resourceAnalyzers []string

		for _, pkg := range resource.Packages {
			// Retrieve analyzers from given relationships
			analyzers := resourcePkgIDToAnalyzers[CreateResourcePkgID(resource.ID, pkg.ID)]
			if len(analyzers) > 0 {
				resourceAnalyzers = append(resourceAnalyzers, analyzers...)
			}
		}

		resources[i].WithAnalyzers(slice.RemoveStringDuplicates(resourceAnalyzers))
	}

	return resources
}

func (r *ResourceTableHandler) setCountFilters(tx *gorm.DB, filters *CountFilters) (*gorm.DB, error) {
	if filters == nil {
		return tx, nil
	}

	// set application ids filter
	resourceIds, err := r.IDsView.GetIDs(GetIDsParams{
		FilterIDs:    filters.ApplicationIDs,
		FilterIDType: ApplicationIDType,
		LookupIDType: ResourceIDType,
	}, true)
	if err != nil {
		return tx, fmt.Errorf("failed to get resource ids by app ids %v: %v", filters.ApplicationIDs, err)
	}
	tx = FilterIs(tx, columnResourceID, resourceIds)

	tx = SeverityFilterGte(tx, columnSeverityCountersHighestSeverity, filters.VulnerabilitySeverityGte)

	return tx, nil
}
