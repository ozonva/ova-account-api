package utils

import "sort"

// spamList represents an example of a list of stop words.
var spamList = []string{
	"bad", "word", "shift",
}

// longSortedSpamList represents an example of an ordered list of stop words.
var longSortedSpamList = []string{"Accept credit cards", "Ad", "All new", "As seen on", "Bargain", "Beneficiary", "Billing", "Bonus", "Cards accepted", "Cash", "Certified", "Cheap", "Claims", "Clearance", "Compare rates", "Credit card offers", "Deal", "Debt", "Discount", "Fantastic", "In accordance with laws", "Income", "Investment", "Join millions", "Lifetime", "Loans", "Luxury", "Marketing solution", "Message contains", "Mortgage rates", "Name brand", "Offer", "Online marketing", "Opt in", "Pre-approved", "Quote", "Rates", "Refinance", "Removal", "Reserves the right", "Score", "Search engine", "Sent in compliance", "Subject to…", "Terms and conditions", "Trial", "Unlimited", "Warranty", "Web traffic", "Work from home"}

// FilterWords filters a list of words by stop words.
// The complexity is O(n*m).
func FilterWords(words []string) []string {
	return FilterSliceString(words, inSpam)
}

// inSpam checks whether the word is in the list of stop words.
func inSpam(word string) bool {
	for i := range spamList {
		if word == spamList[i] {
			return false
		}
	}

	return true
}

// FilterWords2  filters a list of words by stop words.
// The complexity is O(n*log(m)).
func FilterWords2(words []string) []string {
	return FilterSliceString(words, func(s string) bool {
		i := sort.SearchStrings(longSortedSpamList, s)
		if i >= len(longSortedSpamList) {
			return true
		}
		return longSortedSpamList[i] != s
	})
}
