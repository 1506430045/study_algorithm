package queue

type MaxQueue struct {
	q1 *QueueLink
	q2 *QueueLink
}

func Constructor() MaxQueue {
	return MaxQueue{
		q1: InitQueueLink(),
		q2: InitQueueLink(),
	}
}

func (this *MaxQueue) Max_value() int {
	return 0
}

func (this *MaxQueue) Push_back(value int) {

}

func (this *MaxQueue) Pop_front() int {
	return 0
}
