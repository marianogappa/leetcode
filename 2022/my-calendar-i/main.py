# The issue of minimising overlap checks requires keeping the
# bookings in some sort order.
#
# There are no alternatives to a BST, AVL Tree, R/B Tree for
# efficiently keeping the list sorted while adding elements,
# and implementing this would be too tricky so we resort to
# SortedList, which even comes with binary search built-in.
#
# Note that SortedList receives tuples and sorts by the first
# element of the tuple ascendingly, so bookings will be sorted
# by start ascending.
#
# We know there are no overlaps in bookings, which means that,
# when bisect returns, we only need to check previous and
# current index to know whether there's overlap. Careful that
# left and current may both not exist, and that's it.

from sortedcontainers import SortedList

# Time: O(nlogn) for n calls of book, each one being O(logn)
# Space: O(n) to build and maintain the SortedList
class MyCalendar:

    def __init__(self):
        self.bookings = SortedList()

    def book(self, start: int, end: int) -> bool:
        cur_booking = (start, end)
        idx = self.bookings.bisect_right(cur_booking)

        if (idx > 0 and is_overlap(cur_booking, self.bookings[idx-1])) or (idx < len(self.bookings) and is_overlap(cur_booking, self.bookings[idx])):
            return False

        self.bookings.add((start, end))
        return True


def is_overlap(booking1: tuple[int], booking2: tuple[int]) -> bool:
    return latter_booking(booking1, booking2)[0] < earlier_booking(booking1, booking2)[1]


def earlier_booking(booking1: tuple[int], booking2: tuple[int]) -> tuple[int]:
    if booking1[0] <= booking2[0]:
        return booking1
    return booking2


def latter_booking(booking1: tuple[int], booking2: tuple[int]) -> tuple[int]:
    if booking1[0] > booking2[0]:
        return booking1
    return booking2
