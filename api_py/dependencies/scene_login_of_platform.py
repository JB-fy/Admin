from fastapi import Header
from exception.json_exception import JsonException


def scene_login_of_platform(isForce: bool):
    async def func(token: str = Header(default=None, alias="PlatformToken")):
        if token:
            if isForce:
                raise JsonException(39994000)

    return func
