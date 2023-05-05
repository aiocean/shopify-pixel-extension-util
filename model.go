package shopify_pixel_extension_util

// CheckoutCompleted event logs when a visitor completes a purchase. This event is available on the order status and checkout pages.
type CheckoutCompletedEvent struct {
	ID        string                 `json:"id"`
	ClientID  string                 `json:"clientId"`
	Name      string                 `json:"name"`
	Timestamp string                 `json:"timestamp"`
	Context   map[string]interface{} `json:"context"`

	Data CheckoutCompletedData `json:"data"`
}

type CheckoutCompletedData struct {
	Checkout Checkout `json:"checkout"`
}

// CheckoutStarted event logs an instance of a buyer starting the checkout process. This event is available on the checkout page.
type CheckoutStartedEvent struct {
	ID        string                 `json:"id"`
	ClientID  string                 `json:"clientId"`
	Name      string                 `json:"name"`
	Timestamp string                 `json:"timestamp"`
	Context   map[string]interface{} `json:"context"`

	Data CheckoutStartedData `json:"data"`
}

type CheckoutStartedData struct {
	Checkout Checkout `json:"checkout"`
}

// CollectionViewedEvent event logs an instance where a buyer visited a product collection index page. This event is available on the online store page.
type CollectionViewedEvent struct {
	ID        string                 `json:"id"`
	ClientID  string                 `json:"clientId"`
	Name      string                 `json:"name"`
	Timestamp string                 `json:"timestamp"`
	Context   map[string]interface{} `json:"context"`

	Data CollectionViewedData `json:"data"`
}

type CollectionViewedData struct {
	Collection Collection `json:"connection"`
}

type PageViewedEvent struct {
	ID        string                 `json:"id"`
	ClientID  string                 `json:"clientId"`
	Name      string                 `json:"name"`
	Timestamp string                 `json:"timestamp"`
	Context   map[string]interface{} `json:"context"`
}

type PaymentInfoSubmittedEvent struct {
	ID        string                 `json:"id"`
	ClientID  string                 `json:"clientId"`
	Name      string                 `json:"name"`
	Timestamp string                 `json:"timestamp"`
	Context   map[string]interface{} `json:"context"`

	Data PaymentInfoSubmittedData `json:"data"`
}

type PaymentInfoSubmittedData struct {
	Checkout Checkout `json:"checkout"`
}

type ProductAddedToCartEvent struct {
	ID        string                 `json:"id"`
	ClientID  string                 `json:"clientId"`
	Name      string                 `json:"name"`
	Timestamp string                 `json:"timestamp"`
	Context   map[string]interface{} `json:"context"`

	Data ProductAddedToCartData `json:"data"`
}

type ProductAddedToCartData struct {
	CartLine CartLine `json:"cartLine"`
}

type CartLine struct {
	Cost        CartLineCost   `json:"cost"`
	Merchandise ProductVariant `json:"merchandise"`
	Quantity    int            `json:"quantity"`
}

type CartLineCost struct {
	TotalAmount Money `json:"totalAmount"`
}

// product_viewed logs an instance where a buyer visited a product details page
type ProductViewedEvent struct {
	ID        string                 `json:"id"`
	ClientID  string                 `json:"clientId"`
	Name      string                 `json:"name"`
	Timestamp string                 `json:"timestamp"`
	Context   map[string]interface{} `json:"context"`
	Data      ProductViewedData      `json:"data"`
}

type ProductViewedData struct {
	ProductVariant ProductVariant `json:"productVariant"`
}

// search_submitted logs an instance where a buyer performed a search on the storefront
type SearchSubmittedEvent struct {
	ID        string                 `json:"id"`
	ClientID  string                 `json:"clientId"`
	Name      string                 `json:"name"`
	Timestamp string                 `json:"timestamp"`
	Context   map[string]interface{} `json:"context"`
	Data      SearchSubmittedData    `json:"data"`
}

type SearchSubmittedData struct {
	SearchResult SearchResult `json:"searchResult"`
}

// SearchResult represents an object that contains the metadata of when a search has been performed.
type SearchResult struct {
	// Query is the search query that was executed.
	Query string `json:"query"`

	// ProductVariants is a list of product variants returned by the search query.
	ProductVariants []ProductVariant `json:"productVariants"`
}
type Collection struct {
	ID              string           `json:"id"`
	Title           string           `json:"title"`
	ProductVariants []ProductVariant `json:"productVariants"`
}

// Checkout represents a container for all the information required to checkout items and pay.
type Checkout struct {
	// Attributes is a list of attributes accumulated throughout the checkout process.
	Attributes []Attribute `json:"attributes"`

	// CurrencyCode is the three-letter code that represents the currency, for example, USD.
	// Supported codes include standard ISO 4217 codes, legacy codes, and non-standard codes.
	CurrencyCode string `json:"currencyCode"`

	// Email is the email attached to this checkout.
	Email string `json:"email"`

	// LineItems is a list of line item objects, each one containing information about an item in the checkout.
	LineItems []CheckoutLineItem `json:"lineItems"`

	// Order is the resulting order from a paid checkout.
	Order Order `json:"order"`

	// Phone is a unique phone number for the customer.
	// Formatted using E.164 standard. For example, +16135551111.
	Phone string `json:"phone"`

	// ShippingAddress is the shipping address to where the line items will be shipped.
	ShippingAddress MailingAddress `json:"shippingAddress"`

	// ShippingLine is the selected shipping rate by the customer, transitioned to a shipping_line object.
	ShippingLine ShippingRate `json:"shippingLine"`

	// SubtotalPrice is the price at checkout before duties, shipping, and taxes.
	SubtotalPrice Money `json:"subtotalPrice"`

	// Token is a unique identifier for a particular checkout.
	Token string `json:"token"`

	// TotalPrice is the sum of all the prices of all the items in the checkout, including duties, taxes, and discounts.
	TotalPrice Money `json:"totalPrice"`

	// TotalTax is the sum of all the taxes applied to the line items and shipping lines in the checkout.
	TotalTax Money `json:"totalTax"`
}

// MailingAddress represents a mailing address for customers and shipping.
type MailingAddress struct {
	// Address1 is the first line of the address, usually the street address or a P.O. Box number.
	Address1 string `json:"address1"`
	// Address2 is the second line of the address, usually an apartment, suite, or unit number.
	Address2 string `json:"address2"`
	// City is the name of the city, district, village, or town.
	City string `json:"city"`
	// Country is the name of the country.
	Country string `json:"country"`
	// CountryCode is the two-letter code that represents the country, generally following ISO 3166-1 alpha-2 guidelines.
	CountryCode string `json:"countryCode"`
	// Phone is the phone number for this mailing address as entered by the customer.
	Phone string `json:"phone"`
	// Province is the region of the address, such as the province, state, or district.
	Province string `json:"province"`
	// ProvinceCode is the two-letter code for the region, for example, ON.
	ProvinceCode string `json:"provinceCode"`
	// Zip is the ZIP or postal code of the address.
	Zip string `json:"zip"`
}

// ShippingRate represents a shipping rate to be applied to a checkout.
// See: https://shopify.dev/docs/api/pixels/customer-events#shippingrate
type ShippingRate struct {
	// Price is the price of this shipping rate.
	// See: https://shopify.dev/docs/api/pixels/customer-events#money
	Price Money
}

// Order represents a customer’s completed request to purchase one or more products from a shop. An order is created when a customer completes the checkout process.
type Order struct {
	// ID is the ID of the order.
	ID string `json:"id"`
}

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CheckoutLineItem struct {
	ID       string         `json:"id"`
	Quantity int            `json:"quantity"`
	Title    string         `json:"title,omitempty"`
	Variant  ProductVariant `json:"variant,omitempty"`
}

type Image struct {
	Src string `json:"src"`
}

type ProductVariant struct {
	ID      string  `json:"id"`
	Image   Image   `json:"image"`
	Price   Money   `json:"price"`
	Product Product `json:"product"`
	SKU     string  `json:"sku"`
	Title   string  `json:"title"`
}

type Product struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Vendor string `json:"vendor"`
	Type   string `json:"type"`
}

// Money represents a monetary value with currency.
type Money struct {
	// Amount is the decimal money amount.
	Amount float64 `json:"amount"`
	// CurrencyCode is the three-letter code that represents the currency, for example, USD. Supported codes include standard ISO 4217 codes, legacy codes, and non-standard codes.
	CurrencyCode string `json:"currencyCode"`
}
