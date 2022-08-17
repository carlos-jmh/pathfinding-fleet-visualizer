package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"pathfinding-fleet-visualizer/internal/server/utils"
)

type PageInputData struct {
	Start int     `json:"start"`
	End   int     `json:"end"`
	Map   GameMap `json:"gameMap"`
}

type GameMap struct {
	Rows  int
	Cols  int
	Tiles [][]Tile
}

type Tile struct {
	Id       int
	Weight   float64
	Disabled bool
}

type PageOutputData struct {
	Explored []int  `json:"explored"`
	Shortest []int  `json:"shortest"`
	Error    string `json:"error"`
}

// GetPath calls the getDijkstra to find the shortest path
func GetPath(c *gin.Context) {
	pageOutputData := getDijkstra(c)
	c.IndentedJSON(http.StatusOK, pageOutputData)
}

// getDijkstra calls dijkstra to find the shortest path
func getDijkstra(c *gin.Context) *PageOutputData {
	var pageInputData PageInputData
	var pageOutputData = PageOutputData{}

	// Bind receiving JSON to the PageInputData struct
	err := c.BindJSON(&pageInputData)
	if err != nil {
		pageOutputData.Error = err.Error()
		return &pageOutputData
	}

	err = validatePageInputData(&pageInputData)
	if err != nil {
		pageOutputData.Error = err.Error()
		return &pageOutputData
	}

	explored, shortest := pageInputData.Map.dijkstra(pageInputData.Start, pageInputData.End)
	pageOutputData.Explored = explored
	pageOutputData.Shortest = shortest

	return &pageOutputData
}

// dijkstra uses Dijkstra's algorithm to find the shortest path
func (gameMap *GameMap) dijkstra(start, target int) ([]int, []int) { // replace map[int]int with Path later
	var exploredInOrder []int
	var shortestPath []int

	// to keep track of visited Tiles
	var visited = make(map[int]bool)

	// to reconstruct the path from end Tile to start Tile
	var parentsMap = make(map[int]int)

	// to store the weight of each Tile
	var nodeCosts = make(map[int]float64)

	// Initialize the nodeCosts map with the start tile
	for _, tileRow := range gameMap.Tiles {
		for _, tile := range tileRow {
			nodeCosts[tile.Id] = math.Inf(1)
		}
	}
	nodeCosts[start] = 0

	// Initialize the priority queue with the start tile
	var pq = utils.NewHeap()
	pq.Push(&utils.Item{Id: start, Weight: 0})

	for len(*pq.Items) > 0 {
		// Get the tile with the lowest weight
		item := pq.Pop()
		tile := gameMap.getTileById(item.Id)

		if visited[tile.Id] {
			continue
		}

		if tile.Id == target {
			shortestPath = constructPath(parentsMap, target)
			return exploredInOrder, shortestPath
		}

		for _, neighborTile := range gameMap.getNeighbors(tile.Id) {
			if visited[neighborTile.Id] {
				continue
			}

			newCost := nodeCosts[tile.Id] + neighborTile.Weight
			if newCost < nodeCosts[neighborTile.Id] {
				parentsMap[neighborTile.Id] = tile.Id
				nodeCosts[neighborTile.Id] = newCost

				pq.Push(&utils.Item{
					Id:     neighborTile.Id,
					Weight: newCost,
				})
			}
		}

		visited[tile.Id] = true
		exploredInOrder = append(exploredInOrder, tile.Id)
	}

	shortestPath = constructPath(parentsMap, target)

	//return parentsMap, nodeCosts
	return exploredInOrder, shortestPath
}

func validatePageInputData(data *PageInputData) error {
	if data.Map.Rows <= 0 || data.Map.Cols <= 0 {
		return errors.New("map has no rows or columns")
	}
	if data.Start < 1 || data.Start > data.Map.Rows*data.Map.Cols {
		return errors.New("start is out of bounds")
	}
	if data.End < 1 || data.End > data.Map.Rows*data.Map.Cols {
		return errors.New("end is out of bounds")
	}
	if data.Start == data.End {
		return errors.New("start and end are the same")
	}

	return nil
}

// getTileById returns the tile with the given id
func (gameMap *GameMap) getTileById(id int) Tile {
	row := int(math.Floor(float64((id - 1) / gameMap.Cols)))
	col := (id - 1) % gameMap.Cols
	return gameMap.Tiles[row][col]
}

// getNeighbors returns the neighbors surrounding a tile
func (gameMap *GameMap) getNeighbors(id int) []Tile {
	row := int(math.Floor(float64((id - 1) / gameMap.Cols)))
	col := (id - 1) % gameMap.Cols
	var neighbors []Tile

	if row > 0 {
		newTile := gameMap.Tiles[row-1][col]
		if newTile.Disabled == false {
			neighbors = append(neighbors, newTile)
		}
	}
	if row < gameMap.Rows-1 {
		newTile := gameMap.Tiles[row+1][col]
		if newTile.Disabled == false {
			neighbors = append(neighbors, newTile)
		}
	}
	if col > 0 {
		newTile := gameMap.Tiles[row][col-1]
		if newTile.Disabled == false {
			neighbors = append(neighbors, newTile)
		}
	}
	if col < gameMap.Cols-1 {
		newTile := gameMap.Tiles[row][col+1]
		if newTile.Disabled == false {
			neighbors = append(neighbors, newTile)
		}
	}

	return neighbors
}

// constructPath constructs a path from start to target using a parentsMap
func constructPath(parentsMap map[int]int, target int) []int {
	var shortestPath []int

	// Construct the path backwards from the target
	for target != 0 {
		shortestPath = append([]int{target}, shortestPath...)
		target = parentsMap[target]
	}
	return shortestPath
}
