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


class ListItem(BaseModel):
    sceneId: int = Field(default=None, description="场景ID")
    sceneName: str = Field(default=None, description="场景名称")


class ListRes(BaseModel):
    list: list[ListItem]


# @router.get("/", tags=["平台后台/权限管理/场景"], summary="列表")
@router.post("/list", response_model=ListRes, tags=["平台后台/权限管理/场景"], summary="列表")
async def list(
    filter: Filter = Body(default=Filter()),
    field: list[str] = Body(default=[]),
    sort: str = Body(default="id DESC"),
    page: int = Body(default=1, gt=200),
    limit: int = Body(default=10),
):
    print(filter)
    filter = dict(
        builtins.filter(lambda item: item[1] is not None, vars(filter).items())
    )
    print(filter)

    raise JsonException(data={"list": [{"username": "Rick"}, {"username": "Morty"}]})
    return [{"username": "Rick"}, {"username": "Morty"}]
