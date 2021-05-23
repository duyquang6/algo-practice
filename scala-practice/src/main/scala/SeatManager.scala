import collection.mutable.PriorityQueue

class SeatManager(_n: Int) {
    val pq = PriorityQueue[Int]((-_n to -1):_*)
    def reserve(): Int = {
        pq.dequeue()
    }

    def unreserve(seatNumber: Int) {
        pq.enqueue(seatNumber)
    }
}