// @ts-ignore
/* eslint-disable */
// Code generated by protoc-gen-ts-umi. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.

declare namespace NewsV1 {
	/** NoticeListOption */
	type NoticeListOption = {
		type?:number
		total?:number
		page?:number
		page_size?:number
	}
	/** Notice */
	type Notice = {
		type?:string
		title?:string
		content?:string
		to?:string
		status?:number
		error_msg?:string
		birthday?:GoogleProtobuf.Timestamp
	}
	/** NoticeCreateOption */
	type NoticeCreateOption = {
		type?:string
		title?:string
		content?:string
		to?:string
	}
	/** NoticeList */
	type NoticeList = {
		list?:Array<NewsV1.Notice>
		total?:number
		page?:number
		page_size?:number
	}
}

declare namespace GoogleProtobuf {
	/** Timestamp */
	type Timestamp = {
		seconds?:number
		nanos?:number
	}
}
