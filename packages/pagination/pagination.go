// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/run-llama/llama-parse-go/internal/apijson"
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type PaginatedJobsHistory[T any] struct {
	Jobs       []T   `json:"jobs"`
	TotalCount int64 `json:"total_count"`
	Offset     int64 `json:"offset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Jobs        respjson.Field
		TotalCount  respjson.Field
		Offset      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PaginatedJobsHistory[T]) RawJSON() string { return r.JSON.raw }
func (r *PaginatedJobsHistory[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PaginatedJobsHistory[T]) GetNextPage() (res *PaginatedJobsHistory[T], err error) {
	if len(r.Jobs) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	offset := r.Offset
	length := int64(len(r.Jobs))
	next := offset + length

	if next < r.TotalCount && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PaginatedJobsHistory[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PaginatedJobsHistory[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PaginatedJobsHistoryAutoPager[T any] struct {
	page *PaginatedJobsHistory[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPaginatedJobsHistoryAutoPager[T any](page *PaginatedJobsHistory[T], err error) *PaginatedJobsHistoryAutoPager[T] {
	return &PaginatedJobsHistoryAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PaginatedJobsHistoryAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Jobs) == 0 {
		return false
	}
	if r.idx >= len(r.page.Jobs) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Jobs) == 0 {
			return false
		}
	}
	r.cur = r.page.Jobs[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PaginatedJobsHistoryAutoPager[T]) Current() T {
	return r.cur
}

func (r *PaginatedJobsHistoryAutoPager[T]) Err() error {
	return r.err
}

func (r *PaginatedJobsHistoryAutoPager[T]) Index() int {
	return r.run
}

type PaginatedPipelineFiles[T any] struct {
	Files      []T   `json:"files"`
	TotalCount int64 `json:"total_count"`
	Offset     int64 `json:"offset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Files       respjson.Field
		TotalCount  respjson.Field
		Offset      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PaginatedPipelineFiles[T]) RawJSON() string { return r.JSON.raw }
func (r *PaginatedPipelineFiles[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PaginatedPipelineFiles[T]) GetNextPage() (res *PaginatedPipelineFiles[T], err error) {
	if len(r.Files) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	offset := r.Offset
	length := int64(len(r.Files))
	next := offset + length

	if next < r.TotalCount && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PaginatedPipelineFiles[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PaginatedPipelineFiles[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PaginatedPipelineFilesAutoPager[T any] struct {
	page *PaginatedPipelineFiles[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPaginatedPipelineFilesAutoPager[T any](page *PaginatedPipelineFiles[T], err error) *PaginatedPipelineFilesAutoPager[T] {
	return &PaginatedPipelineFilesAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PaginatedPipelineFilesAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Files) == 0 {
		return false
	}
	if r.idx >= len(r.page.Files) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Files) == 0 {
			return false
		}
	}
	r.cur = r.page.Files[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PaginatedPipelineFilesAutoPager[T]) Current() T {
	return r.cur
}

func (r *PaginatedPipelineFilesAutoPager[T]) Err() error {
	return r.err
}

func (r *PaginatedPipelineFilesAutoPager[T]) Index() int {
	return r.run
}

type PaginatedBatchItems[T any] struct {
	Items     []T   `json:"items"`
	TotalSize int64 `json:"total_size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		TotalSize   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PaginatedBatchItems[T]) RawJSON() string { return r.JSON.raw }
func (r *PaginatedBatchItems[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PaginatedBatchItems[T]) GetNextPage() (res *PaginatedBatchItems[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	q := cfg.Request.URL.Query()
	offset, err := strconv.ParseInt(q.Get("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	length := int64(len(r.Items))
	next := offset + length

	if next < r.TotalSize && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PaginatedBatchItems[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PaginatedBatchItems[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PaginatedBatchItemsAutoPager[T any] struct {
	page *PaginatedBatchItems[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPaginatedBatchItemsAutoPager[T any](page *PaginatedBatchItems[T], err error) *PaginatedBatchItemsAutoPager[T] {
	return &PaginatedBatchItemsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PaginatedBatchItemsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PaginatedBatchItemsAutoPager[T]) Current() T {
	return r.cur
}

func (r *PaginatedBatchItemsAutoPager[T]) Err() error {
	return r.err
}

func (r *PaginatedBatchItemsAutoPager[T]) Index() int {
	return r.run
}

type PaginatedCloudDocuments[T any] struct {
	Documents  []T   `json:"documents"`
	TotalCount int64 `json:"total_count"`
	Offset     int64 `json:"offset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Documents   respjson.Field
		TotalCount  respjson.Field
		Offset      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PaginatedCloudDocuments[T]) RawJSON() string { return r.JSON.raw }
func (r *PaginatedCloudDocuments[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PaginatedCloudDocuments[T]) GetNextPage() (res *PaginatedCloudDocuments[T], err error) {
	if len(r.Documents) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	offset := r.Offset
	length := int64(len(r.Documents))
	next := offset + length

	if next < r.TotalCount && next != 0 {
		err = cfg.Apply(option.WithQuery("skip", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PaginatedCloudDocuments[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PaginatedCloudDocuments[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PaginatedCloudDocumentsAutoPager[T any] struct {
	page *PaginatedCloudDocuments[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPaginatedCloudDocumentsAutoPager[T any](page *PaginatedCloudDocuments[T], err error) *PaginatedCloudDocumentsAutoPager[T] {
	return &PaginatedCloudDocumentsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PaginatedCloudDocumentsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Documents) == 0 {
		return false
	}
	if r.idx >= len(r.page.Documents) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Documents) == 0 {
			return false
		}
	}
	r.cur = r.page.Documents[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PaginatedCloudDocumentsAutoPager[T]) Current() T {
	return r.cur
}

func (r *PaginatedCloudDocumentsAutoPager[T]) Err() error {
	return r.err
}

func (r *PaginatedCloudDocumentsAutoPager[T]) Index() int {
	return r.run
}

type PaginatedQuotaConfigurations[T any] struct {
	Items []T   `json:"items"`
	Page  int64 `json:"page"`
	Pages int64 `json:"pages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Page        respjson.Field
		Pages       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PaginatedQuotaConfigurations[T]) RawJSON() string { return r.JSON.raw }
func (r *PaginatedQuotaConfigurations[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PaginatedQuotaConfigurations[T]) GetNextPage() (res *PaginatedQuotaConfigurations[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	currentPage := r.Page
	if r.Pages > 0 && currentPage >= r.Pages {
		return nil, nil
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PaginatedQuotaConfigurations[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PaginatedQuotaConfigurations[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PaginatedQuotaConfigurationsAutoPager[T any] struct {
	page *PaginatedQuotaConfigurations[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPaginatedQuotaConfigurationsAutoPager[T any](page *PaginatedQuotaConfigurations[T], err error) *PaginatedQuotaConfigurationsAutoPager[T] {
	return &PaginatedQuotaConfigurationsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PaginatedQuotaConfigurationsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PaginatedQuotaConfigurationsAutoPager[T]) Current() T {
	return r.cur
}

func (r *PaginatedQuotaConfigurationsAutoPager[T]) Err() error {
	return r.err
}

func (r *PaginatedQuotaConfigurationsAutoPager[T]) Index() int {
	return r.run
}

type PaginatedCursor[T any] struct {
	Items         []T    `json:"items"`
	NextPageToken string `json:"next_page_token"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items         respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PaginatedCursor[T]) RawJSON() string { return r.JSON.raw }
func (r *PaginatedCursor[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PaginatedCursor[T]) GetNextPage() (res *PaginatedCursor[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	next := r.NextPageToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("page_token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PaginatedCursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PaginatedCursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PaginatedCursorAutoPager[T any] struct {
	page *PaginatedCursor[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPaginatedCursorAutoPager[T any](page *PaginatedCursor[T], err error) *PaginatedCursorAutoPager[T] {
	return &PaginatedCursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PaginatedCursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PaginatedCursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *PaginatedCursorAutoPager[T]) Err() error {
	return r.err
}

func (r *PaginatedCursorAutoPager[T]) Index() int {
	return r.run
}

type PaginatedCursorPost[T any] struct {
	Items         []T    `json:"items"`
	NextPageToken string `json:"next_page_token"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items         respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PaginatedCursorPost[T]) RawJSON() string { return r.JSON.raw }
func (r *PaginatedCursorPost[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PaginatedCursorPost[T]) GetNextPage() (res *PaginatedCursorPost[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	next := r.NextPageToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("page_token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PaginatedCursorPost[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PaginatedCursorPost[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PaginatedCursorPostAutoPager[T any] struct {
	page *PaginatedCursorPost[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPaginatedCursorPostAutoPager[T any](page *PaginatedCursorPost[T], err error) *PaginatedCursorPostAutoPager[T] {
	return &PaginatedCursorPostAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PaginatedCursorPostAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PaginatedCursorPostAutoPager[T]) Current() T {
	return r.cur
}

func (r *PaginatedCursorPostAutoPager[T]) Err() error {
	return r.err
}

func (r *PaginatedCursorPostAutoPager[T]) Index() int {
	return r.run
}
