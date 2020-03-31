// See: https://leetcode.com/problems/design-underground-system/discuss/556223/Go-hash-solution
package designundergroundsystem

type Record struct {
	name string
	time int
}
type UndergroundSystem struct {
	metadata map[string]map[string][]int
	going    map[int]Record
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		metadata: make(map[string]map[string][]int),
		going:    make(map[int]Record),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.going[id] = Record{
		name: stationName,
		time: t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	startRecord := this.going[id]

	if this.metadata[startRecord.name] == nil {
		this.metadata[startRecord.name] = make(map[string][]int)
	}

	if len(this.metadata[startRecord.name][stationName]) == 0 {
		this.metadata[startRecord.name][stationName] = append(this.metadata[startRecord.name][stationName], 0, 0)
	}
	this.metadata[startRecord.name][stationName][0] += t - startRecord.time
	this.metadata[startRecord.name][stationName][1]++
	delete(this.going, id)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	if this.metadata[startStation] == nil {
		return 0.0
	}

	if len(this.metadata[startStation][endStation]) == 0 {
		return 0.0
	}
	return float64(this.metadata[startStation][endStation][0]) / float64(this.metadata[startStation][endStation][1])
}
