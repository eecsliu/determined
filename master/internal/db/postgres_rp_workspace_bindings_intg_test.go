//go:build integration
// +build integration

package db

import (
	"context"
	log "github.com/sirupsen/logrus"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/determined-ai/determined/master/internal/config"
	"github.com/determined-ai/determined/master/pkg/etc"
)

func TestAddAndRemoveBindings(t *testing.T) {
	// Test single insert/delete
	// Test bulk insert/delete
	return
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

func TestFindUnboundPools(t *testing.T) {
	ctx := context.Background()
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	user := RequireMockUser(t, db)
	var poolsList []config.ResourcePoolConfig
	for i := 0; i < 10; i++ {
		poolsList = append(poolsList, config.ResourcePoolConfig{PoolName: "TestPool" + strconv.Itoa(i)})
	}
	workspaceNames := []string{"test1", "test2"}
	workspaceIDs, err := MockWorkspaces(workspaceNames, user.ID)
	require.NoError(t, err)

	defer func() {
		err = CleanupMockWorkspace(workspaceIDs)
		if err != nil {
			log.Error("error when cleaning up mock workspaces")
		}
	}()

	for i := 0; i < 9; i++ {
		err = AddRPWorkspaceBindings(ctx, []int32{workspaceIDs[0]}, poolsList[i].PoolName, poolsList)
		require.NoError(t, err, "failure to add rp workspace bindings")
	}

	foundPools, err := FindUnboundResourcePools(ctx, poolsList)
	require.NoError(t, err, "error finding unbound resource pools")
	require.Equal(t, 1, len(foundPools))
	require.Equal(t, "TestPool9", foundPools[0],
		"expected pool name TestPool9 but got pool name %s", foundPools[0])
}

func TestListRPsBoundToWorkspace(t *testing.T) {
	// pretty straightforward
	// don't list binding that are invalid
	// return unbound pools too (we pull from config)
	ctx := context.Background()
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)

	user := RequireMockUser(t, db)
	var poolsList []config.ResourcePoolConfig
	for i := 0; i < 10; i++ {
		poolsList = append(poolsList, config.ResourcePoolConfig{PoolName: "TestPool" + strconv.Itoa(i)})
	}
	workspaceNames := []string{"test1", "test2"}
	workspaceIDs, err := MockWorkspaces(workspaceNames, user.ID)
	require.NoError(t, err)

	defer func() {
		err = CleanupMockWorkspace(workspaceIDs)
		if err != nil {
			log.Error("error when cleaning up mock workspaces")
		}
	}()

	var expectedPools []string
	for i := 0; i < 9; i++ {
		err = AddRPWorkspaceBindings(ctx, []int32{workspaceIDs[0]}, poolsList[i].PoolName, poolsList)
		require.NoError(t, err, "failure to add rp workspace bindings")
		expectedPools = append(expectedPools, poolsList[i].PoolName)
	}

	resourcePools, _, err := ReadRPsBoundToWorkspace(ctx, int(workspaceIDs[0]), 0, 10)
	require.NoError(t, err, "error reading resource pools available to workspace")

	var actualPools []string
	for _, pool := range resourcePools {
		actualPools = append(actualPools, pool.PoolName)
	}
	require.Equal(t, expectedPools, actualPools, "expected available pools to be $t, but got %t", expectedPools, actualPools)

	resourcePools, _, err = ReadRPsBoundToWorkspace(ctx, int(workspaceIDs[1]), 0, 10)
	require.NoError(t, err, "error reading resource pools available to workspace")
	require.Equal(t, 0, len(resourcePools), "expected 0 results, but received %d", len(resourcePools))

	return
}

func TestListAllBindings(t *testing.T) {
	// pretty straightforward
	// list ALL bindings, even invalid ones
	// make sure to return unbound pools too (we pull from config)
	return
}

func TestOverwriteBindings(t *testing.T) {
	require.NoError(t, etc.SetRootPath(RootFromDB))
	db := MustResolveTestPostgres(t)
	MustMigrateTestPostgres(t, db, MigrationsFromDB)
	ctx := context.Background()
	user := RequireMockUser(t, db)
	var existingPools []config.ResourcePoolConfig
	pool := config.ResourcePoolConfig{PoolName: "poolName1"}
	existingPools = append(existingPools, pool)
	// Test overwrite bindings
	poolName := "poolName1" //nolint:goconst
	workspaceNames := []string{"test1", "test2", "test3"}
	workspaceIDs, err := MockWorkspaces(workspaceNames, user.ID)
	require.NoError(t, err)
	err = AddRPWorkspaceBindings(ctx, workspaceIDs, poolName, existingPools)
	require.NoError(t, err)
	err = OverwriteRPWorkspaceBindings(ctx, workspaceIDs, poolName, existingPools)
	require.NoError(t, err)
	// TODO: call list bindings here to make sure it worked
	// Test overwrite pool that's not bound to anything currently
	pool = config.ResourcePoolConfig{PoolName: "poolName2"}
	existingPools = append(existingPools, pool)
	poolName = "poolName2"
	err = OverwriteRPWorkspaceBindings(ctx, workspaceIDs, poolName, existingPools)
	require.NoError(t, err)
	// TODO: call list bindings here to make sure it worked
	err = CleanupMockWorkspace(workspaceIDs)
	require.NoError(t, err)
	return
}

func TestOverwriteFail(t *testing.T) {
	ctx := context.Background()
	var existingPools []config.ResourcePoolConfig
	pool := config.ResourcePoolConfig{PoolName: "poolName1"}
	existingPools = append(existingPools, pool)
	// Test overwrite adding workspace that doesn't exist
	workspaceIDs := []int32{100, 102, 103}
	poolName := "poolName1"
	err := OverwriteRPWorkspaceBindings(ctx, workspaceIDs, poolName, existingPools)
	// db Error that workspace doesn't exist
	require.ErrorContains(t, err, "violates foreign key constraint")
	// Test overwrite pool that doesn't exist
	poolName = "poolNameDoesntExist"
	err = OverwriteRPWorkspaceBindings(ctx, workspaceIDs, poolName, existingPools)
	require.ErrorContains(t, err, "doesn't exist in config")
	return
}

func TestRemoveInvalidBinding(t *testing.T) {
	ctx := context.Background()
	// remove binding that doesn't exist
	poolName := "poolName" //nolint:goconst
	workspaceIDs := []int32{1}
	err := RemoveRPWorkspaceBindings(ctx, workspaceIDs, poolName)
	require.ErrorContains(t, err, "binding doesn't exist")
	// bulk remove bindings that don't exist
	poolName = "poolName"
	workspaceIDs = []int32{1, 2, 3}
	err = RemoveRPWorkspaceBindings(ctx, workspaceIDs, poolName)
	require.ErrorContains(t, err, " binding doesn't exist")
	return
}
