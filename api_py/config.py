from pydantic_settings import BaseSettings
from functools import lru_cache


@lru_cache
def config():
    return ConfigSettings()


class ConfigSettings(BaseSettings):
    is_dev: bool = False
    app_name: str = "JB Admin"
    super_platform_admin_id: int = 1

    server_http_host: str = "0.0.0.0"
    server_http_port: int = 8000

    database_default_url: str
    # database_default: dict

    class Config:
        env_file = ".env"
        extra = "ignore"
