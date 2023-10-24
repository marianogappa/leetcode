#wip

class UndergroundSystem:

    def __init__(self):
        self.unfinished: dict[int, (int, str)] = {}
        self.trips: dict[(str, str), (int, int)] = defaultdict(tuple)
        

    def checkIn(self, id: int, stationName: str, t: int) -> None:
        self.unfinished[id] = (t, stationName)

    def checkOut(self, id: int, stationName: str, t: int) -> None:
        start_t, start_station = self.unfinished[id]
        trip_count, total_time = self.trips.get((start_station, stationName), (0, 0))
        self.trips[(start_station, stationName)] = (trip_count+1, total_time+(t-start_t))

    def getAverageTime(self, startStation: str, endStation: str) -> float:
        trip_count, total_time = self.trips[(startStation, endStation)]
        return total_time/trip_count


# Your UndergroundSystem object will be instantiated and called as such:
# obj = UndergroundSystem()
# obj.checkIn(id,stationName,t)
# obj.checkOut(id,stationName,t)
# param_3 = obj.getAverageTime(startStation,endStation)
