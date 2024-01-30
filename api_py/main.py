from fastapi import FastAPI
from middleware import register_middleware
from exception import register_exception_handler
from router import register_router
import uvicorn

app = FastAPI(docs_url=None, redoc_url="/redoc")
register_middleware(app)
register_exception_handler(app)
register_router(app)

# 启动方式（两种）
# 一般用于调试：uvicorn main:app --host=0.0.0.0 --port=8000 --reload
# 线上服务器用：python3.12 main.py
if __name__ == "__main__":
    uvicorn.run(app="main:app", host="0.0.0.0", port=8000)
