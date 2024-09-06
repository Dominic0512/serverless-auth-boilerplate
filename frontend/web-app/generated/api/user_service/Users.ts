/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

import { UserUserResponse, UserUsersResponse } from "./data-contracts";
import { ContentType, HttpClient, RequestParams } from "./http-client";

export class Users<SecurityDataType = unknown> extends HttpClient<SecurityDataType> {
  /**
   * @description List users
   *
   * @tags User
   * @name ListUsers
   * @summary List users
   * @request GET:/users
   */
  listUsers = (params: RequestParams = {}) =>
    this.request<UserUsersResponse, any>({
      path: `/users`,
      method: "GET",
      type: ContentType.Json,
      format: "json",
      ...params,
    });
  /**
   * @description Create user
   *
   * @tags User
   * @name CreateUser
   * @summary Create user
   * @request POST:/users
   */
  createUser = (params: RequestParams = {}) =>
    this.request<UserUserResponse, any>({
      path: `/users`,
      method: "POST",
      type: ContentType.Json,
      format: "json",
      ...params,
    });
  /**
   * @description Get user by id
   *
   * @tags User
   * @name GetUser
   * @summary Get user by id
   * @request GET:/users/{id}
   */
  getUser = (id: string, params: RequestParams = {}) =>
    this.request<UserUserResponse, any>({
      path: `/users/${id}`,
      method: "GET",
      type: ContentType.Json,
      format: "json",
      ...params,
    });
  /**
   * @description Update user
   *
   * @tags User
   * @name UpdateUser
   * @summary Update user
   * @request PUT:/users/{id}
   */
  updateUser = (id: string, params: RequestParams = {}) =>
    this.request<UserUserResponse, any>({
      path: `/users/${id}`,
      method: "PUT",
      type: ContentType.Json,
      format: "json",
      ...params,
    });
  /**
   * @description Delete user
   *
   * @tags User
   * @name DeleteUser
   * @summary Delete user
   * @request DELETE:/users/{id}
   */
  deleteUser = (id: string, params: RequestParams = {}) =>
    this.request<void, any>({
      path: `/users/${id}`,
      method: "DELETE",
      type: ContentType.Json,
      ...params,
    });
  /**
   * @description Partial update user
   *
   * @tags User
   * @name PartialUpdateUser
   * @summary Partial update user
   * @request PATCH:/users/{id}
   */
  partialUpdateUser = (id: string, params: RequestParams = {}) =>
    this.request<UserUserResponse, any>({
      path: `/users/${id}`,
      method: "PATCH",
      type: ContentType.Json,
      format: "json",
      ...params,
    });
}
