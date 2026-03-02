include "./user/user_svc.thrift"
include "./system/system_svc.thrift"

namespace go inno_agent

service SystemService extends system_svc.SystemService {}
service UserService extends user_svc.UserService {}
