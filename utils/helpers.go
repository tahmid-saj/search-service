package utils

func DeleteAtIndexSliceString(slice []string, index int) []string {
     
	// Append function used to append elements to a slice
	// first parameter as the slice to which the elements 
	// are to be added/appended second parameter is the 
	// element(s) to be appended into the slice
	// return value as a slice
	return append(slice[:index], slice[index+1:]...)
}