from fastapi import FastAPI, Request
from fastapi.encoders import jsonable_encoder
from fastapi.responses import JSONResponse, PlainTextResponse
from starlette.exceptions import HTTPException
from fastapi.exceptions import RequestValidationError

from .json_exception import JsonException
from .raw_exception import RawException


def register_exception_handler(app: FastAPI):
    @app.exception_handler(HTTPException)
    async def http_exception_handler(request: Request, exc: HTTPException):
        return JSONResponse({"code": exc.status_code, "msg": exc.detail, "data": {}})

    @app.exception_handler(RequestValidationError)
    async def validation_exception_handler(request: Request, exc: RequestValidationError):
        # return JSONResponse({"code": 200, "msg": (exc.errors())[0].msg, "data": {}})
        return JSONResponse({"code": 200, "msg": jsonable_encoder({"detail": exc.errors(), "body": exc.body}), "data": {}})

    @app.exception_handler(JsonException)
    async def json_exception_handler(request: Request, exc: JsonException):
        return JSONResponse({"code": exc.code, "msg": exc.msg, "data": exc.data})

    @app.exception_handler(RawException)
    async def raw_exception_handler(request: Request, exc: RawException):
        return PlainTextResponse(exc.msg)
