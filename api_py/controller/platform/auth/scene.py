from fastapi import APIRouter, Body
from exception.json_exception import JsonException
from pydantic import BaseModel, Field
import builtins

router = APIRouter(
    prefix="/auth/scene",
)


class Filter(BaseModel):
    # sceneId: int | None = Field(default=None, description="场景ID")
    sceneId: int = Field(default=None, description="场景ID")
    sceneName: str = Field(default=None, description="场景名称")


class ListReq(BaseModel):
    filter: Filter = Body(default=Filter(), description="过滤条件")
    field: list[str] = Body(
        default=[],
        description="查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力",
    )
    sort: str = Body(default="id DESC", description="排序")
    page: int = Body(default=1, gt=0, description="页码")
    limit: int = Body(default=10, ge=0, description="每页数量。可传0取全部")


class ListItem(BaseModel):
    sceneId: int = Field(default=None, description="场景ID")
    sceneName: str = Field(default=None, description="场景名称")


class ListRes(BaseModel):
    count: int = Field(default=0, description="总数")
    list: builtins.list[ListItem] = Field(default=[], description="列表")


# @router.get("/", response_model=ListRes, tags=["平台后台/权限管理/场景"], summary="列表")
@router.post("/list", tags=["平台后台/权限管理/场景"], summary="列表")
async def list(req: ListReq) -> ListRes:
    print(req)
    filter = dict(
        builtins.filter(lambda item: item[1] is not None, vars(req.filter).items())
    )
    print(req.filter)
    print(filter)
    res = ListRes()
    # res.count = 0
    res.list = [{"sceneName": "平台后台"}]
    # return res
    raise JsonException(data=dict(res))
