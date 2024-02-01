from fastapi import APIRouter, Body
from exception.json_exception import JsonException
from pydantic import BaseModel, Field, validator
import builtins
from datetime import date, datetime
from enum import Enum

router = APIRouter(
    prefix="/auth/scene",
)


class IsStop(int, Enum):
    NO = 0
    YES = 1


class Filter(BaseModel):
    id: int = Field(default=None, gt=0, description="ID")
    idArr: list[int] = Field(default=None, min_length=1, description="ID数组")
    excId: int = Field(default=None, gt=0, description="排除ID")
    excIdArr: list[int] = Field(default=None, min_length=1, description="排除ID数组")
    label: str = Field(
        default=None,
        max_length=30,
        pattern="^[\\p{L}\\p{M}\\p{N}_-]+$",
        description="标签。常用于前端组件",
    )
    sceneId: int = Field(default=None, gt=0, description="场景ID")
    sceneName: str = Field(default=None, max_length=30, description="场景名称")
    # isStop: int = Field(default=None, ge=0, le=1, description="停用：0否 1是")
    isStop: IsStop = Field(default=None, description="停用：0否 1是")
    timeRangeStart: datetime = Field(
        default=None, description="开始时间：YYYY-mm-dd HH:ii:ss"
    )
    timeRangeEnd: datetime = Field(
        default=None, description="结束时间：YYYY-mm-dd HH:ii:ss"
    )

    @validator("idArr", "excIdArr")
    def validator_unique(cls, value):
        if len(value) != len(set(value)):
            raise ValueError("不能有重复值")
        return value

    @validator("timeRangeEnd")
    def validator_datetime(cls, value, values):
        if (
            value is not None
            and "timeRangeStart" in values
            and value < values["timeRangeStart"]
        ):
            raise ValueError("结束时间不能小于开始时间")
        return value


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
    print(req.filter)
    filter = dict(
        builtins.filter(lambda item: item[1] is not None, vars(req.filter).items())
    )
    print(filter)
    res = ListRes()
    # res.count = 0
    res.list = [{"sceneName": "平台后台"}]
    # return res
    raise JsonException(data=dict(res))
