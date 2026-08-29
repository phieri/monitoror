//go:generate mockery --name MergeRequestsService

package gogitlab

import (
	"github.com/xanzy/go-gitlab"
)

type MergeRequestsService interface {
	GetMergeRequest(pid any, mergeRequest int, opt *gitlab.GetMergeRequestsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error)
	ListProjectMergeRequests(pid any, opt *gitlab.ListProjectMergeRequestsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
	ListMergeRequestPipelines(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) ([]*gitlab.PipelineInfo, *gitlab.Response, error)
}
