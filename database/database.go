package database

import "fruits-api/global_product"

func InitFruits() {
	global_product.FruitsList = append(global_product.FruitsList, global_product.Fruits{ID: "1", Name: "Apple", Color: "Red", Image: "https://images.unsplash.com/photo-1568702846914-96b305d2aaeb?w=500", Quantity: 50, Price: 10.99, Description: "Crisp and juicy red apples, perfect for snacking."})
	global_product.FruitsList = append(global_product.FruitsList, global_product.Fruits{ID: "2", Name: "Banana", Color: "Yellow", Image: "https://images.unsplash.com/photo-1528825871115-3581a5387919?w=500", Quantity: 60, Price: 9.99, Description: "Sweet and creamy bananas, rich in potassium."})
	global_product.FruitsList = append(global_product.FruitsList, global_product.Fruits{ID: "3", Name: "Blueberry", Color: "Blue", Image: "https://images.unsplash.com/photo-1498557850523-fd3d118b962e?w=500", Quantity: 70, Price: 8.99, Description: "Tiny antioxidant powerhouses with sweet-tart flavor."})
	global_product.FruitsList = append(global_product.FruitsList, global_product.Fruits{ID: "4", Name: "Orange", Color: "Orange", Image: "https://images.unsplash.com/photo-1547514701-42782101795e?w=500", Quantity: 80, Price: 6.99, Description: "Juicy citrus fruits packed with vitamin C."})
	global_product.FruitsList = append(global_product.FruitsList, global_product.Fruits{ID: "5", Name: "Kiwi", Color: "Green", Image: "https://images.unsplash.com/photo-1550258987-190a2d41a8ba?w=500", Quantity: 90, Price: 7.99, Description: "Fuzzy brown exterior with vibrant green flesh inside."})
	global_product.FruitsList = append(global_product.FruitsList, global_product.Fruits{ID: "6", Name: "Strawberry", Color: "Red", Image: "https://images.unsplash.com/photo-1464965911861-746a04b4bca6?w=500", Quantity: 40, Price: 8.99, Description: "Sweet, heart-shaped berries perfect for desserts."})
}