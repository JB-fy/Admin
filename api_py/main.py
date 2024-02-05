from fastapi import FastAPI, Depends
from middleware import register_middleware
from exception import register_exception_handler
from router import register_router
from config import config
import uvicorn


from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker


engine = create_engine(config().database_default_url)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
Base = declarative_base()


from sqlalchemy import Column, String, Integer, Boolean, ForeignKey
from sqlalchemy.orm import relationship


class Test(Base):
    __tablename__ = "test"
    sceneId = Column(Integer, primary_key=True)
    sceneName = Column(String)
    sceneCode = Column(String, unique=True, index=True)
    sceneConfig = Column(String)
    isStop = Column(Boolean, default=False)
    rels = relationship("TestRel", back_populates="tests")


class TestRel(Base):
    __tablename__ = "test_rel"
    relId = Column(Integer, primary_key=True)
    relName = Column(String)
    sceneId = Column(Integer, ForeignKey("test.sceneId"))
    tests = relationship("Test", back_populates="rels")


from sqlalchemy.orm import Session


def get_db(group: str = "default"):
    async def func() -> Session:
        try:
            match group:
                case "default":
                    db = SessionLocal()
                case _:
                    db = SessionLocal()
            yield db
        finally:
            db.close()

    return func


def create_app():
    app = FastAPI(docs_url=None, redoc_url="/redoc")
    register_middleware(app)
    register_exception_handler(app)
    register_router(app)

    @app.get("/test")
    async def test(db: Session = Depends(get_db())):
        # from sqlalchemy import func
        # db.query(Test, Test.sceneName.label("label"), TestRel.relId, func.GROUP_CONCAT(TestRel.relId).label("relIdArr"))
        # .join(TestRel, TestRel.sceneId == Test.sceneId, isouter=True)
        # .outerjoin(TestRel, TestRel.sceneId == Test.sceneId)
        # .filter(Test.sceneId == 1)
        # .group_by(Test.sceneId)
        # .order_by(Test.sceneId.desc())
        # .offset(0)
        # .limit(1)

        """async for db in get_db()():
        query = db.query(Test)
        print(query)
        result = query.first()
        print(result)"""

        query = db.query(Test)
        print(query)
        result = query.first()
        print(result)
        return

    return app


# 启动方式（两种）
# 一般用于调试：uvicorn main:app --host=0.0.0.0 --port=8000 --reload
# 线上服务器用：python3.12 main.py
if __name__ == "__main__":
    uvicorn.run(
        app="main:create_app",
        host=config().server_http_host,
        port=config().server_http_port,
        reload=config().is_dev,
    )
