package mux

var (
	Rowing_General_0x31       = 0x31
	Rowing_Additional_0x32    = 0x32
	Rowing_Additional_0x33    = 0x33
	Stroke_Data_0x35          = 0x35
	Stroke_Data_0x36          = 0x36
	Split_Interval_0x37       = 0x37
	Split_Interval_0x38       = 0x38
	Workout_Summary_0x39      = 0x39
	Workout_Summary_0x3A      = 0x3A
	Heart_Rate_Belt_Info_0x3B = 0x3B
	Workout_Summary_0x3C      = 0x3C
)

var PM5MultiplexedData = map[string]map[string]int{
	"Mux_0x31": mux0x31,
	"Mux_0x32": mux0x32,
	"Mux_0x33": mux0x33,
	"Mux_0x34": mux0x35,
}

//0x31 C2 rowing general status characteristic
var mux0x31 = map[string]int{
	"Elapsed_Time_Lo":         1,
	"Elapsed_Time_Mid":        2,
	"Elapsed_Time_High":       3,
	"Distance_Lo":             4,
	"Distance_Mid":            5,
	"Distance_High":           6,
	"Workout_Type":            7,
	"Interval_Type":           8,
	"Workout_State":           9,
	"Rowing_State":            10,
	"Stroke_State":            11,
	"Total_Work_Distance_Lo":  12,
	"Total_Work_Distance_Mid": 13,
	"Total_Work_Distance_Hi":  14,
	"Workout_Duration_Lo":     15,
	"Workout_Duration_Mid":    16,
	"Workout_Duration_Hi":     17,
	"Workout_Duration_Type":   18,
	"Drag_Factor":             19,
}

//0x32 C2 rowing additional status 1 characteristic
var mux0x32 = map[string]int{
	"Elapsed_Time_Lo":   1,
	"Elapsed_Time_Mid":  2,
	"Elapsed_Time_High": 3,
	"Speed_Lo":          4,
	"Speed_Hi":          5,
	"Stroke_Rate":       6,
	"Heartrate":         7,
	"Current_Pace_Lo":   8,
	"Current_Pace_Hi":   9,
	"Average_Pace_Lo":   10,
	"Average_Pace_Hi":   11,
	"Rest_Distance_Lo":  12,
	"Rest_Distance_Hi":  13,
	"Rest_Time_Lo":      14,
	"Rest_Time_Mid":     15,
	"Rest_Time_Hi":      16,
	"Average_Power_Lo":  17,
	"Average_Power_Hi":  18,
}

//0x33 C2 rowing additional status 2 characteristic
var mux0x33 = map[string]int{
	"Elapsed_Time_Lo":           1,
	"Elapsed_Time_Mid":          2,
	"Elapsed_Time_High":         3,
	"Interval_Count":            4,
	"Total_Calories_Lo":         5,
	"Total_Calories_Hi":         6,
	"Split_Int_Avg_Pace_Lo":     7,
	"Split_Int_Avg_Pace_Hi":     8,
	"Split_Int_Avg_Power_Lo":    9,
	"Split_Int_Avg_Power_Hi":    10,
	"Split_Int_Avg_Calories_Lo": 11,
	"Split_Int_Avg_Calories_Hi": 12,
	"Last_Split_Time_Lo":        13,
	"Last_Split_Time_Mid":       14,
	"Last_Split_Time_High":      15,
	"Last_Split_Distance_Lo":    16,
	"Last_Split_Distance_Mid":   17,
	"Last_Split_Distance_Hi":    18,
}

//0x35 C2 rowing stroke data characteristic
var mux0x35 = map[string]int{
	"Elapsed_Time_Lo":         1,
	"Elapsed_Time_Mid":        2,
	"Elapsed_Time_High":       3,
	"Distance_Lo":             4,
	"Distance_Mid":            5,
	"Distance_High":           6,
	"Drive_Length":            7,
	"Drive_Time":              8,
	"Stroke_Recovery_Time_Lo": 9,
	"Stroke_Recovery_Time_Hi": 10,
	"Stroke_Distance_Lo":      11,
	"Stroke_Distance_Hi":      12,
	"Peak_Drive_Force_Lo":     13,
	"Peak_Drive_Force_Hi":     14,
	"Average_Drive_Force_Lo":  15,
	"Average_Drive_Force_Hi":  16,
	"Stroke_Count_Lo":         17,
	"Stroke_Count_Hi":         18,
}


