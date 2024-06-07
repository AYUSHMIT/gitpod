// Copyright (c) 2024 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: gitpod/v1/workspace.proto
//
package io.gitpod.publicapi.v1

import com.connectrpc.Headers
import com.connectrpc.ResponseMessage
import com.connectrpc.ServerOnlyStreamInterface

public interface WorkspaceServiceClientInterface {
  /**
   *  GetWorkspace returns a single workspace.
   *
   *  +return NOT_FOUND User does not have access to a workspace with the given
   *  ID +return NOT_FOUND Workspace does not exist
   */
  public suspend fun getWorkspace(request: WorkspaceOuterClass.GetWorkspaceRequest, headers: Headers
      = emptyMap()): ResponseMessage<WorkspaceOuterClass.GetWorkspaceResponse>

  /**
   *  WatchWorkspaceStatus watches the workspaces status changes
   *
   *  workspace_id +return NOT_FOUND Workspace does not exist
   */
  public suspend fun watchWorkspaceStatus(headers: Headers = emptyMap()):
      ServerOnlyStreamInterface<WorkspaceOuterClass.WatchWorkspaceStatusRequest, WorkspaceOuterClass.WatchWorkspaceStatusResponse>

  /**
   *  ListWorkspaces returns a list of workspaces that match the query.
   */
  public suspend fun listWorkspaces(request: WorkspaceOuterClass.ListWorkspacesRequest,
      headers: Headers = emptyMap()): ResponseMessage<WorkspaceOuterClass.ListWorkspacesResponse>

  /**
   *  ListWorkspaceSessions returns a list of workspace sessions that match the
   */
  public suspend
      fun listWorkspaceSessions(request: WorkspaceOuterClass.ListWorkspaceSessionsRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<WorkspaceOuterClass.ListWorkspaceSessionsResponse>

  /**
   *  CreateAndStartWorkspace creates a new workspace and starts it.
   */
  public suspend
      fun createAndStartWorkspace(request: WorkspaceOuterClass.CreateAndStartWorkspaceRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<WorkspaceOuterClass.CreateAndStartWorkspaceResponse>

  /**
   *  StartWorkspace starts an existing workspace.
   *  If the specified workspace is not in stopped phase, this will return the
   *  workspace as is.
   */
  public suspend fun startWorkspace(request: WorkspaceOuterClass.StartWorkspaceRequest,
      headers: Headers = emptyMap()): ResponseMessage<WorkspaceOuterClass.StartWorkspaceResponse>

  /**
   *  UpdateWorkspace updates the workspace.
   */
  public suspend fun updateWorkspace(request: WorkspaceOuterClass.UpdateWorkspaceRequest,
      headers: Headers = emptyMap()): ResponseMessage<WorkspaceOuterClass.UpdateWorkspaceResponse>

  /**
   *  StopWorkspace stops a running workspace.
   */
  public suspend fun stopWorkspace(request: WorkspaceOuterClass.StopWorkspaceRequest,
      headers: Headers = emptyMap()): ResponseMessage<WorkspaceOuterClass.StopWorkspaceResponse>

  /**
   *  DeleteWorkspace deletes a workspace.
   *  When the workspace is running, it will be stopped as well.
   *  Deleted workspaces cannot be started again.
   */
  public suspend fun deleteWorkspace(request: WorkspaceOuterClass.DeleteWorkspaceRequest,
      headers: Headers = emptyMap()): ResponseMessage<WorkspaceOuterClass.DeleteWorkspaceResponse>

  /**
   *  ListWorkspaceClasses enumerates all available workspace classes.
   */
  public suspend fun listWorkspaceClasses(request: WorkspaceOuterClass.ListWorkspaceClassesRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<WorkspaceOuterClass.ListWorkspaceClassesResponse>

  /**
   *  ParseContextURL parses a context URL and returns the workspace metadata and
   *  spec. Not implemented yet.
   */
  public suspend fun parseContextURL(request: WorkspaceOuterClass.ParseContextURLRequest,
      headers: Headers = emptyMap()): ResponseMessage<WorkspaceOuterClass.ParseContextURLResponse>

  /**
   *  GetWorkspaceDefaultImage returns the default workspace image of specified
   *  workspace.
   */
  public suspend
      fun getWorkspaceDefaultImage(request: WorkspaceOuterClass.GetWorkspaceDefaultImageRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<WorkspaceOuterClass.GetWorkspaceDefaultImageResponse>

  /**
   *  SendHeartBeat sends a heartbeat to activate the workspace
   */
  public suspend fun sendHeartBeat(request: WorkspaceOuterClass.SendHeartBeatRequest,
      headers: Headers = emptyMap()): ResponseMessage<WorkspaceOuterClass.SendHeartBeatResponse>

  /**
   *  GetWorkspaceOwnerToken returns an owner token of workspace.
   */
  public suspend
      fun getWorkspaceOwnerToken(request: WorkspaceOuterClass.GetWorkspaceOwnerTokenRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<WorkspaceOuterClass.GetWorkspaceOwnerTokenResponse>

  /**
   *  GetWorkspaceEditorCredentials returns an credentials that is used in editor
   *  to encrypt and decrypt secrets
   */
  public suspend
      fun getWorkspaceEditorCredentials(request: WorkspaceOuterClass.GetWorkspaceEditorCredentialsRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<WorkspaceOuterClass.GetWorkspaceEditorCredentialsResponse>

  /**
   *  CreateWorkspaceSnapshot creates a snapshot of the workspace that can be
   *  shared with others.
   */
  public suspend
      fun createWorkspaceSnapshot(request: WorkspaceOuterClass.CreateWorkspaceSnapshotRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<WorkspaceOuterClass.CreateWorkspaceSnapshotResponse>

  /**
   *  WaitWorkspaceSnapshot waits for the snapshot to be available or failed.
   */
  public suspend
      fun waitForWorkspaceSnapshot(request: WorkspaceOuterClass.WaitForWorkspaceSnapshotRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<WorkspaceOuterClass.WaitForWorkspaceSnapshotResponse>

  /**
   *  UpdateWorkspacePort updates the port of workspace.
   */
  public suspend fun updateWorkspacePort(request: WorkspaceOuterClass.UpdateWorkspacePortRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<WorkspaceOuterClass.UpdateWorkspacePortResponse>
}
