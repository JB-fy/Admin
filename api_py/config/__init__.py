from functools import lru_cache
from .config import ConfigSettings

@lru_cache
def config():
    return ConfigSettings()