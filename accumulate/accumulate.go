package accumulate

const testVersion = 1

/*Accumulate given a collection and an operation to perform on each element of the collection, returns a new collection containing the result of applying that operation to each element of the input collection.
 */
func Accumulate(collection []string, operation func(string) string) []string {
	var resultCollection []string
	for _, elem := range collection {
		result := operation(elem)
		resultCollection = append(resultCollection, result)
	}
	return resultCollection
}
