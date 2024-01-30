from fastapi import FastAPI
from middleware import register_middleware
from exception import register_exception_handler
from router import register_router
import uvicorn

app = FastAPI(docs_url=None, redoc_url="/redoc")
register_middleware(app)
register_exception_handler(app)
register_router(app)

if __name__ == "__main__":
    uvicorn.run(app="main:app", host="0.0.0.0", port=8000, reload=True, workers=1)
