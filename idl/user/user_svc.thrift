namespace go basicuser

include "user.thrift"

// 用户相关接口
service UserService {
    // 登录（支持 phone / email / studentId 三种 authType）
    user.LoginResp Login(1: user.LoginReq req),

    // 注册
    user.RegisterResp Register(1: user.RegisterReq req),

    // 重置密码（需登录态，通过 Authorization header 鉴权）
    user.ResetPasswordResp ResetPassword(1: user.ResetPasswordReq req),

    // 获取个人资料（需登录态）
    user.GetProfileResp GetProfile(1: user.GetProfileReq req),

    // 更新个人资料（需登录态）
    user.UpdateProfileResp UpdateProfile(1: user.UpdateProfileReq req),
}