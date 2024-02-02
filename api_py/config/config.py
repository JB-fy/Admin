from pydantic_settings import BaseSettings


class ConfigSettings(BaseSettings):
    is_dev: bool = False
    app_name: str = "JB Admin"
    super_platform_admin_id: int = 1

    database_default_url: str
    # database_default: dict

    class Config:
        env_file = ".env"
        extra = "ignore"
