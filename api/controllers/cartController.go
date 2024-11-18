package controllers

import (
	"api/config"
	"api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// AddToCart add item
func AddToCart(c *fiber.Ctx) error {
	userID, _ := strconv.Atoi(c.Locals("user_id").(string))
	productID, _ := strconv.Atoi(c.Params("product_id"))
	quantity, _ := strconv.Atoi(c.Params("quantity"))
	var product models.Product
	if err := config.DB.First(&product, productID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	var cartItem models.Cart
	if err := config.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cartItem).Error; err == nil {

		cartItem.Quantity += quantity
		config.DB.Save(&cartItem)
	} else {
		cartItem = models.Cart{
			UserID:    uint(userID),
			ProductID: uint(productID),
			Quantity:  quantity,
		}
		config.DB.Create(&cartItem)
	}

	return c.JSON(fiber.Map{
		"message": "Product added to cart successfully",
	})
}

// RemoveFromCart remove item
func RemoveFromCart(c *fiber.Ctx) error {
	userID, _ := strconv.Atoi(c.Locals("user_id").(string))
	productID, _ := strconv.Atoi(c.Params("product_id"))

	var cartItem models.Cart
	if err := config.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cartItem).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found in cart",
		})
	}

	config.DB.Delete(&cartItem)
	return c.JSON(fiber.Map{
		"message": "Product removed from cart successfully",
	})
}

// UpdateCartQuantity update cart
func UpdateCartQuantity(c *fiber.Ctx) error {
	userID, _ := strconv.Atoi(c.Locals("user_id").(string))
	productID, _ := strconv.Atoi(c.Params("product_id"))
	quantity, _ := strconv.Atoi(c.Params("quantity"))

	var cartItem models.Cart
	if err := config.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cartItem).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found in cart",
		})
	}

	cartItem.Quantity = quantity
	config.DB.Save(&cartItem)
	return c.JSON(fiber.Map{
		"message": "Product quantity updated successfully",
	})
}

// GetCart show products
func GetCart(c *fiber.Ctx) error {
	userID, _ := strconv.Atoi(c.Locals("user_id").(string))

	var cartItems []models.Cart
	if err := config.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Cart is empty",
		})
	}

	var totalPrice float64
	for _, item := range cartItems {
		totalPrice += float64(item.Quantity) * item.Product.Price
	}

	return c.JSON(fiber.Map{
		"cart":       cartItems,
		"totalPrice": totalPrice,
		"itemCount":  len(cartItems),
	})
}

// GetCartTotal show the total price
func GetCartTotal(c *fiber.Ctx) error {

	userID, _ := strconv.Atoi(c.Locals("user_id").(string))

	var cartItems []models.Cart
	if err := config.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Cart is empty or user not found",
		})
	}

	var totalPrice float64
	for _, item := range cartItems {
		totalPrice += float64(item.Quantity) * item.Product.Price
	}

	return c.JSON(fiber.Map{
		"totalPrice": totalPrice,
	})
}
