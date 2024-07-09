from sqlite3 import Connection

class TripsRepository: 
    def __init__(self, conn: Connection) -> None:
        self.__conn = conn