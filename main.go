/*
	This is practice code to help understand the merge-sort algorithm.
	It borrows from a couple of helpful sources:
	
		* T - http://www.tutorialspoint.com/data_structures_algorithms/merge_sort_algorithm.htm
		* G - https://gist.github.com/jzelinskie

	This is meant for educational purposes only.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"bufio"
)

func main() {
	// get list of integers from file
	file, err := os.Open("integers_6.txt")
	if err != nil {         
			log.Fatal("Problem opening source file--",err)
	}
	defer file.Close()

	// put file contents into string array
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
			lines = append(lines, scanner.Text())
	}

	// convert string array to integer array
	var ints = []int{}
	var j int
	for i := 0;i<len(lines);i++ {   
			j, _ = strconv.Atoi(lines[i])
			ints = append(ints,j)
	}
	
	// merge and sort
	b := mergesort(ints)
	_ = b
	// display sorted ints
	//_ = b
	fmt.Println("Merge and sort complete!")
}

/*
Prototype source: T

procedure mergesort( var a as array )
	 if ( n == 1 ) return a

	 var l1 as array = a[0] ... a[n/2]
	 var l2 as array = a[n/2+1] ... a[n]

	 l1 = mergesort( l1 )
	 l2 = mergesort( l2 )

	 return merge( l1, l2 )
end procedure
*/
func mergesort(ary []int) []int {

	if(len(ary)<=1) {
		fmt.Println("length of",ary,"is 1 --  done with array")
		return ary  // this is the break from the recursion
	} else {
			fmt.Println("ary: ",ary)
	}

	ary1 := ary[:len(ary)/2]
	ary2 := ary[(len(ary)/2):]
	fmt.Println("  ary1: ",ary1)
	fmt.Println("  ary2: ",ary2)

	ary1 = mergesort(ary1)
	ary2 = mergesort(ary2)
	fmt.Println("Calling merge for",ary1,ary2)
	return merge(ary1,ary2)
}
/*
Prototype source: T

procedure merge( var a as array, var b as array )

	 var c as array

	 while ( a and b have elements )
			if ( a[0] > b[0] )
				 add b[0] to the end of c
				 remove b[0] from b
			else
				 add a[0] to the end of c
				 remove a[0] from a
			end if
	 end while
	 
	 while ( a has elements )
			add a[0] to the end of c
			remove a[0] from a
	 end while
	 
	 while ( b has elements )
			add b[0] to the end of c
			remove b[0] from b
	 end while
	 
	 return c
	
end procedure
*/
func merge(a1 []int, a2 []int) []int {
	var c []int	
	//fmt.Println("a1: ",a1)
	//fmt.Println("a2: ",a2)
	
	for len(a1)>0 || len(a2)>0 {		// G
		if len(a1)>0 && len(a2)>0 {		// G
			if a1[0] >= a2[0] {
				c = append(c,a2[0])
				a2 = a2[1:]
			} else {
				c = append(c,a1[0])
				a1 = a1[1:]
			}
		} else if len(a1)>0 {
			c = append(c,a1[0])
			a1 = a1[1:]								// using :0 ends array at 12312
		} else if len(a2)>0 {
			c = append(c,a2[0])
			a2 = a2[1:]								// using :0 ends array at 12312
		}
		

	}
	fmt.Println("  Sorted: ",c)
	return c

}