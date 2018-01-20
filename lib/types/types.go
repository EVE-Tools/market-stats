package types

import "time"

// RegionType is a region/type combination
type RegionType struct {
	RegionID int32
	TypeID   int32
}

// RegionStats contains regional statistics
type RegionStats struct {
	RegionID                                      int32     `json:"region_id"`
	TypeID                                        int32     `json:"type_id"`
	GeneratedAt                                   time.Time `json:"generated_at"`
	Date                                          time.Time `json:"date"`
	Highest                                       float64   `json:"highest"`
	Lowest                                        float64   `json:"lowest"`
	Average                                       float64   `json:"average"`
	Volume                                        int64     `json:"volume"`
	OrderCount                                    int64     `json:"order_count"`
	PreviousDate                                  time.Time `json:"previous_date"`
	PreviousHighest                               float64   `json:"previous_highest"`
	PreviousLowest                                float64   `json:"previous_lowest"`
	PreviousAverage                               float64   `json:"previous_average"`
	PreviousVolume                                int64     `json:"previous_volume"`
	PreviousOrderCount                            int64     `json:"previous_order_count"`
	WeekPriceWeightedAverage                      float64   `json:"week_price_weighted_average"`
	WeekPriceAverage                              float64   `json:"week_price_average"`
	WeekPriceAverageStandardDeviation             float64   `json:"week_price_average_standard_deviation"`
	WeekPriceAverageRelativeStandardDeviation     float64   `json:"week_price_average_relative_standard_deviation"`
	WeekISKVolumeAverage                          float64   `json:"week_isk_volume_average"`
	WeekISKVolumeAverageStandardDeviation         float64   `json:"week_isk_volume_average_standard_deviation"`
	WeekISKVolumeAverageRelativeStandardDeviation float64   `json:"week_isk_volume_average_relative_standard_deviation"`
	WeekOrderCountTotal                           int64     `json:"week_order_count_total"`
	WeekOrderCountAverage                         float64   `json:"week_order_count_average"`
	WeekOrderCountStandardDeviation               float64   `json:"week_order_count_standard_deviation"`
	WeekOrderCountRelativeStandardDeviation       float64   `json:"week_order_count_relative_standard_deviation"`
	WeekVolumeTotal                               int64     `json:"week_volume_total"`
	WeekVolumeAverage                             float64   `json:"week_volume_average"`
	WeekVolumeStandardDeviation                   float64   `json:"week_volume_standard_deviation"`
	WeekVolumeRelativeStandardDeviation           float64   `json:"week_volume_relative_standard_deviation"`
}
