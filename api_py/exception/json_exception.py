class JsonException(Exception):
    def __init__(self, code: int = 0, msg: str = "", data: dict = {}):
        self.code = code
        self.msg = msg
        self.data = data
        if self.msg == "":
            self.msg = "成功" if self.code == 0 else "失败"
