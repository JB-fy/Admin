from fastapi import FastAPI
from .common import router as routerCommon
from .platform import router as routerPlatform


def register_router(app: FastAPI):
    app.include_router(routerCommon)
    app.include_router(routerPlatform)
