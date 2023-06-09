package repositories

import (
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/utils"
)

func GetDiscountById(discountID string) (*models.Discount, error) {
	discount := new(models.Discount)

	result := utils.DB.Find(&discount, "id = ?", discountID)
	if result.Error != nil {
		return nil, result.Error
	}

	return discount, nil
}

func GetDiscounts(userID string) ([]models.Discount, error) {
	user_discounts := make([]models.UserDiscount, 0)
	discounts := make([]models.Discount, 0)

	result1 := utils.DB.Find(&user_discounts, "user_id = ?", userID)
	if result1.Error != nil {
		return nil, result1.Error
	}

	for _, userDiscount := range user_discounts {
		discount := new(models.Discount)
		result2 := utils.DB.Find(&discount, "id = ?", userDiscount.DiscountID)
		if result2.Error != nil {
			return nil, result2.Error
		}

		discounts = append(discounts, *discount)
	}

	return discounts, nil
}

// Delete coupon from everyone
func DeleteCouponById(discountID string) error {
	if err := utils.DB.Where("discount_id = ?", discountID).Delete(&models.UserDiscount{}); err.Error != nil {
		return err.Error
	}

	if err := utils.DB.Where("id = ?", discountID).Delete(&models.Discount{}); err.Error != nil {
		return err.Error
	}

	return nil
}

// Remove specific coupon from specific user
func RemoveCouponFromUser(discountID string, userID string) error {
	if err := utils.DB.Where("discount_id = ? AND user_id = ?", discountID, userID).Delete(&models.UserDiscount{}); err.Error != nil {
		return err.Error
	}

	return nil
}

func UpdateCouponById(discountID string, updates map[string]interface{}) error {
	if err := utils.DB.Model(&models.Discount{}).Where("id = ?", discountID).Updates(updates); err.Error != nil {
		return err.Error
	}

	return nil
}

func CreateCoupon(discount *models.Discount) error {
	if err := utils.DB.Create(discount); err.Error != nil {
		return err.Error
	}

	return nil
}

func AddCouponToUser(userDiscount *models.UserDiscount) error {
	if err := utils.DB.Create(userDiscount); err.Error != nil {
		return err.Error
	}

	return nil
}
