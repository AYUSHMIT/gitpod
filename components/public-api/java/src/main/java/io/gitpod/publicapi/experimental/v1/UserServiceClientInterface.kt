// Copyright (c) 2024 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: gitpod/experimental/v1/user.proto
//
package io.gitpod.publicapi.experimental.v1

import com.connectrpc.Headers
import com.connectrpc.ResponseMessage

public interface UserServiceClientInterface {
  /**
   *  GetAuthenticatedUser gets the user info.
   */
  public suspend fun getAuthenticatedUser(request: UserOuterClass.GetAuthenticatedUserRequest,
      headers: Headers = emptyMap()): ResponseMessage<UserOuterClass.GetAuthenticatedUserResponse>

  /**
   *  ListSSHKeys lists the public SSH keys.
   */
  public suspend fun listSSHKeys(request: UserOuterClass.ListSSHKeysRequest, headers: Headers =
      emptyMap()): ResponseMessage<UserOuterClass.ListSSHKeysResponse>

  /**
   *  CreateSSHKey adds a public SSH key.
   */
  public suspend fun createSSHKey(request: UserOuterClass.CreateSSHKeyRequest, headers: Headers =
      emptyMap()): ResponseMessage<UserOuterClass.CreateSSHKeyResponse>

  /**
   *  GetSSHKey retrieves an ssh key by ID.
   */
  public suspend fun getSSHKey(request: UserOuterClass.GetSSHKeyRequest, headers: Headers =
      emptyMap()): ResponseMessage<UserOuterClass.GetSSHKeyResponse>

  /**
   *  DeleteSSHKey removes a public SSH key.
   */
  public suspend fun deleteSSHKey(request: UserOuterClass.DeleteSSHKeyRequest, headers: Headers =
      emptyMap()): ResponseMessage<UserOuterClass.DeleteSSHKeyResponse>

  public suspend fun getGitToken(request: UserOuterClass.GetGitTokenRequest, headers: Headers =
      emptyMap()): ResponseMessage<UserOuterClass.GetGitTokenResponse>

  public suspend fun blockUser(request: UserOuterClass.BlockUserRequest, headers: Headers =
      emptyMap()): ResponseMessage<UserOuterClass.BlockUserResponse>
}
