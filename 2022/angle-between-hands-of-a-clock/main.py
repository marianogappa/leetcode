# Time: O(1)
# Space: O(1)
class Solution:
    def angleClock(self, hour: int, minutes: int) -> float:
        angle_of_an_hour = 30.0
        hour_hand_angle = ((hour + (minutes / 60.0)) * angle_of_an_hour) % 360
        minute_hand_angle = (minutes / 60.0) * 360.0
        
        return min(abs(hour_hand_angle - minute_hand_angle), 360 - abs(minute_hand_angle - hour_hand_angle))
