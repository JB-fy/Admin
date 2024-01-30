from pydantic_settings import BaseSettings


class ConfigSettings(BaseSettings):
    app_env: str = "dev"
    app_name: str = "JB Admin"
    super_platform_admin_id: int = 1

    class Config:
        env_file = ".env"
        extra = "ignore"
