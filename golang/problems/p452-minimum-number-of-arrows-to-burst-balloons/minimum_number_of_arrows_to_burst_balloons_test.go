package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumNumberOfArrowsToBurstBalloons(t *testing.T) {
	type data struct {
		points [][]int
		except int
	}

	testData := []data{
		// {
		// 	points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
		// 	except: 2,
		// },
		// {
		// 	points: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		// 	except: 4,
		// },
		// {
		// 	points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
		// 	except: 2,
		// },
		{
			points: [][]int{{3,9},{7,12},{3,8},{6,8},{9,10},{2,9},{0,9},{3,9},{0,6},{2,8}},
			except: 2,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, findMinArrowShots(d.points))
	}
}
