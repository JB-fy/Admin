from fastapi import FastAPI
from middleware import register_middleware
from exception import register_exception_handler
from router import register_router


app = FastAPI(docs_url=None, redoc_url="/redoc")
register_middleware(app)
register_exception_handler(app)
register_router(app)
