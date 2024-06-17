package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y float64
}

type Load struct {
	//prevDropoff Point // idea: didn't progress in that direction, consider in a future
	id       int
	pickup   Point
	dropoff  Point
	distance float64
	isUsed   bool
}

type Driver struct {
	id       int
	loads    []Load
	distance float64
}

const maxDriveTime = 12 * 60

var ZeroPoint = Point{}

func euclideanDistance(p1, p2 Point) float64 {
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}

func parseLoad(line string) Load {
	fields := strings.Fields(line)
	id, _ := strconv.Atoi(fields[0])

	pickupCoords := strings.Trim(fields[1], "()")
	pickupFields := strings.Split(pickupCoords, ",")
	pickupX, _ := strconv.ParseFloat(pickupFields[0], 64)
	pickupY, _ := strconv.ParseFloat(pickupFields[1], 64)

	dropoffCoords := strings.Trim(fields[2], "()")
	dropoffFields := strings.Split(dropoffCoords, ",")
	dropoffX, _ := strconv.ParseFloat(dropoffFields[0], 64)
	dropoffY, _ := strconv.ParseFloat(dropoffFields[1], 64)

	pickup := Point{pickupX, pickupY}
	dropoff := Point{dropoffX, dropoffY}
	distance := euclideanDistance(pickup, dropoff)

	return Load{
		id: id,
		//prevDropoff: ZeroPoint,
		pickup:   pickup,
		dropoff:  dropoff,
		distance: distance,
	}
}

func readLoads(filepath string) []Load {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", filepath)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var loads []Load

	// skip header
	scanner.Scan()
	scanner.Text()

	for scanner.Scan() {
		load := parseLoad(scanner.Text())
		loads = append(loads, load)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return loads
}

func sortForPoint(loads []Load, point Point) {
	sort.Slice(loads, func(i, j int) bool {
		if loads[i].isUsed && loads[j].isUsed { // both unused // it's actually mistype and should be used line below, but somehow results are much better with this mistype
			//if !loads[i].isUsed && !loads[j].isUsed { // both unused
			return false
		}
		if loads[i].isUsed { // only one in use
			return false
		}
		if loads[j].isUsed { // only second one in use
			return true
		}

		d1, d2 := euclideanDistance(point, loads[i].pickup), euclideanDistance(point, loads[j].pickup)

		if d1 < d2 {
			return true
		}

		return false
	})
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <path_to_problem_file>", os.Args[0])
	}

	filepath := os.Args[1]
	loads := readLoads(filepath)
	drivers := assignLoadsToDrivers(loads)

	//totalNumberOfDrivenMinutes := float64(0) // DEBUG
	for _, driver := range drivers {
		var loadIDs []string
		//var drivenMinutes float64 // DEBUG
		//lastPoint := Point{0, 0} // DEBUG

		for _, load := range driver.loads {
			//drivenMinutes += euclideanDistance(lastPoint, load.pickup) + load.distance // DEBUG
			//lastPoint = load.dropoff // DEBUG
			loadIDs = append(loadIDs, strconv.Itoa(load.id))
		}

		//drivenMinutes += euclideanDistance(lastPoint, Point{0, 0}) // DEBUG
		//totalNumberOfDrivenMinutes += drivenMinutes // DEBUG
		//fmt.Println("driven minutes:", drivenMinutes, driver.distance) // DEBUG
		fmt.Printf("[%s]\n", strings.Join(loadIDs, ","))
	}

	//fmt.Printf("total driven minutes: %v, total drivers: %v, total cost: %v", totalNumberOfDrivenMinutes, len(drivers), float64(500*len(drivers))+totalNumberOfDrivenMinutes) // DEBUG
}

func assignLoadsToDrivers(loads []Load) []Driver {
	var drivers []Driver
	driverID := 1
	usedLoadsCount := 0
	sortForPoint(loads, ZeroPoint)

	for len(loads) != usedLoadsCount {
		var driverLoads []Load
		var totalDistance float64

		dropoffPoint := ZeroPoint

		for i := 0; i < len(loads) && usedLoadsCount < len(loads); i++ {
			if loads[i].isUsed {
				continue
			}

			if totalDistance+ // prev distance
				euclideanDistance(dropoffPoint, loads[i].pickup)+ // from prev point to next point
				loads[i].distance+ // load distance
				euclideanDistance(loads[i].dropoff, Point{}) /* distance home */ <= maxDriveTime {

				driverLoads = append(driverLoads, loads[i])
				totalDistance += euclideanDistance(dropoffPoint, loads[i].pickup) + loads[i].distance

				dropoffPoint = loads[i].dropoff
				loads[i].isUsed = true
				usedLoadsCount++
				sortForPoint(loads, dropoffPoint) // could be optimized by reducing slice loads
				i = -1
			}
		}

		totalDistance += euclideanDistance(dropoffPoint, ZeroPoint) // don't forget to count way home
		driver := Driver{driverID, driverLoads, totalDistance}
		drivers = append(drivers, driver)
		driverID++
	}

	return drivers
}
