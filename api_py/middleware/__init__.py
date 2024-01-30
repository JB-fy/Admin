from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from .i18n import I18n
from .log import Log


def register_middleware(app: FastAPI):
    app.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )
    """ @app.middleware("i18n")
    async def i18n(request: Request, call_next):
        response = await call_next(request)
        return response """
    app.add_middleware(I18n)
    # app.add_middleware(Log)
