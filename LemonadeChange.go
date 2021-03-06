/*
At a lemonade stand, each lemonade costs $5.

Customers are standing in a queue to buy from you, and order one at a time (in the order specified by bills).

Each customer will only buy one lemonade and pay with either a $5, $10, or $20 bill.  You must provide the correct change to each customer, so that the net transaction is that the customer pays $5.

Note that you don't have any change in hand at first.

Return true if and only if you can provide every customer with correct change.



Example 1:

Input: [5,5,5,10,20]
Output: true
Explanation:
From the first 3 customers, we collect three $5 bills in order.
From the fourth customer, we collect a $10 bill and give back a $5.
From the fifth customer, we give a $10 bill and a $5 bill.
Since all customers got correct change, we output true.
Example 2:

Input: [5,5,10]
Output: true
Example 3:

Input: [10,10]
Output: false
Example 4:

Input: [5,5,10,10,20]
Output: false
Explanation:
From the first two customers in order, we collect two $5 bills.
For the next two customers in order, we collect a $10 bill and give back a $5 bill.
For the last customer, we can't give change of $15 back because we only have two $10 bills.
Since not every customer received correct change, the answer is false.
*/

package main

import (
	"fmt"
)

func main() {
	tests := [][]int{[]int{5, 5, 5, 10, 20}, []int{5, 5, 10}, []int{10, 10}, []int{5, 5, 10, 10, 20}}

	for _, test := range tests {
		fmt.Println(lemonadeChange(test))
	}
}

func lemonadeChange(bills []int) bool {
	fivers := 0
	tens := 0
	twenties := 0

	for _, tx := range bills {
		if tx == 5 {
			fivers++
		}

		if tx == 10 {
			if fivers > 0 {
				fivers--
				tens++
			} else {
				return false
			}
		}

		if tx == 20 {
			if tens > 0 && fivers > 0 {
				tens--
				fivers--
				twenties++
			} else if fivers >= 3 {
				fivers = fivers - 3
				twenties++
			} else {
				return false
			}
		}
	}

	return true
}
