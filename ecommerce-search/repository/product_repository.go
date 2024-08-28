package repository

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/LGROW101/ecommerce-search/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Search(query string) ([]model.Product, error) {
	var products []model.Product

	// Sanitize and limit the input
	query = sanitizeInput(query)
	if len(query) == 0 {
		return products, fmt.Errorf("invalid search query")
	}

	// Get synonyms for the query
	var synonyms []string
	err := r.db.Raw("SELECT unnest(synonyms) FROM synonyms WHERE word = ?", strings.ToLower(query)).Scan(&synonyms).Error
	if err != nil {
		return products, fmt.Errorf("error fetching synonyms: %w", err)
	}

	// Limit the number of search terms
	searchTerms := limitSearchTerms(append([]string{query}, synonyms...), 5)

	// Create the base query
	baseQuery := r.db.Preload("Categories").
		Joins("LEFT JOIN product_categories ON products.product_id = product_categories.product_id").
		Joins("LEFT JOIN categories ON product_categories.category_id = categories.category_id")

	// Create the WHERE clause
	whereClause := "products.search_vector @@ plainto_tsquery('english', ?)"
	whereArgs := []interface{}{strings.Join(searchTerms, " | ")}

	for _, term := range searchTerms {
		whereClause += " OR products.search_text ILIKE ? OR categories.category_name ILIKE ?"
		whereArgs = append(whereArgs, "%"+term+"%", "%"+term+"%")
	}

	// Execute the query
	err = baseQuery.Where(whereClause, whereArgs...).
		Group("products.product_id").
		Find(&products).Error

	if err != nil {
		return products, fmt.Errorf("error executing search query: %w", err)
	}

	return products, nil
}

func (r *ProductRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.Preload("Categories").First(&product, id).Error
	return &product, err
}

func sanitizeInput(input string) string {
	// Remove any characters that aren't alphanumeric, spaces, or common punctuation
	reg := regexp.MustCompile(`[^a-zA-Z0-9\s\p{Thai}.,!?-]`)
	sanitized := reg.ReplaceAllString(input, "")

	// Limit the length of the input
	if len(sanitized) > 100 {
		sanitized = sanitized[:100]
	}

	return strings.TrimSpace(sanitized)
}

func limitSearchTerms(terms []string, maxTerms int) []string {
	if len(terms) > maxTerms {
		return terms[:maxTerms]
	}
	return terms
}
