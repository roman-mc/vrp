package main

// NOTE: greedy (most inefficient)
//func assignLoadsToDrivers(loads []Load) []Driver {
//	var drivers []Driver
//	driverID := 1
//
//	usedLoadsCount := 0
//	for len(loads) != usedLoadsCount {
//		var driverLoads []Load
//		var totalDistance float64
//
//		dropoffPoint := ZeroPoint
//
//		for i := 0; i < len(loads) && usedLoadsCount < len(loads); i++ {
//			if loads[i].isUsed {
//				continue
//			}
//
//			if totalDistance+ // prev distance
//				euclideanDistance(dropoffPoint, loads[i].pickup)+ // from prev point to next point
//				loads[i].distance+ // load distance
//				euclideanDistance(loads[i].dropoff, Point{}) /* distance home */ <= maxDriveTime {
//
//				driverLoads = append(driverLoads, loads[i])
//				totalDistance += euclideanDistance(dropoffPoint, loads[i].pickup) + loads[i].distance
//
//				dropoffPoint = loads[i].dropoff
//				loads[i].isUsed = true
//				usedLoadsCount++
//			}
//		}
//
//		totalDistance += euclideanDistance(dropoffPoint, ZeroPoint) // don't forget to count way home
//		driver := Driver{driverID, driverLoads, totalDistance}
//		drivers = append(drivers, driver)
//		driverID++
//	}
//
//	return drivers
//}

//
//func sortForPoint(loads []Load, point Point) {
//	sort.Slice(loads, func(i, j int) bool {
//		if loads[i].isUsed && loads[j].isUsed {
//			return false
//		}
//		if loads[i].isUsed {
//			return false
//		}
//		if loads[j].isUsed {
//			return true
//		}
//
//		d1, d2 := euclideanDistance(point, loads[i].pickup), euclideanDistance(point, loads[j].pickup)
//
//		if d1 < d2 {
//			return true
//		}
//
//		return false
//	})
//}

// greedy heap, barely better than just greedy one, dunno if it's useful considering it affects performance by order of magnitude
//func assignLoadsToDrivers(loads []Load) []Driver {
//	var drivers []Driver
//	driverID := 1
//	usedLoadsCount := 0
//	sortForPoint(loads, ZeroPoint)
//
//	for len(loads) != usedLoadsCount {
//		var driverLoads []Load
//		var totalDistance float64
//
//		dropoffPoint := ZeroPoint
//
//		for i := 0; i < len(loads) && usedLoadsCount < len(loads); i++ {
//			if loads[i].isUsed {
//				continue
//			}
//
//			if totalDistance+ // prev distance
//				euclideanDistance(dropoffPoint, loads[i].pickup)+ // from prev point to next point
//				loads[i].distance+ // load distance
//				euclideanDistance(loads[i].dropoff, Point{}) /* distance home */ <= maxDriveTime {
//
//				driverLoads = append(driverLoads, loads[i])
//				totalDistance += euclideanDistance(dropoffPoint, loads[i].pickup) + loads[i].distance
//
//				dropoffPoint = loads[i].dropoff
//				loads[i].isUsed = true
//				usedLoadsCount++
//				sortForPoint(loads, dropoffPoint) // could be optimized by reducing slice loads
//				i = -1
//			}
//		}
//
//		totalDistance += euclideanDistance(dropoffPoint, ZeroPoint) // don't forget to count way home
//		driver := Driver{driverID, driverLoads, totalDistance}
//		drivers = append(drivers, driver)
//		driverID++
//	}
//
//	return drivers
//}
