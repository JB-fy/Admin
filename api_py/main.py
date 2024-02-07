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


from sqlalchemy import Column, String, Integer, Boolean, ForeignKey, DateTime
from sqlalchemy.orm import relationship


class Action(Base):
    __tablename__ = "auth_action"
    actionId = Column(Integer, primary_key=True)
    actionName = Column(String)
    actionCode = Column(String, unique=True, index=True)
    remark = Column(String)
    isStop = Column(Boolean, default=False)
    updatedAt = Column(DateTime)
    createdAt = Column(DateTime)
    rels = relationship("ActionRelToScene", back_populates="actions")


class ActionRelToScene(Base):
    __tablename__ = "auth_action_rel_to_scene"
    actionId = Column(Integer, ForeignKey("auth_action.actionId"), primary_key=True)
    sceneId = Column(Integer, primary_key=True)
    updatedAt = Column(DateTime)
    createdAt = Column(DateTime)
    actions = relationship("Action", back_populates="rels")


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
        """from sqlalchemy import func
        (
            db.query(Action, Action.actionName.label("label"), func.GROUP_CONCAT(ActionRelToScene.sceneId).label("sceneIdArr"))
            .join(ActionRelToScene, ActionRelToScene.actionId == Action.actionId, isouter=True)
            .outerjoin(ActionRelToScene, ActionRelToScene.actionId == Action.actionId)
            .filter(Action.actionId == 1)
            .group_by(Action.actionId)
            .order_by(Action.actionId.desc())
            .offset(0)
            .limit(1)
        )"""

        """async for db in get_db()():
        query = db.query(Action)
        print(query)
        result = query.first()
        print(result)"""

        query = db.query(Action)
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
