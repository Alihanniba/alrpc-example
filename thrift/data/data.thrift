namespace go data

struct Args {
    1: required i32 age;
    2: required string name;
}

struct Reply {
    1: string title;
}

service RpcService {
    Reply hello(1: Args req)
}