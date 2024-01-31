from pydantic_settings import BaseSettings


class ConfigSettings(BaseSettings):
    is_dev: bool
    app_name: str
    super_platform_admin_id: int

    class Config:
        env_file = ".env"
        extra = "ignore"
