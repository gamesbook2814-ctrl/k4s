export interface SubCategory {
  name: string;
  href: string;
  isNew?: boolean;
}

export interface MegaMenuSection {
  title: string;
  items: SubCategory[];
}

export const mainCategories = [
  { name: "All Products", href: "/products" },
  { name: "Same Day Delivery", href: "/same-day-delivery" },
  { name: "Calendars & Diaries", href: "/calendars-diaries" },
  { name: "Packaging", href: "/packaging" },
  { name: "Apparel", href: "/apparel" },
  { name: "Stationery", href: "/stationery" },
  { name: "Corporate Gifts", href: "/corporate-gifts" },
  { name: "Photo Gifts", href: "/photo-gifts" },
  { name: "Drinkware", href: "/drinkware" },
];

export const allProductsMegaMenu: MegaMenuSection[] = [
  {
    title: "Browse All Categories",
    items: [
      { name: "Calendars & Diaries", href: "/calendars-diaries" },
      { name: "Apparels", href: "/apparel" },
      { name: "Photo Gifts", href: "/photo-gifts" },
      { name: "Rewards and Recognition", href: "/rewards-recognition" },
      { name: "Business Stationery", href: "/stationery" },
      { name: "Corporate Gifts", href: "/corporate-gifts" },
      { name: "Packaging", href: "/packaging" },
      { name: "Marketing & Promo", href: "/marketing-promo" },
      { name: "Premium Products", href: "/premium-products" },
      { name: "Drinkware", href: "/drinkware", isNew: true },
      { name: "Gift Hampers", href: "/gift-hampers" },
      { name: "Sample Kit", href: "/sample-kit" },
    ],
  },
  {
    title: "Best Sellers",
    items: [
      { name: "Visiting Cards", href: "/products/visiting-cards" },
      { name: "Round Neck T-shirts", href: "/products/round-neck-tshirts" },
      { name: "Labels", href: "/products/labels" },
      { name: "Booklets", href: "/products/booklets" },
      { name: "Stickers", href: "/products/stickers" },
      { name: "Brochures", href: "/products/brochures" },
      { name: "Notebooks", href: "/products/notebooks" },
      { name: "Mugs", href: "/products/mugs" },
      { name: "Button Badge", href: "/products/button-badge" },
      { name: "Photo Books", href: "/products/photo-books" },
      { name: "Tote Bags", href: "/products/tote-bags" },
      { name: "Photo Prints", href: "/products/photo-prints" },
    ],
  },
  {
    title: "Calendars & Diaries",
    items: [
      { name: "View All →", href: "/calendars-diaries" },
      { name: "Everyday Collection", href: "/products/everyday-calendars" },
      { name: "Classic Collection", href: "/products/classic-calendars" },
      { name: "Prestige Collection", href: "/products/prestige-calendars" },
      { name: "Calendars", href: "/products/calendars" },
    ],
  },
  {
    title: "Apparels",
    items: [
      { name: "View All →", href: "/apparel" },
      { name: "Custom Round Neck T-shirts", href: "/products/round-neck-tshirts" },
      { name: "Custom Polo T-shirts", href: "/products/polo-tshirts" },
      { name: "Custom Hoodies", href: "/products/hoodies" },
      { name: "Custom Caps", href: "/products/caps" },
    ],
  },
  {
    title: "Photo Gifts",
    items: [
      { name: "View All →", href: "/photo-gifts" },
      { name: "Photo Frames", href: "/products/photo-frames" },
      { name: "Photo Prints", href: "/products/photo-prints" },
      { name: "Acrylic Photo Prints", href: "/products/acrylic-prints" },
      { name: "PhotoBooks", href: "/products/photo-books", isNew: true },
    ],
  },
  {
    title: "Rewards and Recognition",
    items: [
      { name: "View All →", href: "/rewards-recognition" },
      { name: "Crystal & Acrylic Awards", href: "/products/crystal-awards" },
      { name: "Wooden Trophies & Mementos", href: "/products/wooden-trophies" },
      { name: "Certificates", href: "/products/certificates" },
      { name: "Medals", href: "/products/medals" },
    ],
  },
  {
    title: "Packaging",
    items: [
      { name: "View All →", href: "/packaging" },
      { name: "Flat Mailer Boxes", href: "/products/mailer-boxes" },
      { name: "Stickers", href: "/products/stickers" },
      { name: "Labels", href: "/products/labels" },
      { name: "Tote Bags", href: "/products/tote-bags", isNew: true },
    ],
  },
  {
    title: "Marketing & Promo",
    items: [
      { name: "View All →", href: "/marketing-promo" },
      { name: "Banners", href: "/products/banners" },
      { name: "Booklets", href: "/products/booklets" },
      { name: "Brochures", href: "/products/brochures" },
      { name: "Standees", href: "/products/standees" },
    ],
  },
  {
    title: "Premium Products",
    items: [
      { name: "View All →", href: "/premium-products" },
      { name: "Business Cards", href: "/products/premium-business-cards" },
      { name: "Cards", href: "/products/premium-cards" },
      { name: "Stickers", href: "/products/premium-stickers" },
      { name: "Notebooks", href: "/products/premium-notebooks" },
    ],
  },
  {
    title: "Drinkware",
    items: [
      { name: "View All →", href: "/drinkware", isNew: true },
      { name: "Best Sellers", href: "/products/drinkware-bestsellers" },
      { name: "Premium Mugs", href: "/products/premium-mugs" },
      { name: "Premium Flasks", href: "/products/premium-flasks" },
      { name: "Premium Sipper Bottles", href: "/products/sipper-bottles" },
    ],
  },
  {
    title: "Business Stationery",
    items: [
      { name: "View All →", href: "/stationery" },
      { name: "Business Cards", href: "/products/business-cards" },
      { name: "Letterheads", href: "/products/letterheads" },
      { name: "Stamps", href: "/products/stamps" },
      { name: "Envelopes", href: "/products/envelopes" },
    ],
  },
  {
    title: "Gift Hampers",
    items: [
      { name: "View All →", href: "/gift-hampers" },
      { name: "Employee Engagement Kits", href: "/products/employee-kits" },
      { name: "Welcome Kits", href: "/products/welcome-kits" },
      { name: "Awards & Trophies", href: "/products/awards-trophies" },
      { name: "Sustainable Kits", href: "/products/sustainable-kits" },
    ],
  },
  {
    title: "Corporate Gifts",
    items: [
      { name: "View All →", href: "/corporate-gifts" },
      { name: "Drinkware", href: "/drinkware" },
      { name: "Backpack", href: "/products/backpack", isNew: true },
      { name: "Awards & Trophies", href: "/products/awards-trophies" },
      { name: "Employee Engagement Kits", href: "/products/employee-kits" },
    ],
  },
  {
    title: "Sample Kit",
    items: [
      { name: "View All →", href: "/sample-kit" },
      { name: "Flat Mailer Boxes", href: "/products/mailer-boxes" },
      { name: "Stickers", href: "/products/stickers" },
      { name: "Flexible Pouches", href: "/products/flexible-pouches" },
      { name: "Tote Bags", href: "/products/tote-bags" },
    ],
  },
  {
    title: "Shop By Needs",
    items: [
      { name: "For Startups", href: "/shop-by/startups" },
      { name: "Events & Promotions", href: "/shop-by/events" },
      { name: "Cafe & Restaurant Essentials", href: "/shop-by/cafe-restaurant" },
      { name: "Employee Engagement kits", href: "/products/employee-kits" },
    ],
  },
  {
    title: "Stores & Services",
    items: [
      { name: "Same Day Delivery", href: "/same-day-delivery" },
      { name: "Bulk Buying", href: "/bulk-buying" },
      { name: "Store Locator", href: "/store-locator" },
      { name: "New Launches", href: "/new-launches" },
      { name: "Blog", href: "/blog" },
    ],
  },
];

export const categoryMegaMenus: Record<string, MegaMenuSection[]> = {
  "Calendars & Diaries": [
    {
      title: "Calendars & Diaries",
      items: [
        { name: "View All →", href: "/calendars-diaries" },
        { name: "Everyday Collection", href: "/products/everyday-calendars" },
        { name: "Classic Collection", href: "/products/classic-calendars" },
        { name: "Prestige Collection", href: "/products/prestige-calendars" },
        { name: "Wall Calendars", href: "/products/wall-calendars" },
        { name: "Desktop Calendars", href: "/products/desktop-calendars" },
      ],
    },
    {
      title: "Diaries",
      items: [
        { name: "Executive Diaries", href: "/products/executive-diaries" },
        { name: "Pocket Diaries", href: "/products/pocket-diaries" },
        { name: "Premium Diaries", href: "/products/premium-diaries" },
        { name: "Custom Diaries", href: "/products/custom-diaries" },
      ],
    },
  ],
  "Packaging": [
    {
      title: "Boxes",
      items: [
        { name: "Flat Mailer Boxes", href: "/products/mailer-boxes" },
        { name: "Product Boxes", href: "/products/product-boxes" },
        { name: "Gift Boxes", href: "/products/gift-boxes" },
        { name: "Food Boxes", href: "/products/food-boxes" },
      ],
    },
    {
      title: "Labels & Stickers",
      items: [
        { name: "Stickers", href: "/products/stickers" },
        { name: "Labels", href: "/products/labels" },
        { name: "Roll Labels", href: "/products/roll-labels" },
        { name: "Sheet Labels", href: "/products/sheet-labels" },
      ],
    },
    {
      title: "Bags",
      items: [
        { name: "Tote Bags", href: "/products/tote-bags", isNew: true },
        { name: "Paper Bags", href: "/products/paper-bags" },
        { name: "Cotton Bags", href: "/products/cotton-bags" },
      ],
    },
  ],
  "Apparel": [
    {
      title: "T-Shirts",
      items: [
        { name: "Round Neck T-shirts", href: "/products/round-neck-tshirts" },
        { name: "Polo T-shirts", href: "/products/polo-tshirts" },
        { name: "V-Neck T-shirts", href: "/products/vneck-tshirts" },
        { name: "Full Sleeve T-shirts", href: "/products/full-sleeve-tshirts" },
      ],
    },
    {
      title: "Outerwear",
      items: [
        { name: "Hoodies", href: "/products/hoodies" },
        { name: "Sweatshirts", href: "/products/sweatshirts" },
        { name: "Jackets", href: "/products/jackets" },
      ],
    },
    {
      title: "Accessories",
      items: [
        { name: "Caps", href: "/products/caps" },
        { name: "Aprons", href: "/products/aprons" },
        { name: "Bags", href: "/products/bags" },
      ],
    },
  ],
  "Corporate Gifts": [
    {
      title: "Corporate Gifts",
      items: [
        { name: "View All →", href: "/corporate-gifts" },
        { name: "Drinkware", href: "/drinkware" },
        { name: "Backpacks", href: "/products/backpack", isNew: true },
        { name: "Awards & Trophies", href: "/products/awards-trophies" },
        { name: "Employee Engagement Kits", href: "/products/employee-kits" },
      ],
    },
    {
      title: "Gift Hampers",
      items: [
        { name: "Welcome Kits", href: "/products/welcome-kits" },
        { name: "Festival Hampers", href: "/products/festival-hampers" },
        { name: "Onboarding Kits", href: "/products/onboarding-kits" },
        { name: "Sustainable Kits", href: "/products/sustainable-kits" },
      ],
    },
    {
      title: "Premium",
      items: [
        { name: "Executive Gifts", href: "/products/executive-gifts" },
        { name: "Leather Goods", href: "/products/leather-goods" },
        { name: "Tech Accessories", href: "/products/tech-accessories" },
      ],
    },
  ],
  "Photo Gifts": [
    {
      title: "Photo Products",
      items: [
        { name: "Photo Frames", href: "/products/photo-frames" },
        { name: "Photo Prints", href: "/products/photo-prints" },
        { name: "Acrylic Photo Prints", href: "/products/acrylic-prints" },
        { name: "Canvas Prints", href: "/products/canvas-prints" },
        { name: "PhotoBooks", href: "/products/photo-books", isNew: true },
      ],
    },
    {
      title: "Personalized Gifts",
      items: [
        { name: "Photo Mugs", href: "/products/photo-mugs" },
        { name: "Photo Cushions", href: "/products/photo-cushions" },
        { name: "Photo Magnets", href: "/products/photo-magnets" },
        { name: "Photo Keychains", href: "/products/photo-keychains" },
      ],
    },
  ],
  "Drinkware": [
    {
      title: "Mugs",
      items: [
        { name: "Ceramic Mugs", href: "/products/ceramic-mugs" },
        { name: "Magic Mugs", href: "/products/magic-mugs" },
        { name: "Travel Mugs", href: "/products/travel-mugs" },
        { name: "Premium Mugs", href: "/products/premium-mugs" },
      ],
    },
    {
      title: "Bottles & Flasks",
      items: [
        { name: "Sipper Bottles", href: "/products/sipper-bottles" },
        { name: "Premium Flasks", href: "/products/premium-flasks" },
        { name: "Sports Bottles", href: "/products/sports-bottles" },
        { name: "Copper Bottles", href: "/products/copper-bottles" },
      ],
    },
  ],
  "Stationery": [
    {
      title: "Business Cards",
      items: [
        { name: "Standard Cards", href: "/products/standard-cards" },
        { name: "Premium Cards", href: "/products/premium-cards" },
        { name: "Luxury Cards", href: "/products/luxury-cards" },
        { name: "Spot UV Cards", href: "/products/spot-uv-cards" },
      ],
    },
    {
      title: "Office Stationery",
      items: [
        { name: "Letterheads", href: "/products/letterheads" },
        { name: "Envelopes", href: "/products/envelopes" },
        { name: "Notepads", href: "/products/notepads" },
        { name: "Stamps", href: "/products/stamps" },
      ],
    },
    {
      title: "Notebooks",
      items: [
        { name: "Hardcover Notebooks", href: "/products/hardcover-notebooks" },
        { name: "Spiral Notebooks", href: "/products/spiral-notebooks" },
        { name: "Customized Notebooks", href: "/products/customized-notebooks" },
      ],
    },
  ],
};
