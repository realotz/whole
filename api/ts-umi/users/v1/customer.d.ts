// @ts-ignore
/* eslint-disable */
// Code generated by protoc-gen-ts-umi. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.

declare namespace UsersV1 {
	/** CustomerDeleteOption */
	type CustomerDeleteOption = {
		ids?:Array<number>
	}
	/** CustomerList */
	type CustomerList = {
		list?:Array<UsersV1.Customer>
	}
	/** Customer */
	type Customer = {
		id?:number
		uuid?:string
		name?:string
		account?:string
		nick_name?:string
		role?:string
		email?:string
		mobile?:string
		id_card?:string
		sex?:string
		birthday?:GoogleProtobuf.Timestamp
		last_ip?:string
		last_time?:GoogleProtobuf.Timestamp
		create_time?:GoogleProtobuf.Timestamp
		update_time?:GoogleProtobuf.Timestamp
		avatar?:string
	}
	/** CustomerLogin */
	type CustomerLogin = {
		account?:string
		password?:string
		captcha?:string
		code?:string
	}
	/** CustomerLoginRes */
	type CustomerLoginRes = {
		token?:string
		role?:string
		expiration_time?:number
		customer?:UsersV1.Customer
	}
	/** CustomerListOption */
	type CustomerListOption = {
		name?:string
		tag?:string
		email?:string
		mobile?:string
		sex?:string
		id_card?:string
		keyword?:string
		create_time?:string
	}
	/** CustomerGetOption */
	type CustomerGetOption = {
		id?:number
	}
	/** CustomerOption */
	type CustomerOption = {
		name?:string
		account?:string
		nick_name?:string
		role?:string
		email?:string
		mobile?:string
		id_card?:string
		password?:string
		sex?:string
		birthday?:GoogleProtobuf.Timestamp
		id?:number
		avatar?:string
	}
}

declare namespace GoogleProtobuf {
	/** Timestamp */
	type Timestamp = {
		seconds?:number
		nanos?:number
	}
}

