// @ts-ignore
/* eslint-disable */
// Code generated by protoc-gen-ts-umi. DO NOT EDIT.
import {request} from 'umi';

const APIService = '/api';
// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.

type Options = {
  [key: string]: any
}

/** CreateCategory  创建文章类目 /api */
export async function CreateCategory(params: ApiCmsV1.Category, options?: Options) {
	return request<ApiCmsV1.CreateCategoryReply>(APIService + '/api.cms.v1.CategoryService/CreateCategory', {
    	method: 'POST',
    	params: {...params},
    	...(options || {}),
	});
}

/** UpdateCategory  更新文章类目 /api */
export async function UpdateCategory(params: ApiCmsV1.Category, options?: Options) {
	return request<ApiCmsV1.UpdateCategoryReply>(APIService + '/api.cms.v1.CategoryService/UpdateCategory', {
    	method: 'POST',
    	params: {...params},
    	...(options || {}),
	});
}

/** DeleteCategory  批量删除文章类目 /api */
export async function DeleteCategory(params: ApiCmsV1.DeleteCategoryRequest, options?: Options) {
	return request<ApiCmsV1.DeleteCategoryReply>(APIService + '/api.cms.v1.CategoryService/DeleteCategory', {
    	method: 'POST',
    	params: {...params},
    	...(options || {}),
	});
}

/** GetCategory  获取文章类目详情 /api */
export async function GetCategory(params: ApiCmsV1.GetCategoryRequest, options?: Options) {
	return request<ApiCmsV1.Category>(APIService + '/api.cms.v1.CategoryService/GetCategory', {
    	method: 'POST',
    	params: {...params},
    	...(options || {}),
	});
}

/** ListCategory  查询文章类目列表 /api */
export async function ListCategory(params: ApiCmsV1.ListCategoryRequest, options?: Options) {
	return request<ApiCmsV1.ListCategoryReply>(APIService + '/api.cms.v1.CategoryService/ListCategory', {
    	method: 'POST',
    	params: {...params},
    	...(options || {}),
	});
}

