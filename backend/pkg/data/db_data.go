package data

import "tradelist/pkg/api"

func GetSubcategories() []api.Subcategory {
	var subcategories = []api.Subcategory{
		{CategoryId: 1, Name: "Accounting"},
		{CategoryId: 1, Name: "HR"},
		{CategoryId: 1, Name: "Legal"},
		{CategoryId: 1, Name: "Customer Service"},
		{CategoryId: 1, Name: "Healthcare"},
		{CategoryId: 1, Name: "Hospitality"},
		{CategoryId: 1, Name: "Housekeeping"},
		{CategoryId: 1, Name: "Software"},
		{CategoryId: 1, Name: "Accounting"},

		{CategoryId: 2, Name: "For Sale"},
		{CategoryId: 2, Name: "To Rent"},
		{CategoryId: 2, Name: "To Share"},
		{CategoryId: 2, Name: "Sublet"},
		{CategoryId: 2, Name: "Storage"},

		{CategoryId: 3, Name: "Appliances"},
		{CategoryId: 3, Name: "Audio equipment"},
		{CategoryId: 3, Name: "Books"},
		{CategoryId: 3, Name: "Clothes"},
		{CategoryId: 3, Name: "Computers"},
		{CategoryId: 3, Name: "Furniture"},
		{CategoryId: 3, Name: "Gym equipment"},
		{CategoryId: 3, Name: "Sports equipment"},

		{CategoryId: 4, Name: "Computers"},
		{CategoryId: 4, Name: "Entertainment"},
		{CategoryId: 4, Name: "Financial"},
		{CategoryId: 4, Name: "Mobile"},
		{CategoryId: 4, Name: "Pets"},
		{CategoryId: 4, Name: "Real Estate"},
		{CategoryId: 4, Name: "Travel & Tourism"},
		{CategoryId: 4, Name: "Wedding"},

		{CategoryId: 5, Name: "Artists"},
		{CategoryId: 5, Name: "Events"},
		{CategoryId: 5, Name: "Lost & Found"},
		{CategoryId: 5, Name: "Rideshare & Car Pooling"},
	}
	return subcategories
}
