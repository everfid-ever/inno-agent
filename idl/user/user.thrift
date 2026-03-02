namespace go basicuser

include "../base.thrift"

/*
 用户服务 - common struct
 */

// 用户 Profile（字段暂时留空）
struct UserProfile {
}

// 用户登录请求
// authType: phone-password / phone-verify / email-password / email-verify / studentId-password / studentId-verify
struct LoginReq {
    1: required string authType,
    2: required string authId,
    3: required string verify,
}

// 用户登录响应
struct LoginResp {
    1: required base.Response resp,
    2: optional string        token,
    3: optional bool          isNew,
    4: optional string        name,
    5: optional string        avatar,
}

// 用户注册请求
struct RegisterReq {
    1: required string authType,
    2: required string authId,
    3: required string verify,
    4: optional string password,
}

// 用户注册响应
struct RegisterResp {
    1: required base.Response resp,
    2: optional string        token,
}

// 重置密码请求（需登录态）
struct ResetPasswordReq {
    1: required string newPassword,
}

// 重置密码响应
struct ResetPasswordResp {
    1: required base.Response resp,
}

// 获取个人资料请求
struct GetProfileReq {
}

// 获取个人资料响应
struct GetProfileResp {
    1: required base.Response resp,
    2: optional string        name,
    3: optional string        avatar,
    4: optional UserProfile   profile,
}

// 更新个人资料请求
struct UpdateProfileReq {
    1: optional string      name,
    2: optional string      avatar,
    3: optional UserProfile profile,
}

// 更新个人资料响应
struct UpdateProfileResp {
    1: required base.Response resp,
}