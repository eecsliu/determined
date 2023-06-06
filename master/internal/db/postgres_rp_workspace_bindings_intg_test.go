package db

import (
	"context"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestAddAndRemoveBindings(t *testing.T) {
	ctx := context.Background()
	pgDB := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, pgDB, MigrationsFromDB)

	user := RequireMockUser(t, pgDB)
	workspaceIDs, err := MockWorkspaces([]string{"test1", "test2", "test3", "test4"}, user.ID)
	require.NoError(t, err, "failed creating workspaces: %t", err)
	defer func() {
		err = CleanupMockWorkspace(workspaceIDs)
		if err != nil {
			log.Errorf("error when cleaning up mock workspaces")
		}
	}()

	var int32IDs []int32
	for _, workspaceID := range workspaceIDs {
		int32IDs = append(int32IDs, int32(workspaceID))
	}
	// Test single insert/delete
	// Test bulk insert/delete
	err = AddRPWorkspaceBindings(ctx, []int32{int32IDs[0]}, "test pool")
	require.NoError(t, err, "failed to add bindings: %t", err)

	var values []RPWorkspaceBinding
	count, err := Bun().NewSelect().Model(&values).ScanAndCount(ctx)
	require.NoError(t, err, "error when scanning DB: %t", err)
	require.Equal(t, 1, count, "expected 1 item in DB, found %d", count)
	require.Equal(t, workspaceIDs[0], values[0].WorkspaceID,
		"expected workspaceID to be %d, but it is %d", workspaceIDs[0], values[0].WorkspaceID)
	require.Equal(t, "test pool", values[0].PoolName,
		"expected pool name to be 'test pool', but got %s", values[0].PoolName)

	err = RemoveRPWorkspaceBindings(ctx, []int32{int32IDs[0]}, "test pool")
	require.NoError(t, err, "failed to remove bindings: %t", err)

	count, err = Bun().NewSelect().Model(&values).ScanAndCount(ctx)
	require.NoError(t, err, "error when scanning DB: %t", err)
	require.Equal(t, 0, count, "expected 0 items in DB, found %d", count)

	err = AddRPWorkspaceBindings(ctx, int32IDs, "test pool too")
	require.NoError(t, err, "failed to add bindings: %t", err)

	count, err = Bun().NewSelect().Model(&values).Order("workspace_id ASC").ScanAndCount(ctx)
	require.NoError(t, err, "error when scanning DB: %t", err)
	require.Equal(t, 4, count, "expected 3 items in DB, found %d", count)
	for i := 0; i < 4; i++ {
		require.Equal(t, workspaceIDs[i], values[i].WorkspaceID,
			"expected workspaceID to be %d, but it is %d", i+1, values[i].WorkspaceID)
		require.Equal(t, "test pool too", values[i].PoolName,
			"expected pool name to be 'test pool too', but got %s", values[i].PoolName)
	}

	err = RemoveRPWorkspaceBindings(ctx, int32IDs, "test pool too")
	require.NoError(t, err, "failed to remove bindings: %t", err)

	count, err = Bun().NewSelect().Model(&values).ScanAndCount(ctx)
	require.NoError(t, err, "error when scanning DB: %t", err)
	require.Equal(t, 0, count, "expected 0 items in DB, found %d", count)
}

func TestBindingFail(t *testing.T) {
	// Test add the same binding multiple times - should fail
	// Test add same binding among bulk transaction - should fail the entire transaction
	// Test add workspace that doesn't exist
	// Test add RP that doesn't exist
	return
}

func TestListWorkspacesBindingRP(t *testing.T) {
	// pretty straightforward
	// don't list bindings that are invalid
	// if RP is unbound, return nothing
	return
}

func TestListRPsBoundToWorkspace(t *testing.T) {
	// pretty straightforward
	// don't list binding that are invalid
	// return unbound pools too (we pull from config)
	return
}

func TestListAllBindings(t *testing.T) {
	// pretty straightforward
	// list ALL bindings, even invalid ones
	// make sure to return unbound pools too (we pull from config)
	return
}

func TestOverwriteBindings(t *testing.T) {
	// Test overwrite bindings
	// Test overwrite pool that's not bound to anything currently
	return
}

func TestOverwriteFail(t *testing.T) {
	// Test overwrite adding workspace that doesn't exist
	// Test overwrite pool that doesn't exist
	return
}

func TestRemoveInvalidBinding(t *testing.T) {
	// remove binding that doesn't exist
	// bulk remove bindings that don't exist
	return
}
