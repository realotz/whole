// @ts-ignore
/* eslint-disable */
// Code generated by protoc-gen-ts-umi. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.

declare namespace AdminV1 {
	/** EmployeeList */
	type EmployeeList = {
		list?:Array<AdminV1.Employee>
		total?:number
		page?:number
		page_size?:number
	}
	/** Employee */
	type Employee = {
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
	/** CaptchaReq */
	type CaptchaReq = {
		to?:string
		captcha?:string
		challenge?:string
		tk?:string
	}
	/** EmployeeLogin */
	type EmployeeLogin = {
		account?:string
		password?:string
		captcha?:string
		code?:string
	}
	/** EmployeeLoginRes */
	type EmployeeLoginRes = {
		token?:string
		role?:string
		expiration_time?:number
		employee?:AdminV1.Employee
	}
	/** EmployeeListOption */
	type EmployeeListOption = {
		name?:string
		tag?:string
		email?:string
		mobile?:string
		sex?:string
		role?:string
		type?:string
		create_time?:string
	}
	/** EmployeeOption */
	type EmployeeOption = {
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
	/** EmployeeDeleteOption */
	type EmployeeDeleteOption = {
		ids?:Array<number>
	}
	/** EmployeeGetOption */
	type EmployeeGetOption = {
		id?:number
	}
}

declare namespace GoogleProtobuf {
	/** Timestamp */
	type Timestamp = {
		seconds?:number
		nanos?:number
	}
}
