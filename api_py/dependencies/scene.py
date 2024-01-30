from fastapi import Request
from exception.json_exception import JsonException


async def scene(request: Request):
    pathArr = str.split(request.url.path, "/")
    sceneCode = pathArr[1]
    if sceneCode == "":
        raise JsonException(39999998)
    """
	sceneInfo, _ := daoAuth.Scene.ParseDbCtx(r.GetCtx()).Where(daoAuth.Scene.Columns().SceneCode, sceneCode).One()
	if sceneInfo.IsEmpty() {
		r.SetError(utils.NewErrorCode(r.GetCtx(), 39999998, ``))
		return
	}
	if sceneInfo[daoAuth.Scene.Columns().IsStop].Uint() == 1 {
		r.SetError(utils.NewErrorCode(r.GetCtx(), 39999997, ``))
		return
	}
	utils.SetCtxSceneInfo(r, sceneInfo) """
