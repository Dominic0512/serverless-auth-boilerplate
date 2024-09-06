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

import { AuthGenerateAuthURLResponse, AuthTokenResponse } from "./data-contracts";
import { ContentType, HttpClient, RequestParams } from "./http-client";

export class Auth<SecurityDataType = unknown> extends HttpClient<SecurityDataType> {
  /**
   * @description Currently, the authorization is integrated with Auth0. This endpoint will generate an authorization URL for the client to redirect to the Auth0 login page.
   *
   * @tags Auth
   * @name GenerateAuthUrl
   * @summary Generate oauth login url
   * @request GET:/auth/oauth-url
   */
  generateAuthUrl = (params: RequestParams = {}) =>
    this.request<AuthGenerateAuthURLResponse, any>({
      path: `/auth/oauth-url`,
      method: "GET",
      type: ContentType.Json,
      format: "json",
      ...params,
    });
  /**
   * @description SignIn with oauth code
   *
   * @tags Auth
   * @name SignIn
   * @summary SignIn with oauth code
   * @request POST:/auth/sign-in
   */
  signIn = (params: RequestParams = {}) =>
    this.request<AuthTokenResponse, any>({
      path: `/auth/sign-in`,
      method: "POST",
      type: ContentType.Json,
      format: "json",
      ...params,
    });
  /**
   * @description SignUp with oauth code
   *
   * @tags Auth
   * @name SignUp
   * @summary SignUp with oauth code
   * @request POST:/auth/sign-up
   */
  signUp = (params: RequestParams = {}) =>
    this.request<AuthTokenResponse, any>({
      path: `/auth/sign-up`,
      method: "POST",
      type: ContentType.Json,
      format: "json",
      ...params,
    });
}
