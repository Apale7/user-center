include "./base.thrift"
namespace go apale.user_center
namespace py apale.user_center

struct User {
    1: required i32 id,
    2: required string username,
    3: required string password,
    4: required i32 user_extra_id,
}

struct UserExtra {
    1: required i32 id,
    2: required string nickname,
    3: required string email,
    4: required string phone_number,
}

struct RegisterRequest {
    1: required User user,
    2: required UserExtra user_extra,

    255: base.Base base,
}

struct RegisterResponse {
    1: required i32 id,

    255: base.BaseResp base_resp,
}

struct LoginRequest {
    1: required string username,
    2: required string password,

    255: base.Base base,
}

struct LoginResponse {
    1: required i32 id,

    255: base.BaseResp base_resp,
}

struct LogoutRequest {
    1: required i32 id,

    255 : base.Base base,
}

struct LogoutResponse {
    255: base.BaseResp base_resp,
}

service UserCenter{
    RegisterResponse Register(1: RegisterRequest req),
    LoginResponse Login(1: LoginRequest req),
    LogoutResponse Logout(1: LogoutRequest req),
}