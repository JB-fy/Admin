from fastapi import APIRouter, Depends
from dependencies import scene, scene_login_of_platform
from controller.platform.auth.scene import router as routerAuthScene

router = APIRouter(
    prefix="/platform",
    dependencies=[Depends(scene), Depends(scene_login_of_platform(True))],
)

router.include_router(routerAuthScene)
""" routerAuth=APIRouter(prefix="/auth")
routerAuth.include_router(routerAuthScene)
router.include_router(routerAuth) """
