# json-comparer
Json Comparer compares two files consist of json array. It is able to compare if the objects or attributes are not in the same order in the files.

## Comparing Algorithms
It supports changing comparing algorithm by implementing **DataSetComparer** interface. After that a new type should be added to **dataSetFactory**. Algorithm type can be selected via passing the algorithm type to *application/DataSetCompare.CompareDataSets* function.

There are currently two differrent type of comparing algorithms

### LoopingTwoDataSetComparer
LoopingTwoDataSetComparer stores the json objects inside a **JSONObject** array and compare the two datasets. It checks every objects if they are equal inside a nested for loop. When a matched object is found, the objects are deleted from the array not to compare them multiple times.

It has also a different type of algorithm to check if the objects are same. This algorith can be changed by implementing **ObjectComparer** interface.

It only supports one-level deep json objects.

#### FindAttributeByKeyObjectComparer
FindAttributeByKeyObjectComparer checks if objects are eqaul. Attributes of an objects stored in a map type. It checks if the attribute key exists in the other object and the values are equal.

### HashMapDataSetComparer
HashMapDataSetComparer stores the json object inside a map. It creates a hash key for every object and store it as a key of the object. It uses one for loop and check if the other dataset has a object with the same hash.

### Banchmark results for algorithm

##### Ordered Dataset:

HashMapDataSetComparer is faster if objects are stored randomly in the datasets. But if the orders are same in the datasets it is slower than LoopingTwoDataSetComparer.

##### Random ordered Dataset:
BenchmarkWithDirectlyFindAttributeKeyByMapKeyRandomOrdered-8   	1000000000	         0.0313 ns/op	       0 B/op	       0 allocs/op
BenchmarkWithDirectlyFindObjectWithHashMapRandomOrdered-8      	1000000000	         0.0198 ns/op	       0 B/op	       0 allocs/op

##### Ordered Dataset:
BenchmarkWithDirectlyFindAttributeKeyByMapKeyOrdered-8         	1000000000	         0.0103 ns/op	       0 B/op	       0 allocs/op
BenchmarkWithDirectlyFindObjectWithHashMapOrdered-8            	1000000000	         0.0188 ns/op	       0 B/op	       0 allocs/op


# Example
File names are passed as arguments to main.go shown below

go run main.go "./MOCK_DATA.json" "./MOCK_DATA copy.json"
