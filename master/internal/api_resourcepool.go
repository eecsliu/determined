package internal

import (
	"context"

	"github.com/determined-ai/determined/proto/pkg/apiv1"
)

func (a *apiServer) GetResourcePools(
	_ context.Context, req *apiv1.GetResourcePoolsRequest,
) (*apiv1.GetResourcePoolsResponse, error) {
	resp, err := a.m.rm.GetResourcePools(a.m.system, req)
	if err != nil {
		return nil, err
	}
	return resp, a.paginate(&resp.Pagination, &resp.ResourcePools, req.Offset, req.Limit)
}

func (a *apiServer) BindRPToWorkspace(
	ctx context.Context, req *apiv1.BindRPToWorkspaceRequest,
) (*apiv1.BindRPToWorkspaceResponse, error) {
	return nil, nil
}

func (a *apiServer) OverwriteRPWorkspaceBindings(
	ctx context.Context, req *apiv1.OverwriteRPWorkspaceBindingsRequest,
) (*apiv1.OverwriteRPWorkspaceBindingsResponse, error) {
	return nil, nil
}

func (a *apiServer) UnbindRPFromWorkspace(
	ctx context.Context, req *apiv1.UnbindRPFromWorkspaceRequest,
) (*apiv1.UnbindRPFromWorkspaceResponse, error) {
	return nil, nil
}

func (a *apiServer) ListWorkspacesBoundToRP(
	ctx context.Context, req *apiv1.ListWorkspacesBoundToRPRequest,
) (*apiv1.ListWorkspacesBoundToRPResponse, error) {
	return nil, nil
}

func (a *apiServer) ListRPWorkspaceBindings(
	ctx context.Context, req *apiv1.ListRPWorkspaceBindingsRequest,
) (*apiv1.ListRPWorkspaceBindingsResponse, error) {
	return nil, nil
}
