package queue

import (
	cq "github.com/EbicHecker/queue"
	"github.com/FUNHAVER-Gaming/game-generator/pkg/models"
	"time"
)

// Have 1 Queue: Global Queue
// For the first 5 minutes, try to match roles as best as possible
// In between 5 and 10 minutes, loosen the restrictions on the "optimized" team rules for RBMM
// After 10 minutes, forgo any role based MM, and simply put 10 people together
// If the players time in the queue is > 1.5 hours, remove from queue.

// Have 3 queues: Global Queue, Short Term, and Long Term queue.
// - Short Term queue: Holds players for 10 minutes, and tries to find a role based optimized match
// - Long Term queue: Holds players for 1.5 hours (but greater than 10 minutes), and tries to simply match 10 players together
// - Global Queue: Holds all players, and is used to track averages / roles based in queue

// Have 5 Queues: Controller Queue, Sentinel Queue, Duelist Queue, Initiator Queue, Any Role Queue
// - Queues would fill with respective roles, and the system would use the established role based MM to create games based off the content of those queues
// - Players in the "Any Role" queue would be counter for any role that is required to make an optimized match
// - Players can queue for 2 roles at the same time. IE> A duelist is fine playing initiator, but not controller
// - Queues are prioritized based off of time spent in the queue, not rating

var queue *cq.ConcurrentQueue[*models.Player]

func init() {
	queue = cq.NewConcurrentQueue[*models.Player]()
}

// AddToQueue add a player to the queue, and return their spot in it
func AddToQueue(player *models.Player) int {
	queue.Enqueue(player)
	return queue.Len()
}

func GetQueuePosition(player *models.Player) int {
	for index, p := range queue.ToSlice() {
		if p.UserID == player.UserID {
			return index
		}
	}
	return -1
}

func queueRunner() {
	for {
		if queue.IsEmpty() {
			time.Sleep(1 * time.Second)
			continue
		}

		//If there aren't 10 people queueing, we do not care from here
		if queue.Len() < 10 {
			time.Sleep(500 * time.Millisecond)
			continue
		}

		//There is EXACTLY 10 people in queue, put them into a game
		if queue.Len() == 10 {
			//TODO insert into game
			continue
		}

	}
}
