struct Request {
	1: string uuid
}

struct UserFeatureItem {
    1: list<string> clickList,
    2: list<string> clickTimeList
    3: list<string> playList
    4: list<string> playTimeList
    5: list<string> playListT1
    6: list<string> playTimeListT1
    7: map<string,i64> playIdx
}

service UserFeatureRPCService {
    UserFeatureItem GetUserFeature(1: Request req)
}