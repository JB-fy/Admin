from fastapi import Request
from exception.json_exception import JsonException


async def scene(request: Request):
    pathArr = str.split(request.url.path, "/")
    sceneCode = pathArr[1]
    if sceneCode == "":
        raise JsonException(39999998)
    """ sceneInfo, _ := db.query(Scene).filter(Scene.sceneCode == sceneCode).first()
    if sceneInfo.IsEmpty() {
		r.SetError(utils.NewErrorCode(r.GetCtx(), 39999998, ``))
		return
	}
	if sceneInfo[daoAuth.Scene.Columns().IsStop].Uint() == 1 {
		r.SetError(utils.NewErrorCode(r.GetCtx(), 39999997, ``))
		return
	} """
    # request.sceneInfo = {"sceneId":1, "sceneCode":sceneCode}
    # 官方建议用state变量增加自定义内容
    request.state.sceneInfo = {"sceneId": 1, "sceneCode": sceneCode}
