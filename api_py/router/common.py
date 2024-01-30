from fastapi import APIRouter, Depends
from fastapi.responses import RedirectResponse

router = APIRouter()


@router.get("/")
async def root():
    return RedirectResponse(url="/admin/platform", status_code=302)
