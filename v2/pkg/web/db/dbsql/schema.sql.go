// Code generated by sqlc. DO NOT EDIT.
// source: schema.sql

package dbsql

import (
	"context"
	"database/sql"
)

const addIssue = `-- name: AddIssue :exec
INSERT INTO "public".issues
	(matchedat, title, severity, createdat, updatedat, scansource, issuestate, description, author, cvss, cwe, labels, issuedata, issuetemplate, templatename, remediation, debug, scanid) 
VALUES 
    ($1, $2, $3, NOW(), NOW(), $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
`

type AddIssueParams struct {
	Matchedat     sql.NullString
	Title         sql.NullString
	Severity      sql.NullString
	Scansource    sql.NullString
	Issuestate    sql.NullString
	Description   sql.NullString
	Author        sql.NullString
	Cvss          sql.NullFloat64
	Cwe           []int32
	Labels        []string
	Issuedata     sql.NullString
	Issuetemplate sql.NullString
	Templatename  sql.NullString
	Remediation   sql.NullString
	Debug         sql.NullString
	Scanid        sql.NullInt64
}

func (q *Queries) AddIssue(ctx context.Context, arg AddIssueParams) error {
	_, err := q.db.Exec(ctx, addIssue,
		arg.Matchedat,
		arg.Title,
		arg.Severity,
		arg.Scansource,
		arg.Issuestate,
		arg.Description,
		arg.Author,
		arg.Cvss,
		arg.Cwe,
		arg.Labels,
		arg.Issuedata,
		arg.Issuetemplate,
		arg.Templatename,
		arg.Remediation,
		arg.Debug,
		arg.Scanid,
	)
	return err
}

const addScan = `-- name: AddScan :one
INSERT INTO "public".scans
	( name, status, scantime, hosts, scansource, templates, targets, config, runnow, reporting, scheduleoccurence, scheduletime) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12 ) RETURNING id
`

type AddScanParams struct {
	Name              sql.NullString
	Status            sql.NullString
	Scantime          sql.NullInt64
	Hosts             sql.NullInt64
	Scansource        sql.NullString
	Templates         []string
	Targets           []string
	Config            sql.NullString
	Runnow            sql.NullBool
	Reporting         sql.NullString
	Scheduleoccurence sql.NullString
	Scheduletime      sql.NullString
}

func (q *Queries) AddScan(ctx context.Context, arg AddScanParams) (int64, error) {
	row := q.db.QueryRow(ctx, addScan,
		arg.Name,
		arg.Status,
		arg.Scantime,
		arg.Hosts,
		arg.Scansource,
		arg.Templates,
		arg.Targets,
		arg.Config,
		arg.Runnow,
		arg.Reporting,
		arg.Scheduleoccurence,
		arg.Scheduletime,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const addTarget = `-- name: AddTarget :exec
INSERT INTO public.targets
	( name, createdat, updatedat, internalid, filename, total) VALUES ($1, NOW(), NOW(), $2, $3, $4)
`

type AddTargetParams struct {
	Name       sql.NullString
	Internalid sql.NullString
	Filename   sql.NullString
	Total      sql.NullInt64
}

func (q *Queries) AddTarget(ctx context.Context, arg AddTargetParams) error {
	_, err := q.db.Exec(ctx, addTarget,
		arg.Name,
		arg.Internalid,
		arg.Filename,
		arg.Total,
	)
	return err
}

const addTemplate = `-- name: AddTemplate :exec
INSERT INTO public.templates
( name, folder, "path", contents, createdat, updatedat, hash) VALUES ($1, $2, $3 , $4, NOW(), NOW(), $5)
`

type AddTemplateParams struct {
	Name     sql.NullString
	Folder   sql.NullString
	Path     string
	Contents string
	Hash     sql.NullString
}

func (q *Queries) AddTemplate(ctx context.Context, arg AddTemplateParams) error {
	_, err := q.db.Exec(ctx, addTemplate,
		arg.Name,
		arg.Folder,
		arg.Path,
		arg.Contents,
		arg.Hash,
	)
	return err
}

const deleteIssue = `-- name: DeleteIssue :exec
DELETE FROM "public".issues WHERE id=$1
`

func (q *Queries) DeleteIssue(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteIssue, id)
	return err
}

const deleteIssueByScanID = `-- name: DeleteIssueByScanID :exec
DELETE FROM "public".issues WHERE scanid=$1
`

func (q *Queries) DeleteIssueByScanID(ctx context.Context, scanid sql.NullInt64) error {
	_, err := q.db.Exec(ctx, deleteIssueByScanID, scanid)
	return err
}

const deleteScan = `-- name: DeleteScan :exec
DELETE FROM "public".scans WHERE id=$1
`

func (q *Queries) DeleteScan(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteScan, id)
	return err
}

const deleteTarget = `-- name: DeleteTarget :exec
DELETE FROM public.targets WHERE ID=$1
`

func (q *Queries) DeleteTarget(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteTarget, id)
	return err
}

const deleteTemplate = `-- name: DeleteTemplate :exec
DELETE FROM public.templates WHERE path=$1
`

func (q *Queries) DeleteTemplate(ctx context.Context, path string) error {
	_, err := q.db.Exec(ctx, deleteTemplate, path)
	return err
}

const getIssue = `-- name: GetIssue :one
SELECT matchedat, title, severity, createdat, updatedat, scansource, issuestate, description, author, cvss, cwe, labels, 
	issuedata, issuetemplate, templatename, remediation, debug, id, scanid
FROM
	"public".issues WHERE id=$1 LIMIT 1
`

func (q *Queries) GetIssue(ctx context.Context, id int64) (Issue, error) {
	row := q.db.QueryRow(ctx, getIssue, id)
	var i Issue
	err := row.Scan(
		&i.Matchedat,
		&i.Title,
		&i.Severity,
		&i.Createdat,
		&i.Updatedat,
		&i.Scansource,
		&i.Issuestate,
		&i.Description,
		&i.Author,
		&i.Cvss,
		&i.Cwe,
		&i.Labels,
		&i.Issuedata,
		&i.Issuetemplate,
		&i.Templatename,
		&i.Remediation,
		&i.Debug,
		&i.ID,
		&i.Scanid,
	)
	return i, err
}

const getIssues = `-- name: GetIssues :many
SELECT id, scanid, matchedat, title, severity, createdat, updatedat, scansource
FROM
	"public".issues
`

type GetIssuesRow struct {
	ID         int64
	Scanid     sql.NullInt64
	Matchedat  sql.NullString
	Title      sql.NullString
	Severity   sql.NullString
	Createdat  sql.NullTime
	Updatedat  sql.NullTime
	Scansource sql.NullString
}

func (q *Queries) GetIssues(ctx context.Context) ([]GetIssuesRow, error) {
	rows, err := q.db.Query(ctx, getIssues)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetIssuesRow
	for rows.Next() {
		var i GetIssuesRow
		if err := rows.Scan(
			&i.ID,
			&i.Scanid,
			&i.Matchedat,
			&i.Title,
			&i.Severity,
			&i.Createdat,
			&i.Updatedat,
			&i.Scansource,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getIssuesMatches = `-- name: GetIssuesMatches :many
SELECT id, matchedat, templatename, severity, author
FROM
	"public".issues WHERE scanid=$1
`

type GetIssuesMatchesRow struct {
	ID           int64
	Matchedat    sql.NullString
	Templatename sql.NullString
	Severity     sql.NullString
	Author       sql.NullString
}

func (q *Queries) GetIssuesMatches(ctx context.Context, scanid sql.NullInt64) ([]GetIssuesMatchesRow, error) {
	rows, err := q.db.Query(ctx, getIssuesMatches, scanid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetIssuesMatchesRow
	for rows.Next() {
		var i GetIssuesMatchesRow
		if err := rows.Scan(
			&i.ID,
			&i.Matchedat,
			&i.Templatename,
			&i.Severity,
			&i.Author,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getScan = `-- name: GetScan :one
SELECT name, status, scantime, hosts, scansource, templates, targets, config, runnow, reporting, scheduleoccurence, 
	scheduletime, id
FROM
	"public".scans WHERE id=$1 LIMIT 1
`

func (q *Queries) GetScan(ctx context.Context, id int64) (Scan, error) {
	row := q.db.QueryRow(ctx, getScan, id)
	var i Scan
	err := row.Scan(
		&i.Name,
		&i.Status,
		&i.Scantime,
		&i.Hosts,
		&i.Scansource,
		&i.Templates,
		&i.Targets,
		&i.Config,
		&i.Runnow,
		&i.Reporting,
		&i.Scheduleoccurence,
		&i.Scheduletime,
		&i.ID,
	)
	return i, err
}

const getScans = `-- name: GetScans :many
SELECT name, status, scantime, hosts, scansource, templates, targets, config, runnow, reporting, scheduleoccurence, 
	scheduletime, id
FROM
	"public".scans
`

func (q *Queries) GetScans(ctx context.Context) ([]Scan, error) {
	rows, err := q.db.Query(ctx, getScans)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Scan
	for rows.Next() {
		var i Scan
		if err := rows.Scan(
			&i.Name,
			&i.Status,
			&i.Scantime,
			&i.Hosts,
			&i.Scansource,
			&i.Templates,
			&i.Targets,
			&i.Config,
			&i.Runnow,
			&i.Reporting,
			&i.Scheduleoccurence,
			&i.Scheduletime,
			&i.ID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getScansBySearchKey = `-- name: GetScansBySearchKey :many
SELECT name, status, scantime, hosts, scansource, templates, targets, config, runnow, reporting, scheduleoccurence, 
	scheduletime, id
FROM
	"public".scans WHERE name LIKE '%'||$1||'%'
`

func (q *Queries) GetScansBySearchKey(ctx context.Context, dollar_1 sql.NullString) ([]Scan, error) {
	rows, err := q.db.Query(ctx, getScansBySearchKey, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Scan
	for rows.Next() {
		var i Scan
		if err := rows.Scan(
			&i.Name,
			&i.Status,
			&i.Scantime,
			&i.Hosts,
			&i.Scansource,
			&i.Templates,
			&i.Targets,
			&i.Config,
			&i.Runnow,
			&i.Reporting,
			&i.Scheduleoccurence,
			&i.Scheduletime,
			&i.ID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSettingByName = `-- name: GetSettingByName :one
SELECT settingdata, datatype
FROM
	"public".settings WHERE name=$1 LIMIT 1
`

type GetSettingByNameRow struct {
	Settingdata sql.NullString
	Datatype    sql.NullString
}

func (q *Queries) GetSettingByName(ctx context.Context, name sql.NullString) (GetSettingByNameRow, error) {
	row := q.db.QueryRow(ctx, getSettingByName, name)
	var i GetSettingByNameRow
	err := row.Scan(&i.Settingdata, &i.Datatype)
	return i, err
}

const getSettings = `-- name: GetSettings :many
SELECT settingdata, datatype, name
FROM
	"public".settings
`

func (q *Queries) GetSettings(ctx context.Context) ([]Setting, error) {
	rows, err := q.db.Query(ctx, getSettings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Setting
	for rows.Next() {
		var i Setting
		if err := rows.Scan(&i.Settingdata, &i.Datatype, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTarget = `-- name: GetTarget :one
SELECT name, internalid, filename, total, createdat, updatedat
FROM
	public.targets WHERE ID=$1 LIMIT 1
`

type GetTargetRow struct {
	Name       sql.NullString
	Internalid sql.NullString
	Filename   sql.NullString
	Total      sql.NullInt64
	Createdat  sql.NullTime
	Updatedat  sql.NullTime
}

func (q *Queries) GetTarget(ctx context.Context, id int64) (GetTargetRow, error) {
	row := q.db.QueryRow(ctx, getTarget, id)
	var i GetTargetRow
	err := row.Scan(
		&i.Name,
		&i.Internalid,
		&i.Filename,
		&i.Total,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const getTargetByName = `-- name: GetTargetByName :one
SELECT id, internalid, filename, total, createdat, updatedat
FROM
	public.targets WHERE name=$1 LIMIT 1
`

type GetTargetByNameRow struct {
	ID         int64
	Internalid sql.NullString
	Filename   sql.NullString
	Total      sql.NullInt64
	Createdat  sql.NullTime
	Updatedat  sql.NullTime
}

func (q *Queries) GetTargetByName(ctx context.Context, name sql.NullString) (GetTargetByNameRow, error) {
	row := q.db.QueryRow(ctx, getTargetByName, name)
	var i GetTargetByNameRow
	err := row.Scan(
		&i.ID,
		&i.Internalid,
		&i.Filename,
		&i.Total,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const getTargets = `-- name: GetTargets :many
SELECT id, name, createdat, updatedat, internalid, filename, total
FROM
	public.targets
`

type GetTargetsRow struct {
	ID         int64
	Name       sql.NullString
	Createdat  sql.NullTime
	Updatedat  sql.NullTime
	Internalid sql.NullString
	Filename   sql.NullString
	Total      sql.NullInt64
}

func (q *Queries) GetTargets(ctx context.Context) ([]GetTargetsRow, error) {
	rows, err := q.db.Query(ctx, getTargets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTargetsRow
	for rows.Next() {
		var i GetTargetsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Createdat,
			&i.Updatedat,
			&i.Internalid,
			&i.Filename,
			&i.Total,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTargetsForSearch = `-- name: GetTargetsForSearch :many
SELECT id, name, createdat, updatedat, internalid, filename, total
FROM
	"public".targets WHERE name LIKE '%'||$1||'%' OR filename LIKE '%'||$1||'%'
`

type GetTargetsForSearchRow struct {
	ID         int64
	Name       sql.NullString
	Createdat  sql.NullTime
	Updatedat  sql.NullTime
	Internalid sql.NullString
	Filename   sql.NullString
	Total      sql.NullInt64
}

func (q *Queries) GetTargetsForSearch(ctx context.Context, dollar_1 sql.NullString) ([]GetTargetsForSearchRow, error) {
	rows, err := q.db.Query(ctx, getTargetsForSearch, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTargetsForSearchRow
	for rows.Next() {
		var i GetTargetsForSearchRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Createdat,
			&i.Updatedat,
			&i.Internalid,
			&i.Filename,
			&i.Total,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTemplateContents = `-- name: GetTemplateContents :one
SELECT contents FROM public.templates WHERE path=$1 LIMIT 1
`

func (q *Queries) GetTemplateContents(ctx context.Context, path string) (string, error) {
	row := q.db.QueryRow(ctx, getTemplateContents, path)
	var contents string
	err := row.Scan(&contents)
	return contents, err
}

const getTemplates = `-- name: GetTemplates :many
SELECT id, name, folder, "path", createdat, updatedat, hash
FROM
	"public".templates
`

type GetTemplatesRow struct {
	ID        int64
	Name      sql.NullString
	Folder    sql.NullString
	Path      string
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Hash      sql.NullString
}

func (q *Queries) GetTemplates(ctx context.Context) ([]GetTemplatesRow, error) {
	rows, err := q.db.Query(ctx, getTemplates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTemplatesRow
	for rows.Next() {
		var i GetTemplatesRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Folder,
			&i.Path,
			&i.Createdat,
			&i.Updatedat,
			&i.Hash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTemplatesByFolder = `-- name: GetTemplatesByFolder :many
SELECT id, name, "path", createdat, updatedat, hash
FROM
	"public".templates WHERE folder=$1
`

type GetTemplatesByFolderRow struct {
	ID        int64
	Name      sql.NullString
	Path      string
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Hash      sql.NullString
}

func (q *Queries) GetTemplatesByFolder(ctx context.Context, folder sql.NullString) ([]GetTemplatesByFolderRow, error) {
	rows, err := q.db.Query(ctx, getTemplatesByFolder, folder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTemplatesByFolderRow
	for rows.Next() {
		var i GetTemplatesByFolderRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Path,
			&i.Createdat,
			&i.Updatedat,
			&i.Hash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTemplatesByFolderOne = `-- name: GetTemplatesByFolderOne :one
SELECT id, name, "path", createdat, updatedat, hash
FROM
	"public".templates WHERE folder=$1 LIMIT 1
`

type GetTemplatesByFolderOneRow struct {
	ID        int64
	Name      sql.NullString
	Path      string
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Hash      sql.NullString
}

func (q *Queries) GetTemplatesByFolderOne(ctx context.Context, folder sql.NullString) (GetTemplatesByFolderOneRow, error) {
	row := q.db.QueryRow(ctx, getTemplatesByFolderOne, folder)
	var i GetTemplatesByFolderOneRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Path,
		&i.Createdat,
		&i.Updatedat,
		&i.Hash,
	)
	return i, err
}

const getTemplatesBySearchKey = `-- name: GetTemplatesBySearchKey :many
SELECT id, name, folder, "path", createdat, updatedat, hash
FROM
	"public".templates WHERE path LIKE '%'||$1||'%'
`

type GetTemplatesBySearchKeyRow struct {
	ID        int64
	Name      sql.NullString
	Folder    sql.NullString
	Path      string
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Hash      sql.NullString
}

func (q *Queries) GetTemplatesBySearchKey(ctx context.Context, dollar_1 sql.NullString) ([]GetTemplatesBySearchKeyRow, error) {
	rows, err := q.db.Query(ctx, getTemplatesBySearchKey, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTemplatesBySearchKeyRow
	for rows.Next() {
		var i GetTemplatesBySearchKeyRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Folder,
			&i.Path,
			&i.Createdat,
			&i.Updatedat,
			&i.Hash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTemplatesForScan = `-- name: GetTemplatesForScan :many
SELECT path, contents FROM public.templates WHERE folder=$1 OR path=$1 OR path LIKE $1||'%'
`

type GetTemplatesForScanRow struct {
	Path     string
	Contents string
}

func (q *Queries) GetTemplatesForScan(ctx context.Context, folder sql.NullString) ([]GetTemplatesForScanRow, error) {
	rows, err := q.db.Query(ctx, getTemplatesForScan, folder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTemplatesForScanRow
	for rows.Next() {
		var i GetTemplatesForScanRow
		if err := rows.Scan(&i.Path, &i.Contents); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setSettings = `-- name: SetSettings :exec
INSERT INTO "public".settings
	( settingdata, datatype, name) VALUES ( $1, $2, $3) ON CONFLICT (name) DO UPDATE SET settingdata=$1
`

type SetSettingsParams struct {
	Settingdata sql.NullString
	Datatype    sql.NullString
	Name        sql.NullString
}

func (q *Queries) SetSettings(ctx context.Context, arg SetSettingsParams) error {
	_, err := q.db.Exec(ctx, setSettings, arg.Settingdata, arg.Datatype, arg.Name)
	return err
}

const updateIssue = `-- name: UpdateIssue :exec
UPDATE "public".issues SET issuestate=$2 WHERE id=$1
`

type UpdateIssueParams struct {
	ID         int64
	Issuestate sql.NullString
}

func (q *Queries) UpdateIssue(ctx context.Context, arg UpdateIssueParams) error {
	_, err := q.db.Exec(ctx, updateIssue, arg.ID, arg.Issuestate)
	return err
}

const updateSettings = `-- name: UpdateSettings :exec
UPDATE "public".settings SET settingdata=$1 WHERE name=$2
`

type UpdateSettingsParams struct {
	Settingdata sql.NullString
	Name        sql.NullString
}

func (q *Queries) UpdateSettings(ctx context.Context, arg UpdateSettingsParams) error {
	_, err := q.db.Exec(ctx, updateSettings, arg.Settingdata, arg.Name)
	return err
}

const updateTargetMetadata = `-- name: UpdateTargetMetadata :exec
UPDATE targets SET total=total+$1 AND updatedAt=NOW() WHERE id=$2
`

type UpdateTargetMetadataParams struct {
	Total sql.NullInt64
	ID    int64
}

func (q *Queries) UpdateTargetMetadata(ctx context.Context, arg UpdateTargetMetadataParams) error {
	_, err := q.db.Exec(ctx, updateTargetMetadata, arg.Total, arg.ID)
	return err
}

const updateTemplate = `-- name: UpdateTemplate :exec
UPDATE public.templates SET contents=$1, updatedat=$2, hash=$4 WHERE path=$3
`

type UpdateTemplateParams struct {
	Contents  string
	Updatedat sql.NullTime
	Path      string
	Hash      sql.NullString
}

func (q *Queries) UpdateTemplate(ctx context.Context, arg UpdateTemplateParams) error {
	_, err := q.db.Exec(ctx, updateTemplate,
		arg.Contents,
		arg.Updatedat,
		arg.Path,
		arg.Hash,
	)
	return err
}