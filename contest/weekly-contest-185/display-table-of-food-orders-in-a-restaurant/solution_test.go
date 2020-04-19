package displaytableoffoodordersinarestaurant

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDisplayTable(t *testing.T) {
	tests := []struct {
		in, want [][]string
	}{
		{in: [][]string{{"David", "3", "Ceviche"}, {"Corina", "10", "Beef Burrito"}, {"David", "3", "Fried Chicken"}, {"Carla", "5", "Water"}, {"Carla", "5", "Ceviche"}, {"Rous", "3", "Ceviche"}}, want: [][]string{{"Table", "Beef Burrito", "Ceviche", "Fried Chicken", "Water"}, {"3", "0", "2", "1", "0"}, {"5", "0", "1", "0", "1"}, {"10", "1", "0", "0", "0"}}},
		{in: [][]string{{"James", "12", "Fried Chicken"}, {"Ratesh", "12", "Fried Chicken"}, {"Amadeus", "12", "Fried Chicken"}, {"Adam", "1", "Canadian Waffles"}, {"Brianna", "1", "Canadian Waffles"}}, want: [][]string{{"Table", "Canadian Waffles", "Fried Chicken"}, {"1", "2", "0"}, {"12", "0", "3"}}},
		{in: [][]string{{"Laura", "2", "Bean Burrito"}, {"Jhon", "2", "Beef Burrito"}, {"Melissa", "2", "Soda"}}, want: [][]string{{"Table", "Bean Burrito", "Beef Burrito", "Soda"}, {"2", "1", "1", "1"}}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := displayTable(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
